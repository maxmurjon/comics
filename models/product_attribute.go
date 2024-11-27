package models

import "time"

type ProductAttribute struct {
	ID          int       `json:"id"`
	ProductID   string    `json:"product_id"`
	AttributeID int       `json:"attribute_id"`
	Value       string    `json:"value"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt	time.Time `json:"updated_at"`
}

type CreateProductAttribute struct {
	ProductID   string    `json:"product_id"`
	AttributeID int       `json:"attribute_id"`
	Value       string    `json:"value"`
}

type UpdateProductAttribute struct {
	ID          int    `json:"id"`
	ProductID   string    `json:"product_id"`
	AttributeID int       `json:"attribute_id"`
	Value       string    `json:"value"`
}

type GetListProductAttributeRequest struct {
	Offset int    `json:"offset"`
	Limit  int    `json:"limit"`
	Search string `json:"search"`
}

type GetListProductAttributeResponse struct {
	Count             int                 `json:"count"`
	ProductAttributes []*ProductAttribute `json:"product_attributes"`
}
