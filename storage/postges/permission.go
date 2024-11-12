package postges

import (
	"comics/models"
	"comics/pkg/helper/helper"
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

type permissionRepo struct {
	db *pgxpool.Pool
}

func (u *permissionRepo) Create(ctx context.Context, req *models.CreatePermission) (*models.PrimaryKey, error) {
	query := `
		INSERT INTO permissions (
			name,
			description
		) VALUES ($1,$2
		)
		RETURNING id;
	`

	var newID int
	err := u.db.QueryRow(ctx, query, req.PermissionName,req.Description).Scan(&newID)
	if err != nil {
		return nil, err
	}

	pKey := &models.PrimaryKey{
		Id: newID,
	}

	return pKey, nil
}

// GetByID retrieves a permission by its ID.
func (u *permissionRepo) GetByID(ctx context.Context, req *models.PrimaryKey) (*models.Permission, error) {
	res := &models.Permission{}
	query := `SELECT
		id,
		name,
		description
	FROM
		permissions
	WHERE
		id = $1`

	err := u.db.QueryRow(ctx, query, req.Id).Scan(
		&res.Id,
		&res.PermissionName,
		&res.Description,
	)
	if err != nil {
		return res, err
	}

	return res, nil
}

// GetList retrieves a list of permissions with pagination and optional search functionality.
func (u *permissionRepo) GetList(ctx context.Context, req *models.GetListPermissionRequest) (*models.GetListPermissionResponse, error) {
	res := &models.GetListPermissionResponse{}
	params := make(map[string]interface{})
	var arr []interface{}

	query := `SELECT
		id,
		name,
		description
	FROM
		permissions`
	filter := " WHERE 1=1"
	offset := " OFFSET 0"
	limit := " LIMIT 10"

	// Implement search on permission_name only
	if len(req.Search) > 0 {
		params["search"] = req.Search
		filter += " AND permission_name ILIKE '%' || :search || '%'"
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
	cQ := `SELECT count(1) FROM permissions` + filter
	cQ, arr = helper.ReplaceQueryParams(cQ, params)
	err := u.db.QueryRow(ctx, cQ, arr...).Scan(
		&res.Count,
	)
	if err != nil {
		return res, err
	}

	// Main query for retrieving permissions
	q := query + filter + offset + limit
	q, arr = helper.ReplaceQueryParams(q, params)
	rows, err := u.db.Query(ctx, q, arr...)
	if err != nil {
		return res, err
	}
	defer rows.Close()

	for rows.Next() {
		obj := &models.Permission{}
		err = rows.Scan(
			&obj.Id,
			&obj.PermissionName,
			&obj.Description,
		)
		if err != nil {
			return res, err
		}

		res.Permissions = append(res.Permissions, obj)
	}

	return res, nil
}

// Update updates a permission in the permissions table.
func (u *permissionRepo) Update(ctx context.Context, req *models.UpdatePermission) (int64, error) {
	query := `UPDATE permissions SET
		name = :name,
		description =:description
	WHERE
		id = :id`

	params := map[string]interface{}{
		"id":        req.Id,
		"name": req.PermissionName,
		"description":req.Description,
	}

	q, arr := helper.ReplaceQueryParams(query, params)
	result, err := u.db.Exec(ctx, q, arr...)
	if err != nil {
		return 0, err
	}

	rowsAffected := result.RowsAffected()

	return rowsAffected, err
}

// Delete deletes a permission from the permissions table by its ID.
func (u *permissionRepo) Delete(ctx context.Context, req *models.PrimaryKey) (int64, error) {
	query := `DELETE FROM permissions WHERE id = $1`

	result, err := u.db.Exec(ctx, query, req.Id)
	if err != nil {
		return 0, err
	}

	rowsAffected := result.RowsAffected()

	return rowsAffected, err
}
