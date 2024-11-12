package models

type OrderItem struct {
	OrderId int `json:"order_id"`
	ComicId int `json:"comic_id"`
	Quantity int `json:"quantity"`
	Price float64 `json:"price"`
}

type CreateOrderItem struct {
	OrderId int `json:"order_id"`
	ComicId int `json:"comic_id"`
	Quantity int `json:"quantity"`
	Price float64 `json:"price"`
}

type UpdateOrderItem struct {
	OrderId int `json:"order_id"`
	ComicId int `json:"comic_id"`
	Quantity int `json:"quantity"`
	Price float64 `json:"price"`
}

type GetListOrderItemRequest struct {
	Offset int    `json:"offset"`
	Limit  int    `json:"limit"`
	Search string `json:"search"`
}

type GetListOrderItemResponse struct {
	Count     int          `json:"count"`
	OrderItem []*OrderItem `json:"OrderItem"`
}
