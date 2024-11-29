package models

type Category struct {
	ID          string `json:"id"`
	Name    string `json:"name"`
	Description string `json:"description"`
}

type CreateCategory struct {
	Name    string `json:"name"`
	Description string `json:"description"`
}

type UpdateCategory struct {
	ID          string `json:"id"`
	Name    string `json:"name"`
	Description string `json:"description"`
}

type GetListCategoryRequest struct {
	Offset int    `json:"offset"`
	Limit  int    `json:"limit"`
	Search string `json:"search"`
}

type GetListCategoryResponse struct {
	Count int     `json:"count"`
	Categories []*Category `json:"categories"`
}
