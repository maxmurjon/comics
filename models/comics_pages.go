package models

import "time"

type ComicsPages struct {
	Id         string    `json:"id"`
	ComicId    int       `json:"comic_id"`
	PageNumber int       `json:"page_number"`
	PageUrl     int       `json:"page_url"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

type CreateComicsPages struct {
	ComicId    int       `json:"comic_id"`
	PageNumber int       `json:"page_number"`
	PageUrl     int       `json:"page_url"`
}

type UpdateComicsPages struct {
	Id        string    `json:"id"`
	ComicId    int       `json:"comic_id"`
	PageNumber int       `json:"page_number"`
	PageUrl     int       `json:"page_url"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

type GetListComicsPagesRequest struct {
	Offset int    `json:"offset"`
	Limit  int    `json:"limit"`
	Search string `json:"search"`
}

type GetListComicsPagesResponse struct {
	Count       int            `json:"count"`
	ComicsPages []*ComicsPages `json:"comics_pages"`
}
