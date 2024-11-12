package postges

import (
	"comics/models"
	"comics/pkg/helper/helper"
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

type orderRepo struct {
	db *pgxpool.Pool
}

func (u *orderRepo) Create(ctx context.Context, req *models.CreateOrder) (*models.PrimaryKey, error) {
	var id int

	query := `INSERT INTO order(
		user_id,
		order_date,
		total_price,
		status,
		created_at,
		updated_at 
	) VALUES ($1, $2, $3, $4,now(), now() RETURNING id`

	err := u.db.QueryRow(ctx, query,
		req.UserId,
		req.OrderDate,
		req.TotalPrice,
		req.Status,
	).Scan(&id)

	pKey := &models.PrimaryKey{
		Id: id,
	}

	return pKey, err

}

func (u *orderRepo) GetByID(ctx context.Context, req *models.PrimaryKey) (*models.Order, error) {
	res := &models.Order{}
	query := `
        SELECT
            id,
		user_id,
		order_date,
		total_price,
		status,
		created_at,
		updated_at 
        FROM
            "order"
        WHERE
            id = $1`

	err := u.db.QueryRow(ctx, query, req.Id).Scan(
		&res.Id,
		&res.UserId,
		&res.OrderDate,
		&res.TotalPrice,
		&res.Status,
		&res.CreatedAt,
		&res.UpdatedAt,
	)
	if err != nil {
		return res, err
	}

	return res, nil
}

func (u *orderRepo) GetList(ctx context.Context, req *models.GetListOrderRequest) (*models.GetListOrderResponse, error) {
	res := &models.GetListOrderResponse{}
	params := make(map[string]interface{})
	var arr []interface{}
	query := `SELECT
		   id,
		user_id,
		order_date,
		total_price,
		status,
		created_at,
		updated_at 
	FROM
		"order"`
	filter := " WHERE 1=1"
	order := " ORDER BY created_at"
	arrangement := " DESC"
	offset := " OFFSET 0"
	limit := " LIMIT 10"

	if len(req.Search) > 0 {
		params["search"] = req.Search
		filter += " AND ((name || phone || is_active || login) ILIKE ('%' || :search || '%'))"
	}

	if req.Offset > 0 {
		params["offset"] = req.Offset
		offset = " OFFSET :offset"
	}

	if req.Limit > 0 {
		params["limit"] = req.Limit
		limit = " LIMIT :limit"
	}

	cQ := `SELECT count(1) FROM "order"` + filter
	cQ, arr = helper.ReplaceQueryParams(cQ, params)
	err := u.db.QueryRow(ctx, cQ, arr...).Scan(
		&res.Count,
	)
	if err != nil {
		return res, err
	}

	q := query + filter + order + arrangement + offset + limit

	q, arr = helper.ReplaceQueryParams(q, params)
	rows, err := u.db.Query(ctx, q, arr...)
	if err != nil {
		return res, err
	}
	defer rows.Close()

	for rows.Next() {
		obj := &models.Order{}

		err = rows.Scan(
			&obj.Id,
			&obj.UserId,
			&obj.OrderDate,
			&obj.TotalPrice,
			&obj.Status,
			&obj.CreatedAt, // created_at as time.Time
			&obj.UpdatedAt, // updated_at as time.Time
		)

		if err != nil {
			return res, err
		}

		res.Order = append(res.Order, obj)
	}

	return res, nil
}

func (u *orderRepo) Update(ctx context.Context, req *models.UpdateOrder) (id int64, err error) {
	query := `UPDATE "orders" SET
		user_id=:user_id,
		order_date=:order_date,
		total_price=:total_price,
		status=:status,
		updated_at=now()
	WHERE
		id = :id`

	params := map[string]interface{}{
		"id":          req.Id,
		"user_id":     req.UserId,
		"order_date":  req.OrderDate,
		"total_price": req.TotalPrice,
		"status":      req.Status,
	}

	q, arr := helper.ReplaceQueryParams(query, params)
	result, err := u.db.Exec(ctx, q, arr...)
	if err != nil {
		return 0, err
	}

	rowsAffected := result.RowsAffected()

	return rowsAffected, err
}

func (u *orderRepo) Delete(ctx context.Context, req *models.PrimaryKey) (id int64, err error) {
	query := `DELETE FROM "order" WHERE id = $1`

	result, err := u.db.Exec(ctx, query, req.Id)
	if err != nil {
		return 0, err
	}

	rowsAffected := result.RowsAffected()

	return rowsAffected, err
}
