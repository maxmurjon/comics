package models

import "time"

type ProductImage struct {
	ID        int       `json:"id"`
	ProductID int       `json:"product_id"`
	ImageUrl  string    `json:"image_url"`
	IsPrimary bool      `json:"is_primary"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type CreateProductImage struct {
	ProductID int    `json:"product_id"`
	ImageUrl  string `json:"image_url"`
	IsPrimary bool   `json:"is_primary"`
}

type UpdateProductImage struct {
	ID        int    `json:"id"`
	ProductID int    `json:"product_id"`
	ImageUrl  string `json:"image_url"`
	IsPrimary bool   `json:"is_primary"`
}

type GetListProductImageRequest struct {
	Offset int    `json:"offset"`
	Limit  int    `json:"limit"`
	Search string `json:"search"`
}

type GetListProductImageResponse struct {
	Count         int             `json:"count"`
	ProductImages []*ProductImage `json:"product_image"`
}
