package models

type Category struct {
	ID          string `json:"id"`
	Name    string `json:"name"`
	CreatedAt string `json:"created_at"`
}

type CreateCategory struct {
	Name    string `json:"name"`
}

type UpdateCategory struct {
	ID          string `json:"id"`
	Name    string `json:"name"`
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
