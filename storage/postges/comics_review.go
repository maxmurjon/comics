package postges

import (
	"comics/models"
	"comics/pkg/helper/helper"
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

type comicsReviewRepo struct {
	db *pgxpool.Pool
}

func (u *comicsReviewRepo) Create(ctx context.Context, req *models.CreateComicsReview) (*models.PrimaryKey, error) {
	var id int

	query := `INSERT INTO comics_review(
		comic_id,
		user_id,
		rating,
		review,
		created_at,
		updated_at 
	) VALUES ($1, $2, $3, $4, now(), now()) RETURNING id`

	err := u.db.QueryRow(ctx, query,
		req.ComicId,
		req.UserId,
		req.Rating,
		req.Review,
	).Scan(&id)

	pKey := &models.PrimaryKey{
		Id: id,
	}

	return pKey, err
}

func (u *comicsReviewRepo) GetByID(ctx context.Context, req *models.PrimaryKey) (*models.ComicsReview, error) {
	res := &models.ComicsReview{}
	query := `
        SELECT
            id,
		comic_id,
		user_id,
		rating,
		review,
		created_at,
		updated_at 
        FROM
            "comics_review"
        WHERE
            id = $1`

	err := u.db.QueryRow(ctx, query, req.Id).Scan(
		&res.Id,
		&res.ComicId,
		&res.UserId,
		&res.Rating,
		&res.Review,
		&res.CreatedAt, // created_at as time.Time
		&res.UpdatedAt, // updated_at as time.Time
	)
	if err != nil {
		return res, err
	}

	return res, nil
}

func (u *comicsReviewRepo) GetList(ctx context.Context, req *models.GetListComicsReviewRequest) (*models.GetListComicsReviewResponse, error) {
	res := &models.GetListComicsReviewResponse{}
	params := make(map[string]interface{})
	var arr []interface{}
	query := `SELECT
		    id,
		comic_id,
		user_id,
		rating,
		review,
		created_at,
		updated_at 
	FROM
		"comics_review"`
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

	cQ := `SELECT count(1) FROM "comics_review"` + filter
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
		obj := &models.ComicsReview{}

		err = rows.Scan(
			&obj.Id,
			&obj.ComicId,
			&obj.UserId,
			&obj.Rating,
			&obj.Review,
			&obj.CreatedAt,
			&obj.UpdatedAt,
		)

		if err != nil {
			return res, err
		}

		res.ComicsReview = append(res.ComicsReview, obj)
	}

	return res, nil
}

func (u *comicsReviewRepo) Update(ctx context.Context, req *models.UpdateComicsReview) (id int64, err error) {
	query := `UPDATE "comics_review" SET
		comic_id =:comic_id,
		user_id =:user_id,
		rating =:rating,
		review =:review,
		updated_at=now()
	WHERE
		id = :id`

	params := map[string]interface{}{
		"id":               req.Id,
		"comic_id":            req.ComicId,
		"user_id":           req.UserId,
		"rating":      req.Rating,
		"review":            req.Review,
	}

	q, arr := helper.ReplaceQueryParams(query, params)
	result, err := u.db.Exec(ctx, q, arr...)
	if err != nil {
		return 0, err
	}

	rowsAffected := result.RowsAffected()

	return rowsAffected, err
}

func (u *comicsReviewRepo) Delete(ctx context.Context, req *models.PrimaryKey) (id int64, err error) {
	query := `DELETE FROM "comics_review" WHERE id = $1`

	result, err := u.db.Exec(ctx, query, req.Id)
	if err != nil {
		return 0, err
	}

	rowsAffected := result.RowsAffected()

	return rowsAffected, err
}
