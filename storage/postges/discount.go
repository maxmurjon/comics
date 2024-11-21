package postges

import (
	"comics/models"
	"comics/pkg/helper/helper"
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

type discountRepo struct {
	db *pgxpool.Pool
}

func (u *discountRepo) Create(ctx context.Context, req *models.CreateDiscount) (*models.PrimaryKey, error) {
	var id int

	query := `INSERT INTO discount(
		id,
		code,
		discount_percentage,
		valid_until
	) VALUES ($1, $2, $3, $4) RETURNING id`

	err := u.db.QueryRow(ctx, query,
		req.Code,
		req.DiscountPersentage,
		req.ValidUntil,
	).Scan(&id)

	pKey := &models.PrimaryKey{
		Id: id,
	}

	return pKey, err

}

func (u *discountRepo) GetByID(ctx context.Context, req *models.PrimaryKey) (*models.Discount, error) {
	res := &models.Discount{}
	query := `
        SELECT
        id,
		code,
		discount_percentage,
		valid_until
        FROM
            "discount"
        WHERE
            id = $1`

	err := u.db.QueryRow(ctx, query, req.Id).Scan(
		&res.Id,
		&res.Code,
		&res.DiscountPersentage,
		&res.ValidUntil,
	)
	if err != nil {
		return res, err
	}

	return res, nil
}

func (u *discountRepo) GetList(ctx context.Context, req *models.GetListDiscountRequest) (*models.GetListDiscountResponse, error) {
	res := &models.GetListDiscountResponse{}
	params := make(map[string]interface{})
	var arr []interface{}
	query := `SELECT
		id,
		code,
		discount_percentage,
		valid_until 
	FROM
		"discount"`
	filter := " WHERE 1=1"
	discount := " discount BY created_at"
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

	cQ := `SELECT count(1) FROM "discount"` + filter
	cQ, arr = helper.ReplaceQueryParams(cQ, params)
	err := u.db.QueryRow(ctx, cQ, arr...).Scan(
		&res.Count,
	)
	if err != nil {
		return res, err
	}

	q := query + filter + discount + arrangement + offset + limit

	q, arr = helper.ReplaceQueryParams(q, params)
	rows, err := u.db.Query(ctx, q, arr...)
	if err != nil {
		return res, err
	}
	defer rows.Close()

	for rows.Next() {
		obj := &models.Discount{}

		err = rows.Scan(
			&obj.Id,
			&obj.Code,
			&obj.DiscountPersentage,
			&obj.ValidUntil,
		)

		if err != nil {
			return res, err
		}

		res.Discount = append(res.Discount, obj)
	}

	return res, nil
}

func (u *discountRepo) Update(ctx context.Context, req *models.UpdateDiscount) (id int64, err error) {
	query := `UPDATE "discounts" SET
		code=:code,
		discount_persentage=:discount_persentage,
		ValidUntil=:ValidUntil
	WHERE
		id = :id`

	params := map[string]interface{}{
		"id":          req.Id,
		"user_id":     req.Code,
		"discount_date":  req.DiscountPersentage,
		"total_price": req.ValidUntil,
	}

	q, arr := helper.ReplaceQueryParams(query, params)
	result, err := u.db.Exec(ctx, q, arr...)
	if err != nil {
		return 0, err
	}

	rowsAffected := result.RowsAffected()

	return rowsAffected, err
}

func (u *discountRepo) Delete(ctx context.Context, req *models.PrimaryKey) (id int64, err error) {
	query := `DELETE FROM "discount" WHERE id = $1`

	result, err := u.db.Exec(ctx, query, req.Id)
	if err != nil {
		return 0, err
	}

	rowsAffected := result.RowsAffected()

	return rowsAffected, err
}
