package postges

import (
	"comics/models"
	"comics/pkg/helper/helper"
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

type productAttributeRepo struct {
	db *pgxpool.Pool
}

func (u *productAttributeRepo) Create(ctx context.Context, req *models.CreateProductAttribute) (*models.PrimaryKey, error) {
	query := `
		INSERT INTO product_attributes (
			product_id,
			key,
			value,
		) VALUES ($1,$2,$3)
		RETURNING id;
	`

	var newID int
	err := u.db.QueryRow(ctx, query, req.ProductID, req.Key, req.Value).Scan(&newID)
	if err != nil {
		return nil, err
	}

	pKey := &models.PrimaryKey{
		Id: newID,
	}

	return pKey, nil
}

// GetByID retrieves a productAttribute by its ID.
func (u *productAttributeRepo) GetByID(ctx context.Context, req *models.PrimaryKey) (*models.ProductAttribute, error) {
	res := &models.ProductAttribute{}
	query := `SELECT
		id,
		product_id,
		key,
		value
	FROM
		product_attributes
	WHERE
		id = $1`

	err := u.db.QueryRow(ctx, query, req.Id).Scan(
		&res.ID,
		&res.ProductID,
		&res.Key,
		&res.Value,
	)
	if err != nil {
		return res, err
	}

	return res, nil
}

// GetList retrieves a list of productAttributes with pagination and optional search functionality.
func (u *productAttributeRepo) GetList(ctx context.Context, req *models.GetListProductAttributeRequest) (*models.GetListProductAttributeResponse, error) {
	res := &models.GetListProductAttributeResponse{}
	params := make(map[string]interface{})
	var arr []interface{}

	query := `SELECT
		id,
		product_id,
		key,
		value
	FROM
		product_attributes`
	filter := " WHERE 1=1"
	offset := " OFFSET 0"
	limit := " LIMIT 10"

	// Implement search on productAttribute_name only
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
	cQ := `SELECT count(1) FROM product_attributes` + filter
	cQ, arr = helper.ReplaceQueryParams(cQ, params)
	err := u.db.QueryRow(ctx, cQ, arr...).Scan(
		&res.Count,
	)
	if err != nil {
		return res, err
	}

	// Main query for retrieving productAttributes
	q := query + filter + offset + limit
	q, arr = helper.ReplaceQueryParams(q, params)
	rows, err := u.db.Query(ctx, q, arr...)
	if err != nil {
		return res, err
	}
	defer rows.Close()

	for rows.Next() {
		obj := &models.ProductAttribute{}
		err = rows.Scan(
			&obj.ID,
			&obj.ProductID,
			&obj.Key,
			&obj.Value,
		)
		if err != nil {
			return res, err
		}

		res.ProductAttributes = append(res.ProductAttributes, obj)
	}

	return res, nil
}

// Update updates a productAttribute in the productAttributes table.
func (u *productAttributeRepo) Update(ctx context.Context, req *models.UpdateProductAttribute) (int64, error) {
	query := `UPDATE product_attributes SET
		product_id = :product_id,
		key =:key,
		value = :value
	WHERE
		id = :id`

	params := map[string]interface{}{
		"id":         req.ID,
		"product_id": req.ProductID,
		"key": req.Key,
		"value": req.Value,
	}

	q, arr := helper.ReplaceQueryParams(query, params)
	result, err := u.db.Exec(ctx, q, arr...)
	if err != nil {
		return 0, err
	}

	rowsAffected := result.RowsAffected()

	return rowsAffected, err
}

// Delete deletes a productAttribute from the productAttributes table by its ID.
func (u *productAttributeRepo) Delete(ctx context.Context, req *models.PrimaryKey) (int64, error) {
	query := `DELETE FROM product_attributes WHERE id = $1`

	result, err := u.db.Exec(ctx, query, req.Id)
	if err != nil {
		return 0, err
	}

	rowsAffected := result.RowsAffected()

	return rowsAffected, err
}
