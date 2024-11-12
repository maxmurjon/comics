package postges

import (
	"comics/models"
	"comics/pkg/helper/helper"
	"context"

	"github.com/google/uuid"

	"github.com/jackc/pgx/v5/pgxpool"
)

type comicsRepo struct {
	db *pgxpool.Pool
}

func (u *comicsRepo) Create(ctx context.Context, req *models.CreateComics) (*models.PrimaryKeyUUID, error) {

	uuid, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}

	query := `INSERT INTO comics(
		id,
		title,
		author,
		description,
		genre,
		release_date,
		popularity_score,
		poster_url,
		price,
		is_active,
		created_at,
		updated_at 
	) VALUES ($1, $2, $3, $4, $5, $6,$7,$8,$9,$10 now(), now()	)`

	_, err = u.db.Exec(ctx, query,
		uuid.String(),
		req.Title,
		req.Author,
		req.Description,
		req.Genre,
		req.ReleaseDate,
		req.PopularityScore,
		req.PosterUrl,
		req.Price,
		req.IsActive,
	)

	pKey := &models.PrimaryKeyUUID{
		Id: uuid.String(),
	}

	return pKey, err

}

func (u *comicsRepo) GetByID(ctx context.Context, req *models.PrimaryKeyUUID) (*models.Comics, error) {
	res := &models.Comics{}
	query := `
        SELECT
            id,
		title,
		author,
		description,
		genre,
		release_date,
		popularity_score,
		poster_url,
		price,
		is_active,
		created_at,
		updated_at
        FROM
            "comics"
        WHERE
            id = $1`

	err := u.db.QueryRow(ctx, query, req.Id).Scan(
		&res.Id,
		&res.Title,
		&res.Author,
		&res.Description,
		&res.Genre,
		&res.ReleaseDate,
		&res.PopularityScore,
		&res.PosterUrl,
		&res.Price,
		&res.IsActive,
		&res.CreatedAt, // created_at as time.Time
		&res.UpdatedAt, // updated_at as time.Time
	)
	if err != nil {
		return res, err
	}

	return res, nil
}

func (u *comicsRepo) GetList(ctx context.Context, req *models.GetListComicsRequest) (*models.GetListComicsResponse, error) {
	res := &models.GetListComicsResponse{}
	params := make(map[string]interface{})
	var arr []interface{}
	query := `SELECT
		   id,
		title,
		author,
		description,
		genre,
		release_date,
		popularity_score,
		poster_url,
		price,
		is_active,
		created_at,
		updated_at
	FROM
		"comics"`
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

	cQ := `SELECT count(1) FROM "comics"` + filter
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
		obj := &models.Comics{}

		err = rows.Scan(
			&obj.Id,
			&obj.Title,
			&obj.Author,
			&obj.Description,
			&obj.Genre,
			&obj.ReleaseDate,
			&obj.PopularityScore,
			&obj.PosterUrl,
			&obj.Price,
			&obj.IsActive,
			&obj.CreatedAt, // created_at as time.Time
			&obj.UpdatedAt, // updated_at as time.Time
		)

		if err != nil {
			return res, err
		}

		res.Comics = append(res.Comics, obj)
	}

	return res, nil
}

func (u *comicsRepo) Update(ctx context.Context, req *models.UpdateComics) (id int64, err error) {
	query := `UPDATE "comicss" SET
		title =:title,
		author =:author,
		description =:description,
		genre =:genre,
		release_date =:release_date,
		popularity_score =:popularity_score,
		poster_url =:poster_url,
		price =:price,
		is_active =:is_active,
		updated_at=now()
	WHERE
		id = :id`

	params := map[string]interface{}{
		"id":            req.Id,
		"title": req.Title,
		"author": req.Author,
		"description": req.Description,
		"genre": req.Genre,
		"release_date": req.ReleaseDate,
		"popularity_score": req.PopularityScore,
		"poster_url": req.PosterUrl,
		"price": req.Price,
		"is_active": req.IsActive,
	}

	q, arr := helper.ReplaceQueryParams(query, params)
	result, err := u.db.Exec(ctx, q, arr...)
	if err != nil {
		return 0, err
	}

	rowsAffected := result.RowsAffected()

	return rowsAffected, err
}

func (u *comicsRepo) Delete(ctx context.Context, req *models.PrimaryKeyUUID) (id int64, err error) {
	query := `DELETE FROM "comics" WHERE id = $1`

	result, err := u.db.Exec(ctx, query, req.Id)
	if err != nil {
		return 0, err
	}

	rowsAffected := result.RowsAffected()

	return rowsAffected, err
}
