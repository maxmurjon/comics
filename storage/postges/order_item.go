package postges

import (
	"comics/models"
	"comics/pkg/helper/helper"
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

type orderItemRepo struct {
	db *pgxpool.Pool
}

func (u *orderItemRepo) Create(ctx context.Context, req *models.CreateOrderItem) (*models.PrimaryKey, error) {
	query := `INSERT INTO order_items(
		order_id,
		comic_id,
		quantity,
		price
	) VALUES ($1, $2, $3, $4)`

	_,err := u.db.Exec(ctx, query,
		req.OrderId,
		req.ComicId,
		req.Quantity,
		req.Price,
	)

	pKey := &models.PrimaryKey{
		Id: req.OrderId,
	}

	return pKey, err

}

func (u *orderItemRepo) GetByID(ctx context.Context, req *models.PrimaryKey) (*models.OrderItem, error) {
	res := &models.OrderItem{}
	query := `
        SELECT
			order_id,
			comic_id,
			quantity,
			price
        FROM
            "order_items"
        WHERE
            order_id = $1`

	err := u.db.QueryRow(ctx, query, req.Id).Scan(
		&res.OrderId,
		&res.ComicId,
		&res.Quantity,
		&res.Price,
	)
	if err != nil {
		return res, err
	}

	return res, nil
}

func (u *orderItemRepo) GetList(ctx context.Context, req *models.GetListOrderItemRequest) (*models.GetListOrderItemResponse, error) {
	res := &models.GetListOrderItemResponse{}
	params := make(map[string]interface{})
	var arr []interface{}
	query := `SELECT
			order_id,
			comic_id,
			quantity,
			price
		FROM
			"order_items"`
	filter := " WHERE 1=1"
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

	cQ := `SELECT count(1) FROM "order_items"` + filter
	cQ, arr = helper.ReplaceQueryParams(cQ, params)
	err := u.db.QueryRow(ctx, cQ, arr...).Scan(
		&res.Count,
	)
	if err != nil {
		return res, err
	}

	q := query + filter + offset + limit

	q, arr = helper.ReplaceQueryParams(q, params)
	rows, err := u.db.Query(ctx, q, arr...)
	if err != nil {
		return res, err
	}
	defer rows.Close()

	for rows.Next() {
		obj := &models.OrderItem{}

		err = rows.Scan(
			&obj.OrderId,
			&obj.ComicId,
			&obj.Quantity,
			&obj.Price,
		)

		if err != nil {
			return res, err
		}

		res.OrderItem = append(res.OrderItem, obj)
	}

	return res, nil
}

func (u *orderItemRepo) Update(ctx context.Context, req *models.UpdateOrderItem) (id int64, err error) {
	query := `UPDATE "order_items" SET
		order_id=:order_id,
		comic_id=:comic_id,
		quantity=:quantity,
		price=:price
	WHERE
		order_id = :order_id`

	params := map[string]interface{}{
		"order_id":     req.OrderId,
		"comic_id":  req.ComicId,
		"quantity": req.Quantity,
		"price":      req.Price,
	}

	q, arr := helper.ReplaceQueryParams(query, params)
	result, err := u.db.Exec(ctx, q, arr...)
	if err != nil {
		return 0, err
	}

	rowsAffected := result.RowsAffected()

	return rowsAffected, err
}

func (u *orderItemRepo) Delete(ctx context.Context, req *models.PrimaryKey) (id int64, err error) {
	query := `DELETE FROM "order_items" WHERE order_id = $1`

	result, err := u.db.Exec(ctx, query, req.Id)
	if err != nil {
		return 0, err
	}

	rowsAffected := result.RowsAffected()

	return rowsAffected, err
}
