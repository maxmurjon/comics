package postges

import (
	"comics/models"
	"comics/pkg/helper/helper"
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

type rolePermissionRepo struct {
	db *pgxpool.Pool
}

func (u *rolePermissionRepo) Create(ctx context.Context, req *models.CreateRolePermission) (*models.PrimaryKey, error) {
	query := `
		INSERT INTO role_permissions (
			permission_id,
			role_id
		) VALUES ($1,$2
		)
		RETURNING id;
	`

	var newID int
	err := u.db.QueryRow(ctx, query, req.PermissionId,req.RoleId).Scan(&newID)
	if err != nil {
		return nil, err
	}

	pKey := &models.PrimaryKey{
		Id: newID,
	}

	return pKey, nil
}

// GetByID retrieves a rolePermission by its ID.
func (u *rolePermissionRepo) GetByID(ctx context.Context, req *models.PrimaryKey) (*models.RolePermission, error) {
	res := &models.RolePermission{}
	query := `SELECT
		permission_id,
		role_id
	FROM
		role_permissions
	WHERE
		role_id = $1`

	err := u.db.QueryRow(ctx, query, req.Id).Scan(
		&res.PermissionId,
		&res.RoleId,
	)
	if err != nil {
		return res, err
	}

	return res, nil
}

// GetList retrieves a list of rolePermissions with pagination and optional search functionality.
func (u *rolePermissionRepo) GetList(ctx context.Context, req *models.GetListRolePermissionRequest) (*models.GetListRolePermissionResponse, error) {
	res := &models.GetListRolePermissionResponse{}
	params := make(map[string]interface{})
	var arr []interface{}

	query := `SELECT
		permission_id,
		role_id
	FROM
		role_permissions`
	filter := " WHERE 1=1"
	offset := " OFFSET 0"
	limit := " LIMIT 10"

	// Implement search on rolePermission_name only
	if len(req.Search) > 0 {
		params["search"] = req.Search
		filter += " AND user_role_name ILIKE '%' || :search || '%'"
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
	cQ := `SELECT count(1) FROM role_permissions` + filter
	cQ, arr = helper.ReplaceQueryParams(cQ, params)
	err := u.db.QueryRow(ctx, cQ, arr...).Scan(
		&res.Count,
	)
	if err != nil {
		return res, err
	}

	// Main query for retrieving rolePermissions
	q := query + filter + offset + limit
	q, arr = helper.ReplaceQueryParams(q, params)
	rows, err := u.db.Query(ctx, q, arr...)
	if err != nil {
		return res, err
	}
	defer rows.Close()

	for rows.Next() {
		obj := &models.RolePermission{}
		err = rows.Scan(
			&obj.PermissionId,
			&obj.RoleId,
		)
		if err != nil {
			return res, err
		}

		res.RolePermission = append(res.RolePermission, obj)
	}

	return res, nil
}

// Update updates a rolePermission in the rolePermissions table.
func (u *rolePermissionRepo) Update(ctx context.Context, req *models.UpdateRolePermission) (int64, error) {
	query := `UPDATE role_permissions SET
		role_id = :role_id,
		permission_id = :permission_id
	WHERE
		role_id = :role_id`

	params := map[string]interface{}{
		"role_id":        req.RoleId,
		"permission_id": req.PermissionId,
	}

	q, arr := helper.ReplaceQueryParams(query, params)
	result, err := u.db.Exec(ctx, q, arr...)
	if err != nil {
		return 0, err
	}

	rowsAffected := result.RowsAffected()

	return rowsAffected, err
}

// Delete deletes a rolePermission from the rolePermissions table by its ID.
func (u *rolePermissionRepo) Delete(ctx context.Context, req *models.PrimaryKey) (int64, error) {
	query := `DELETE FROM role_permissions WHERE id = $1`

	result, err := u.db.Exec(ctx, query, req.Id)
	if err != nil {
		return 0, err
	}

	rowsAffected := result.RowsAffected()

	return rowsAffected, err
}
