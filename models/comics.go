package models

import "time"

type PrimaryKeyUUID struct {
	Id string `json:"id"`
}

type Comics struct {
	Id          string    `json:"id"`
	Title   string    `json:"first_name"`
	Author    string    `json:"last_name"`
	Description string    `json:"phone_number"`
	Genre    string    `json:"password"`
	ReleaseDate string `json:"release_date"`
	PopularityScore string `json:"popularity_score"`
	PosterUrl string `json:"poster_url"`
	Price string `json:"price"`
	IsActive string `json:"is_active"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time    `json:"updated_at"`
}


type CreateComics struct {
	Title   string    `json:"first_name"`
	Author    string    `json:"last_name"`
	Description string    `json:"phone_number"`
	Genre    string    `json:"password"`
	ReleaseDate string `json:"release_date"`
	PopularityScore string `json:"popularity_score"`
	PosterUrl string `json:"poster_url"`
	Price string `json:"price"`
	IsActive string `json:"is_active"`
}

type UpdateComics struct {
	Id          string    `json:"id"`
	Title   string    `json:"first_name"`
	Author    string    `json:"last_name"`
	Description string    `json:"phone_number"`
	Genre    string    `json:"password"`
	ReleaseDate string `json:"release_date"`
	PopularityScore string `json:"popularity_score"`
	PosterUrl string `json:"poster_url"`
	Price string `json:"price"`
	IsActive string `json:"is_active"`
}

type GetListComicsRequest struct {
	Offset int    `json:"offset"`
	Limit  int    `json:"limit"`
	Search string `json:"search"`
}

type GetListComicsResponse struct {
	Count int     `json:"count"`
	Comics []*Comics `json:"comics"`
}
