package postges

import (
	"comics/models"
	"comics/pkg/helper/helper"
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

type productRepo struct {
	db *pgxpool.Pool
}

func (u *productRepo) Create(ctx context.Context, req *models.CreateProduct) (*models.PrimaryKey, error) {
	query := `
		INSERT INTO products (
			name,
			description,
			price,
			stock_quantity,
			category_id,
			created_at,
			updated_at
		) VALUES ($1,$2,$3,$4,$5,now(),now())
		RETURNING id;
	`

	var newID int
	err := u.db.QueryRow(ctx, query,
		req.Name,
		req.Description,
		req.Price,
		req.StockQuantity,
		req, req.CategoryID,
	).Scan(&newID)
	if err != nil {
		return nil, err
	}

	pKey := &models.PrimaryKey{
		Id: newID,
	}

	return pKey, nil
}

// GetByID retrieves a role by its ID.
func (u *productRepo) GetByID(ctx context.Context, req *models.PrimaryKey) (*models.Product, error) {
	res := &models.Product{}
	query := `SELECT
		id,
		name,
		description,
		price,
		stock_quantity,
		category_id,
		created_at,
		updated_at
	FROM
		products
	WHERE
		id = $1`

	err := u.db.QueryRow(ctx, query, req.Id).Scan(
		&res.ID,
		&res.Name,
		&res.Description,
		&res.Price,
		&res.StockQuantity,
		&res.CategoryID,
		&res.CreatedAt,
		&res.UpdatedAt,
	)
	if err != nil {
		return res, err
	}

	return res, nil
}

// GetList retrieves a list of roles with pagination and optional search functionality.
func (u *productRepo) GetList(ctx context.Context, req *models.GetListProductRequest) (*models.GetListProductResponse, error) {
	res := &models.GetListProductResponse{}
	params := make(map[string]interface{})
	var arr []interface{}

	query := `SELECT
		id,
		name,
		description,
		price,
		stock_quantity,
		category_id,
		created_at,
		updated_at
	FROM
		products`
	filter := " WHERE 1=1"
	offset := " OFFSET 0"
	limit := " LIMIT 10"

	// Implement search on name only
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
	cQ := `SELECT count(1) FROM products` + filter
	cQ, arr = helper.ReplaceQueryParams(cQ, params)
	err := u.db.QueryRow(ctx, cQ, arr...).Scan(
		&res.Count,
	)
	if err != nil {
		return res, err
	}

	// Main query for retrieving roles
	q := query + filter + offset + limit
	q, arr = helper.ReplaceQueryParams(q, params)
	rows, err := u.db.Query(ctx, q, arr...)
	if err != nil {
		return res, err
	}
	defer rows.Close()

	for rows.Next() {
		obj := &models.Product{}
		err = rows.Scan(
			&obj.ID,
			&obj.Name,
			&obj.Description,
			&obj.Price,
			&obj.StockQuantity,
			&obj.CategoryID,
			&obj.CreatedAt,
			&obj.UpdatedAt,
		)
		if err != nil {
			return res, err
		}

		res.Products = append(res.Products, obj)
	}

	return res, nil
}

// Update updates a role in the roles table.
func (u *productRepo) Update(ctx context.Context, req *models.UpdateProduct) (int64, error) {
	query := `UPDATE products SET
		name = :name,
		decription = :decription,
		price = :price,
		stock_quantity = :stock_quantity,
		category_id = :category_id
		update_at=now()
	WHERE
		id = :id`

	params := map[string]interface{}{
		"id":          req.ID,
		"name":        req.Name,
		"description": req.Description,
		"price": req.Price,
		"stock_quantity": req.StockQuantity,
		"category_id": req.CategoryID,
		
	}

	q, arr := helper.ReplaceQueryParams(query, params)
	result, err := u.db.Exec(ctx, q, arr...)
	if err != nil {
		return 0, err
	}

	rowsAffected := result.RowsAffected()

	return rowsAffected, err
}

// Delete deletes a role from the roles table by its ID.
func (u *productRepo) Delete(ctx context.Context, req *models.PrimaryKey) (int64, error) {
	query := `DELETE FROM products WHERE id = $1`

	result, err := u.db.Exec(ctx, query, req.Id)
	if err != nil {
		return 0, err
	}

	rowsAffected := result.RowsAffected()

	return rowsAffected, err
}
