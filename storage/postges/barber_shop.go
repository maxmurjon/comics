package postges

// import (
// 	"comics/config"
// 	"comics/models"
// 	"comics/pkg/helper/helper"
// 	"context"
// 	"fmt"

// 	"github.com/google/uuid"
// 	"github.com/jackc/pgx/v5/pgxpool"
// )

// type barberShopRepo struct {
// 	db *pgxpool.Pool
// }

// func (u *barberShopRepo) Create(ctx context.Context, req *models.CreateBarberShop) (*models.BarberShopPrimaryKey, error) {
// 	query := `
// 		INSERT INTO barber_shops (
// 			id,
// 			name,
// 			address,
// 			latitude,
// 			longitude,
// 			created_at,
// 			updated_at
// 		) VALUES (
// 			$1, $2, $3, $4, $5, now(), now()
// 		)
// 		RETURNING id;
// 	`

// 	fmt.Println(req)
// 	uuid, err := uuid.NewRandom()
// 	if err != nil {
// 		return nil, err
// 	}

// 	_, err = u.db.Exec(ctx, query,
// 		uuid.String(),
// 		req.Name,
// 		req.Address,
// 		req.Latitude,  // latitude to'g'ri
// 		req.Longitute, // longitude to'g'ri
// 	)
// 	if err != nil {
// 		return nil, err
// 	}

// 	pKey := &models.BarberShopPrimaryKey{
// 		Id: uuid.String(),
// 	}

// 	return pKey, nil
// }

// // GetByID retrieves a BarberShop by its ID.
// func (u *barberShopRepo) GetByID(ctx context.Context, req *models.BarberShopPrimaryKey) (*models.BarberShop, error) {
// 	res := &models.BarberShop{}
// 	query := `SELECT
// 		id,
// 			name,
// 			address,
// 			latitude,
// 			longitude,
// 			TO_CHAR(created_at, ` + config.DatabaseQueryTimeLayout + `) AS created_at,
// 			TO_CHAR(updated_at, ` + config.DatabaseQueryTimeLayout + `) AS updated_at
// 	FROM
// 		barber_shops
// 	WHERE
// 		id = $1`

// 	err := u.db.QueryRow(ctx, query, req.Id).Scan(
// 		&res.Id,
// 		&res.Name,
// 		&res.Address,
// 		&res.Latitude,
// 		&res.Longitute,
// 		&res.CreatedAt,
// 		&res.UpdatedAt,
// 	)
// 	if err != nil {
// 		return res, err
// 	}

// 	return res, nil
// }

// // GetList retrieves a list of BarberShops with pagination and optional search functionality.
// func (u *barberShopRepo) GetList(ctx context.Context, req *models.GetListBarberShopRequest) (*models.GetListBarberShopResponse, error) {
// 	res := &models.GetListBarberShopResponse{}
// 	params := make(map[string]interface{})
// 	var arr []interface{}

// 	query := `SELECT
// 		id,
// 			name,
// 			address,
// 			latitude,
// 			longitude,
// 			TO_CHAR(created_at, ` + config.DatabaseQueryTimeLayout + `) AS created_at,
// 			TO_CHAR(updated_at, ` + config.DatabaseQueryTimeLayout + `) AS updated_at
// 	FROM
// 		barber_shops`
// 	filter := " WHERE 1=1"
// 	offset := " OFFSET 0"
// 	limit := " LIMIT 10"

// 	// Implement search on BarberShop_name only
// 	if len(req.Search) > 0 {
// 		params["search"] = req.Search
// 		filter += " AND BarberShop_name ILIKE '%' || :search || '%'"
// 	}

// 	if req.Offset > 0 {
// 		params["offset"] = req.Offset
// 		offset = " OFFSET :offset"
// 	}

// 	if req.Limit > 0 {
// 		params["limit"] = req.Limit
// 		limit = " LIMIT :limit"
// 	}

// 	// Count query
// 	cQ := `SELECT count(1) FROM barber_shops` + filter
// 	cQ, arr = helper.ReplaceQueryParams(cQ, params)
// 	err := u.db.QueryRow(ctx, cQ, arr...).Scan(
// 		&res.Count,
// 	)
// 	if err != nil {
// 		return res, err
// 	}

// 	// Main query for retrieving BarberShops
// 	q := query + filter + offset + limit
// 	q, arr = helper.ReplaceQueryParams(q, params)
// 	rows, err := u.db.Query(ctx, q, arr...)
// 	if err != nil {
// 		return res, err
// 	}
// 	defer rows.Close()

// 	for rows.Next() {
// 		obj := &models.BarberShop{}
// 		err = rows.Scan(
// 			&obj.Id,
// 			&obj.Name,
// 			&obj.Address,
// 			&obj.Latitude,
// 			&obj.Longitute,
// 			&obj.CreatedAt,
// 			&obj.UpdatedAt,
// 		)
// 		if err != nil {
// 			return res, err
// 		}

// 		res.BarberShops = append(res.BarberShops, obj)
// 	}

// 	return res, nil
// }

// // Update updates a BarberShop in the BarberShops table.
// func (u *barberShopRepo) Update(ctx context.Context, req *models.UpdateBarberShop) (int64, error) {
// 	query := `UPDATE barber_shops SET
// 		name = :name,
// 		address = :address,
// 		lattitude = :lattitude,
// 		longitude = :longitude,
// 		updated_at = now()
// 	WHERE
// 		id = :id`

// 	params := map[string]interface{}{
// 		"id":        req.Id,
// 		"name":      req.Name,
// 		"address":   req.Address,
// 		"lattitude": req.Latitude,
// 		"longitude": req.Longitute,
// 	}

// 	q, arr := helper.ReplaceQueryParams(query, params)
// 	result, err := u.db.Exec(ctx, q, arr...)
// 	if err != nil {
// 		return 0, err
// 	}

// 	rowsAffected := result.RowsAffected()

// 	return rowsAffected, err
// }

// // Delete deletes a BarberShop from the BarberShops table by its ID.
// func (u *barberShopRepo) Delete(ctx context.Context, req *models.BarberShopPrimaryKey) (int64, error) {
// 	query := `DELETE FROM barber_shops WHERE id = $1`

// 	result, err := u.db.Exec(ctx, query, req.Id)
// 	if err != nil {
// 		return 0, err
// 	}

// 	rowsAffected := result.RowsAffected()

// 	return rowsAffected, err
// }
