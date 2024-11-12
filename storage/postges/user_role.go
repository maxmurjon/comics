package postges

import (
	"comics/models"
	"comics/pkg/helper/helper"
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

type userRoleRepo struct {
	db *pgxpool.Pool
}

func (u *userRoleRepo) Create(ctx context.Context, req *models.CreateUserRole) (error) {
	query := `
		INSERT INTO user_roles (
			user_id,
			role_id
		) VALUES ($1,$2
		);
	`

	_,err := u.db.Exec(ctx, query, req.UserId,req.RoleId)
	if err != nil {
		return err
	}

	
	return nil
}

// GetByID retrieves a userRole by its ID.
func (u *userRoleRepo) GetByID(ctx context.Context, req string) (*models.UserRole, error) {
	res := &models.UserRole{}
	query := `SELECT
		user_id,
		role_id
	FROM
		user_roles
	WHERE
		user_id = $1`

	err := u.db.QueryRow(ctx, query, req).Scan(
		&res.UserId,
		&res.RoleId,
	)
	if err != nil {
		return res, err
	}

	return res, nil
}

// GetList retrieves a list of userRoles with pagination and optional search functionality.
func (u *userRoleRepo) GetList(ctx context.Context, req *models.GetListUserRoleRequest) (*models.GetListUserRoleResponse, error) {
	res := &models.GetListUserRoleResponse{}
	params := make(map[string]interface{})
	var arr []interface{}

	query := `SELECT
		user_id,
		role_id
	FROM
		user_roles`
	filter := " WHERE 1=1"
	offset := " OFFSET 0"
	limit := " LIMIT 10"

	// Implement search on userRole_name only
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
	cQ := `SELECT count(1) FROM user_roles` + filter
	cQ, arr = helper.ReplaceQueryParams(cQ, params)
	err := u.db.QueryRow(ctx, cQ, arr...).Scan(
		&res.Count,
	)
	if err != nil {
		return res, err
	}

	// Main query for retrieving userRoles
	q := query + filter + offset + limit
	q, arr = helper.ReplaceQueryParams(q, params)
	rows, err := u.db.Query(ctx, q, arr...)
	if err != nil {
		return res, err
	}
	defer rows.Close()

	for rows.Next() {
		obj := &models.UserRole{}
		err = rows.Scan(
			&obj.UserId,
			&obj.RoleId,
		)
		if err != nil {
			return res, err
		}

		res.UserRoles = append(res.UserRoles, obj)
	}

	return res, nil
}

// Update updates a userRole in the userRoles table.
func (u *userRoleRepo) Update(ctx context.Context, req *models.UpdateUserRole) (int64, error) {
	query := `UPDATE user_roles SET
		role_id = :role_id
	WHERE
		user_id = :user_id`

	params := map[string]interface{}{
		"role_id":        req.RoleId,
		"user_id": req.UserId,
	}

	q, arr := helper.ReplaceQueryParams(query, params)
	result, err := u.db.Exec(ctx, q, arr...)
	if err != nil {
		return 0, err
	}

	rowsAffected := result.RowsAffected()

	return rowsAffected, err
}

// Delete deletes a userRole from the userRoles table by its ID.
func (u *userRoleRepo) Delete(ctx context.Context, req *models.PrimaryKey) (int64, error) {
	query := `DELETE FROM user_roles WHERE id = $1`

	result, err := u.db.Exec(ctx, query, req.Id)
	if err != nil {
		return 0, err
	}

	rowsAffected := result.RowsAffected()

	return rowsAffected, err
}
