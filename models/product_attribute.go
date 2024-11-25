package models

type ProductAttribute struct {
	ID             int `json:"id"`
	ProductID string `json:"product_id"`
	Key string `json:"key"`
	Value string `json:"value"`
}

type CreateProductAttribute struct {
	ProductID string `json:"product_id"`
	Key string `json:"key"`
	Value string `json:"value"`
}

type UpdateProductAttribute struct {
	ID             int `json:"id"`
	ProductID string `json:"product_id"`
	Key string `json:"key"`
	Value string `json:"value"`
}

type GetListProductAttributeRequest struct {
	Offset int    `json:"offset"`
	Limit  int    `json:"limit"`
	Search string `json:"search"`
}

type GetListProductAttributeResponse struct {
	Count       int           `json:"count"`
	ProductAttributes []*ProductAttribute `json:"ProductAttributes"`
}
