package postges

import (
	"comics/models"
	"comics/pkg/helper/helper"
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

type paymentRepo struct {
	db *pgxpool.Pool
}

func (u *paymentRepo) Create(ctx context.Context, req *models.CreatePayment) (*models.PrimaryKey, error) {
	var id int

	query := `INSERT INTO payments(
		purchase_id,
		amount,
		payment_method,
		payment_date,
		status
	) VALUES ($1, $2, $3, $4, $5) RETURNING id`

	err := u.db.QueryRow(ctx, query,
		req.PurchaseId,
		req.PaymentMethod,
		req.PaymentDate,
		req.Amount,
		req.Status,
	).Scan(&id)

	pKey := &models.PrimaryKey{
		Id: id,
	}

	return pKey, err

}

func (u *paymentRepo) GetByID(ctx context.Context, req *models.PrimaryKey) (*models.Payment, error) {
	res := &models.Payment{}
	query := `
        SELECT
            id,
            purchase_id,
		amount,
		payment_method,
		payment_date,
		status
        FROM
            "payments"
        WHERE
            id = $1`

	err := u.db.QueryRow(ctx, query, req.Id).Scan(
		&res.Id,
		&res.PurchaseId,
		&res.Amount,
		&res.PaymentMethod,
		&res.PaymentDate,
	)
	if err != nil {
		return res, err
	}

	return res, nil
}

func (u *paymentRepo) GetList(ctx context.Context, req *models.GetListPaymentRequest) (*models.GetListPaymentResponse, error) {
	res := &models.GetListPaymentResponse{}
	params := make(map[string]interface{})
	var arr []interface{}
	query := `SELECT
		id,
		purchase_id,
		amount,
		payment_method,
		payment_date,
		status
	FROM
		"payments"`
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

	cQ := `SELECT count(1) FROM "payments"` + filter
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
		obj := &models.Payment{}

		err = rows.Scan(
			&obj.Id,
			&obj.PurchaseId,
			&obj.Amount,
			&obj.PaymentMethod,
			&obj.PaymentDate,
			&obj.Status,
		)

		if err != nil {
			return res, err
		}

		res.Payment = append(res.Payment, obj)
	}

	return res, nil
}

func (u *paymentRepo) Update(ctx context.Context, req *models.UpdatePayment) (id int64, err error) {
	query := `UPDATE "payments" SET
		id = :id,
		purchase_id = :purchase_id,
		amount = :amount,
		payment_method = :payment_method,
		payment_date = :payment_date,
		status = :status,
	WHERE
		id = :id`

	params := map[string]interface{}{
		"id": req.Id,
		"purchase_id": req.PurchaseId,
		"amount":req.Amount,
		"payment_method":req.PaymentMethod,
		"payment_date":req.PaymentDate,
		"status":req.Status,
	}

	q, arr := helper.ReplaceQueryParams(query, params)
	result, err := u.db.Exec(ctx, q, arr...)
	if err != nil {
		return 0, err
	}

	rowsAffected := result.RowsAffected()

	return rowsAffected, err
}

func (u *paymentRepo) Delete(ctx context.Context, req *models.PrimaryKey) (id int64, err error) {
	query := `DELETE FROM "payments" WHERE id = $1`

	result, err := u.db.Exec(ctx, query, req.Id)
	if err != nil {
		return 0, err
	}

	rowsAffected := result.RowsAffected()

	return rowsAffected, err
}
