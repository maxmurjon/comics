package models

import "time"

type Order struct {
	Id         int       `json:"id"`
	UserId     string    `json:"user_id"`
	OrderDate  string    `json:"order_date"`
	TotalPrice string    `json:"total_price"`
	Status     string    `json:"status"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

type CreateOrder struct {
	UserId     string `json:"user_id"`
	OrderDate  string `json:"order_date"`
	TotalPrice string `json:"total_price"`
	Status     string `json:"status"`
}

type UpdateOrder struct {
	Id         int       `json:"id"`
	UserId     string    `json:"user_id"`
	OrderDate  string    `json:"order_date"`
	TotalPrice string    `json:"total_price"`
	Status     string    `json:"status"`
}

type GetListOrderRequest struct {
	Offset int    `json:"offset"`
	Limit  int    `json:"limit"`
	Search string `json:"search"`
}

type GetListOrderResponse struct {
	Count  int       `json:"count"`
	Order []*Order `json:"order"`
}
