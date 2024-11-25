package postges

import (
	"comics/models"
	"comics/pkg/helper/helper"
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

type comicsPagesRepo struct {
	db *pgxpool.Pool
}

func (u *comicsPagesRepo) Create(ctx context.Context, req *models.CreateComicsPages) (models.PrimaryKey, error) {
	var id int

	query := `INSERT INTO comic_pages(
		comic_id,
		page_number,
		page_url,
		created_at
	) VALUES ($1, $2, $3,  now()) RETURNING id`

	err := u.db.QueryRow(ctx, query,
		req.ComicId,
		req.PageNumber,
		req.PageUrl,
	).Scan(&id)

	pKey := &models.PrimaryKey{
		Id: id,
	}

	return *pKey, err
}

func (u *comicsPagesRepo) GetByID(ctx context.Context, req *models.PrimaryKey) (*models.ComicsPages, error) {
	res := &models.ComicsPages{}
	query := `
        SELECT
            id,
		comic_id,
		page_number,
		page_url,
		created_at
        FROM
            "comic_pages"
        WHERE
            id = $1`

	err := u.db.QueryRow(ctx, query, req.Id).Scan(
		&res.Id,
		&res.ComicId,
		&res.PageNumber,
		&res.PageUrl,
		&res.CreatedAt, // created_at as time.Time
	)
	if err != nil {
		return res, err
	}

	return res, nil
}

func (u *comicsPagesRepo) GetList(ctx context.Context, req *models.GetListComicsPagesRequest) (*models.GetListComicsPagesResponse, error) {
	res := &models.GetListComicsPagesResponse{}
	params := make(map[string]interface{})
	var arr []interface{}
	query := `SELECT
		    id,
		comic_id,
		page_number,
		page_url,
		created_at
	FROM
		"comic_pages"`
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

	cQ := `SELECT count(1) FROM "comics_Pages"` + filter
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
		obj := &models.ComicsPages{}

		err = rows.Scan(
			&obj.Id,
			&obj.ComicId,
			&obj.PageNumber,
			&obj.PageUrl,
			&obj.CreatedAt,
		)

		if err != nil {
			return res, err
		}

		res.ComicsPages = append(res.ComicsPages, obj)
	}

	return res, nil
}

func (u *comicsPagesRepo) Update(ctx context.Context, req *models.UpdateComicsPages) (id int64, err error) {
	query := `UPDATE "comic_Pages" SET
		comic_id =:comic_id,
		page_number =:page_number,
		page_url =:page_url
	WHERE
		id = :id`

	params := map[string]interface{}{
		"id":               req.Id,
		"comic_id":            req.ComicId,
		"page_number":           req.PageNumber,
		"page_url":      req.PageUrl,
	}

	q, arr := helper.ReplaceQueryParams(query, params)
	result, err := u.db.Exec(ctx, q, arr...)
	if err != nil {
		return 0, err
	}

	rowsAffected := result.RowsAffected()

	return rowsAffected, err
}

func (u *comicsPagesRepo) Delete(ctx context.Context, req *models.PrimaryKey) (id int64, err error) {
	query := `DELETE FROM "comic_pages" WHERE id = $1`

	result, err := u.db.Exec(ctx, query, req.Id)
	if err != nil {
		return 0, err
	}

	rowsAffected := result.RowsAffected()

	return rowsAffected, err
}
