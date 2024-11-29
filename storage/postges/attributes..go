package postges

import (
	"comics/models"
	"comics/pkg/helper/helper"
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

type attributeRepo struct {
	db *pgxpool.Pool
}

func (u *attributeRepo) Create(ctx context.Context, req *models.CreateAttribute) (*models.PrimaryKey, error) {
	query := `
		INSERT INTO attributes (
			name,
			data_type,
			created_at,
			updated_at
		) VALUES ($1,$2,now(),now())
		RETURNING id;
	`

	var newID int
	err := u.db.QueryRow(ctx, query, req.Name,req.DataType).Scan(&newID)
	if err != nil {
		return nil, err
	}

	pKey := &models.PrimaryKey{
		Id: newID,
	}

	return pKey, nil
}

// GetByID retrieves a attribute by its ID.
func (u *attributeRepo) GetByID(ctx context.Context, req *models.PrimaryKey) (*models.Attribute, error) {
	res := &models.Attribute{}
	query := `SELECT
		id,
		name,
		data_type,
		created_at,
		updated_at
	FROM
		attributes
	WHERE
		id = $1`

	err := u.db.QueryRow(ctx, query, req.Id).Scan(
		&res.ID,
		&res.Name,
		&res.DataType,
		&res.CreatedAt,
		&res.UpdatedAt,
	)
	if err != nil {
		return res, err
	}

	return res, nil
}

// GetList retrieves a list of attributes with pagination and optional search functionality.
func (u *attributeRepo) GetList(ctx context.Context, req *models.GetListAttributeRequest) (*models.GetListAttributeResponse, error) {
	res := &models.GetListAttributeResponse{}
	params := make(map[string]interface{})
	var arr []interface{}

	query := `SELECT
		id,
		name,
		data_type,
		created_at,
		updated_at
	FROM
		attributes`
	filter := " WHERE 1=1"
	offset := " OFFSET 0"
	limit := " LIMIT 10"

	// Implement search on attribute_name only
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
	cQ := `SELECT count(1) FROM attributes` + filter
	cQ, arr = helper.ReplaceQueryParams(cQ, params)
	err := u.db.QueryRow(ctx, cQ, arr...).Scan(
		&res.Count,
	)
	if err != nil {
		return res, err
	}

	// Main query for retrieving attributes
	q := query + filter + offset + limit
	q, arr = helper.ReplaceQueryParams(q, params)
	rows, err := u.db.Query(ctx, q, arr...)
	if err != nil {
		return res, err
	}
	defer rows.Close()

	for rows.Next() {
		obj := &models.Attribute{}
		err = rows.Scan(
			&obj.ID,
			&obj.Name,
			&obj.DataType,
			&obj.CreatedAt,
			&obj.UpdatedAt,
		)
		if err != nil {
			return res, err
		}

		res.Attributes = append(res.Attributes, obj)
	}

	return res, nil
}

// Update updates a attribute in the attributes table.
func (u *attributeRepo) Update(ctx context.Context, req *models.UpdateAttribute) (int64, error) {
	query := `UPDATE attributes SET
		name = :name,
		data_type = :data_type,
		updated_at = now()
	WHERE
		id = :id`

	params := map[string]interface{}{
		"id":        req.ID,
		"name": req.Name,
		"data_type": req.DataType,
	}

	q, arr := helper.ReplaceQueryParams(query, params)
	result, err := u.db.Exec(ctx, q, arr...)
	if err != nil {
		return 0, err
	}

	rowsAffected := result.RowsAffected()

	return rowsAffected, err
}

// Delete deletes a attribute from the attributes table by its ID.
func (u *attributeRepo) Delete(ctx context.Context, req *models.PrimaryKey) (int64, error) {
	query := `DELETE FROM attributes WHERE id = $1`

	result, err := u.db.Exec(ctx, query, req.Id)
	if err != nil {
		return 0, err
	}

	rowsAffected := result.RowsAffected()

	return rowsAffected, err
}
