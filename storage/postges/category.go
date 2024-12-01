package postges

import (
	"comics/models"
	"comics/pkg/helper/helper"
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

type categoryRepo struct {
	db *pgxpool.Pool
}

func (u *categoryRepo) Create(ctx context.Context, req *models.CreateCategory) (*models.PrimaryKey, error) {
	query := `
		INSERT INTO categories (
			name,
			description,
			image_url
		) VALUES ($1,$2,$3)
		RETURNING id;
	`

	var newID int
	err := u.db.QueryRow(ctx, query, req.Name, req.Description, req.ImageUrl).Scan(&newID)
	if err != nil {
		return nil, err
	}

	pKey := &models.PrimaryKey{
		Id: newID,
	}

	return pKey, nil
}

// GetByID retrieves a category by its ID.
func (u *categoryRepo) GetByID(ctx context.Context, req *models.PrimaryKey) (*models.Category, error) {
	res := &models.Category{}
	query := `SELECT
		id,
		name,
		description,
		image_url
	FROM
		categories
	WHERE
		id = $1`

	err := u.db.QueryRow(ctx, query, req.Id).Scan(
		&res.ID,
		&res.Name,
		&res.Description,
		&res.ImageUrl,
	)
	if err != nil {
		return res, err
	}

	return res, nil
}

// GetList retrieves a list of categorys with pagination and optional search functionality.
func (u *categoryRepo) GetList(ctx context.Context, req *models.GetListCategoryRequest) (*models.GetListCategoryResponse, error) {
	res := &models.GetListCategoryResponse{}
	params := make(map[string]interface{})
	var arr []interface{}

	query := `SELECT
		id,
		name,
		description,
		image_url
	FROM
		categories`
	filter := " WHERE 1=1"
	offset := " OFFSET 0"
	limit := " LIMIT 10"

	// Implement search on category_name only
	if len(req.Search) > 0 {
		params["search"] = req.Search
		filter += " AND name ILIKE '%' || :search || '%'"
	}

	if req.Offset > 0 {
		params["offset"] = req.Offset
		offset = " OFFSET :offset"
	}

	if req.Limit > 0 {
		params["limit"] = req.Limit
		limit = " LIMIT :limit"
	}

	// Count query
	cQ := `SELECT count(1) FROM categories` + filter
	cQ, arr = helper.ReplaceQueryParams(cQ, params)
	err := u.db.QueryRow(ctx, cQ, arr...).Scan(
		&res.Count,
	)
	if err != nil {
		return res, err
	}

	// Main query for retrieving categorys
	q := query + filter + offset + limit
	q, arr = helper.ReplaceQueryParams(q, params)
	rows, err := u.db.Query(ctx, q, arr...)
	if err != nil {
		return res, err
	}
	defer rows.Close()

	for rows.Next() {
		obj := &models.Category{}
		err = rows.Scan(
			&obj.ID,
			&obj.Name,
			&obj.Description,
			&obj.ImageUrl,
		)
		if err != nil {
			return res, err
		}

		res.Categories = append(res.Categories, obj)
	}

	return res, nil
}

// Update updates a category in the categorys table.
func (u *categoryRepo) Update(ctx context.Context, req *models.UpdateCategory) (int64, error) {
	query := `UPDATE categories SET
		name = :name,
		description = :description,
		image_url = :image_url
	WHERE
		id = :id`

	params := map[string]interface{}{
		"id":         req.ID,
		"name":       req.Name,
		"desription": req.Description,
		"image_url":  req.ImageUrl,
	}

	q, arr := helper.ReplaceQueryParams(query, params)
	result, err := u.db.Exec(ctx, q, arr...)
	if err != nil {
		return 0, err
	}

	rowsAffected := result.RowsAffected()

	return rowsAffected, err
}

// Delete deletes a category from the categorys table by its ID.
func (u *categoryRepo) Delete(ctx context.Context, req *models.PrimaryKey) (int64, error) {
	query := `DELETE FROM categories WHERE id = $1`

	result, err := u.db.Exec(ctx, query, req.Id)
	if err != nil {
		return 0, err
	}

	rowsAffected := result.RowsAffected()

	return rowsAffected, err
}
