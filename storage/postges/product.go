package postges

import (
	"comics/models"
	"comics/pkg/helper/helper"
	"context"
	"encoding/json"
	"fmt"

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
			created_at,
			updated_at
		) VALUES ($1,$2,$3,$4,now(),now())
		RETURNING id;
	`

	var newID int
	err := u.db.QueryRow(ctx, query,
		req.Name,
		req.Description,
		req.Price,
		req.StockQuantity,
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
	query := `
		WITH product_data AS (
			SELECT
				p.id,
				p.name,
				p.description,
				p.price,
				p.stock_quantity,
				p.created_at,
				p.updated_at,
				COALESCE(json_agg(DISTINCT jsonb_build_object(
					'id', pi.id,
					'url', pi.image_url,
					'is_primary', pi.is_primary
				)) FILTER (WHERE pi.id IS NOT NULL), '[]') AS images,
				COALESCE(json_agg(DISTINCT jsonb_build_object(
					'id', c.id,
					'name', c.name,
					'description', c.description
				)) FILTER (WHERE c.id IS NOT NULL), '[]') AS categories
			FROM products p
			LEFT JOIN product_images pi ON pi.product_id = p.id
			LEFT JOIN product_categories pc ON pc.product_id = p.id
			LEFT JOIN categories c ON c.id = pc.category_id
			GROUP BY p.id
		)
		SELECT
			COUNT(*) OVER() AS total_count,
			pd.*
		FROM product_data pd
		ORDER BY pd.created_at DESC
		OFFSET $1
		LIMIT $2;
	`

	rows, err := u.db.Query(ctx, query, req.Offset, req.Limit)
	if err != nil {
		return nil, fmt.Errorf("querying products: %w", err)
	}
	defer rows.Close()

	// Prepare to collect products
	var products []*models.ProductInfo

	for rows.Next() {
		var product models.ProductInfo
		var images, categories []byte

		// Scan the row into variables
		err := rows.Scan(
			&res.Count,
			&product.ID,
			&product.Name,
			&product.Description,
			&product.Price,
			&product.StockQuantity,
			&product.CreatedAt,
			&product.UpdatedAt,
			&images,
			&categories,
		)
		if err != nil {
			return nil, fmt.Errorf("scanning product row: %w", err)
		}

		// Parse JSONB data
		if err := json.Unmarshal(images, &product.Images); err != nil {
			return nil, fmt.Errorf("unmarshalling images: %w", err)
		}
		if err := json.Unmarshal(categories, &product.Categories); err != nil {
			return nil, fmt.Errorf("unmarshalling categories: %w", err)
		}

		products = append(products, &product)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("iterating rows: %w", err)
	}

	// Assign result
	res.Products = products
	return res, nil
}

// Update updates a role in the roles table.
func (u *productRepo) Update(ctx context.Context, req *models.UpdateProduct) (int64, error) {
	query := `UPDATE products SET
		name = :name,
		description = :description,
		price = :price,
		stock_quantity = :stock_quantity,
		updated_at=now()
	WHERE
		id = :id`

	params := map[string]interface{}{
		"id":             req.ID,
		"name":           req.Name,
		"description":    req.Description,
		"price":          req.Price,
		"stock_quantity": req.StockQuantity,
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
