package models

type Discount struct {
	Id   int    `json:"id"`
	Code string `json:"code"`
	DiscountPersentage string `json:"discount_persentage"`
	ValidUntil string `json:"valid_until"`
}

type CreateDiscount struct {
	Code string `json:"code"`
	DiscountPersentage string `json:"discount_persentage"`
	ValidUntil string `json:"valid_until"`
}

type UpdateDiscount struct {
	Id           int    `json:"id"`
	Code string `json:"code"`
	DiscountPersentage string `json:"discount_persentage"`
	ValidUntil string `json:"valid_until"`
}

type GetListDiscountRequest struct {
	Offset int    `json:"offset"`
	Limit  int    `json:"limit"`
	Search string `json:"search"`
}

type GetListDiscountResponse struct {
	Count    int         `json:"count"`
	Discount []*Discount `json:"Discount"`
}
