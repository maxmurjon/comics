package models

import "time"

type PrimaryKeyUUID struct {
	Id string `json:"id"`
}

type Comics struct {
	Id          string    `json:"id"`
	Title   string    `json:"title"`
	Author    string    `json:"author"`
	Description string    `json:"description"`
	Genre    string    `json:"genre"`
	ReleaseDate time.Time `json:"release_date"`
	PopularityScore int `json:"popularity_score"`
	PosterUrl string `json:"poster_url"`
	Price float32 `json:"price"`
	IsActive bool `json:"is_active"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time    `json:"updated_at"`
}


type CreateComics struct {
	Title   string    `json:"title"`
	Author    string    `json:"author"`
	Description string    `json:"description"`
	Genre    string    `json:"genre"`
	ReleaseDate time.Time `json:"release_date"`
	PopularityScore int `json:"popularity_score"`
	PosterUrl string `json:"poster_url"`
	Price float32 `json:"price"`
	IsActive bool `json:"is_active"`
}

type UpdateComics struct {
	Id          int    `json:"id"`
	Title   string    `json:"title"`
	Author    string    `json:"author"`
	Description string    `json:"description"`
	Genre    string    `json:"genre"`
	ReleaseDate string `json:"release_date"`
	PopularityScore int `json:"popularity_score"`
	PosterUrl string `json:"poster_url"`
	Price float32 `json:"price"`
	IsActive bool `json:"is_active"`
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
