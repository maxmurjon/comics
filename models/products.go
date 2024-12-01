package models

import "time"

type ProductInfo struct {
	ID            int           `json:"id"`
	Name          string        `json:"name"`
	TranslatedName string       `json:"translated_name,omitempty"`
	Description   string        `json:"description"`
	TranslatedDescription string `json:"translated_description,omitempty"`
	Price         float64       `json:"price"`
	StockQuantity int           `json:"stock_quantity"`
	ImageURLs     []string      `json:"image_urls"`
	Categories    []Category    `json:"categories"`
	Attributes    []Attribute   `json:"attributes"`
	AverageRating float64       `json:"average_rating"`
	ReviewsCount  int           `json:"reviews_count"`
	CreatedAt     time.Time     `json:"created_at"`
	UpdatedAt     time.Time     `json:"updated_at"`
}

type Product struct {
	ID            int       `json:"id"`
	Name          string    `json:"name"`
	Description   string    `json:"description"`
	Price         float64   `json:"price"`          // NUMERIC(10, 2) uchun float64 ishlatiladi
	StockQuantity int       `json:"stock_quantity"` // Ombordagi mahsulot miqdori
	CreatedAt     time.Time `json:"created_at"`     // Yaratilgan vaqt
	UpdatedAt     time.Time `json:"updated_at"`     // Yangilangan vaqt
}

type CreateProduct struct {
	Name          string  `json:"name"`
	Description   string  `json:"description"`
	Price         float64 `json:"price"`          // NUMERIC(10, 2) uchun float64 ishlatiladi
	StockQuantity int     `json:"stock_quantity"` // Ombordagi mahsulot miqdori
}

type UpdateProduct struct {
	ID            int     `json:"id"`
	Name          string  `json:"name"`
	Description   string  `json:"description"`
	Price         float64 `json:"price"`          // NUMERIC(10, 2) uchun float64 ishlatiladi
	StockQuantity int     `json:"stock_quantity"` // Ombordagi mahsulot miqdori
}

type GetListProductRequest struct {
	Offset int    `json:"offset"`
	Limit  int    `json:"limit"`
	Search string `json:"search"`
}

type GetListProductResponse struct {
	Count    int        `json:"count"`
	Products []*Product `json:"products"`
}
