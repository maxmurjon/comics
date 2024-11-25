package postges

import (
	"comics/models"
	"comics/pkg/helper/helper"
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

type productImagesRepo struct {
	db *pgxpool.Pool
}

func (u *productImagesRepo) Create(ctx context.Context, req *models.CreateProductImage) (*models.PrimaryKey, error) {
	query := `
		INSERT INTO product_images (
			product_id,
			image_url,
			is_primary
		) VALUES ($1,$2,$3)
		RETURNING id;
	`

	var newID int
	err := u.db.QueryRow(ctx, query, req.ProductID, req.ImageUrl, req.IsPrimary).Scan(&newID)
	if err != nil {
		return nil, err
	}

	pKey := &models.PrimaryKey{
		Id: newID,
	}

	return pKey, nil
}

// GetByID retrieves a productImages by its ID.
func (u *productImagesRepo) GetByID(ctx context.Context, req *models.PrimaryKey) (*models.ProductImage, error) {
	res := &models.ProductImage{}
	query := `SELECT
		id,
		product_id,
		image_url,
		is_primary
	FROM
		product_images
	WHERE
		id = $1`

	err := u.db.QueryRow(ctx, query, req.Id).Scan(
		&res.ID,
		&res.ProductID,
		&res.ImageUrl,
		&res.IsPrimary,
	)
	if err != nil {
		return res, err
	}

	return res, nil
}

// GetList retrieves a list of productImagess with pagination and optional search functionality.
func (u *productImagesRepo) GetList(ctx context.Context, req *models.GetListProductImageRequest) (*models.GetListProductImageResponse, error) {
	res := &models.GetListProductImageResponse{}
	params := make(map[string]interface{})
	var arr []interface{}

	query := `SELECT
		id,
		product_id,
		image_url,
		is_primary
	FROM
		product_images`
	filter := " WHERE 1=1"
	offset := " OFFSET 0"
	limit := " LIMIT 10"

	// Implement search on productImages_name only
	if len(req.Search) > 0 {
		params["search"] = req.Search
		filter += " AND product_id ILIKE '%' || :search || '%'"
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
	cQ := `SELECT count(1) FROM product_images` + filter
	cQ, arr = helper.ReplaceQueryParams(cQ, params)
	err := u.db.QueryRow(ctx, cQ, arr...).Scan(
		&res.Count,
	)
	if err != nil {
		return res, err
	}

	// Main query for retrieving productImagess
	q := query + filter + offset + limit
	q, arr = helper.ReplaceQueryParams(q, params)
	rows, err := u.db.Query(ctx, q, arr...)
	if err != nil {
		return res, err
	}
	defer rows.Close()

	for rows.Next() {
		obj := &models.ProductImage{}
		err = rows.Scan(
			&obj.ID,
			&obj.ProductID,
			&obj.ImageUrl,
			&obj.IsPrimary,
		)
		if err != nil {
			return res, err
		}

		res.ProductImages = append(res.ProductImages, obj)
	}

	return res, nil
}

// Update updates a productImages in the productImagess table.
func (u *productImagesRepo) Update(ctx context.Context, req *models.UpdateProductImage) (int64, error) {
	query := `UPDATE product_images SET
		product_id = :product_id,
		image_url =:image_url,
		is_primary = :is_primary
	WHERE
		id = :id`

	params := map[string]interface{}{
		"id":         req.ID,
		"product_id": req.ProductID,
		"is_primary": req.IsPrimary,
	}

	q, arr := helper.ReplaceQueryParams(query, params)
	result, err := u.db.Exec(ctx, q, arr...)
	if err != nil {
		return 0, err
	}

	rowsAffected := result.RowsAffected()

	return rowsAffected, err
}

// Delete deletes a productImages from the productImagess table by its ID.
func (u *productImagesRepo) Delete(ctx context.Context, req *models.PrimaryKey) (int64, error) {
	query := `DELETE FROM product_images WHERE id = $1`

	result, err := u.db.Exec(ctx, query, req.Id)
	if err != nil {
		return 0, err
	}

	rowsAffected := result.RowsAffected()

	return rowsAffected, err
}
