package models

import "time"

type ComicsReview struct {
	Id        string    `json:"id"`
	ComicId   int       `json:"comic_id"`
	UserId    string    `json:"user_id"`
	Rating    int       `json:"rating"`
	Review    string    `json:"review"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type CreateComicsReview struct {
	ComicId int    `json:"comic_id"`
	UserId  string `json:"user_id"`
	Rating  int    `json:"rating"`
	Review  string `json:"review"`
}

type UpdateComicsReview struct {
	Id        string    `json:"id"`
	ComicId   int       `json:"comic_id"`
	UserId    string    `json:"user_id"`
	Rating    int       `json:"rating"`
	Review    string    `json:"review"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type GetListComicsReviewRequest struct {
	Offset int    `json:"offset"`
	Limit  int    `json:"limit"`
	Search string `json:"search"`
}

type GetListComicsReviewResponse struct {
	Count        int             `json:"count"`
	ComicsReview []*ComicsReview `json:"comicsReview"`
}
