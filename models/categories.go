package models

type Category struct {
	ID          string `json:"id"`
	Name    string `json:"name"`
	ParentId string `json:"parent_id"`
	CreatedAt string `json:"created_at"`
}

type CreateCategory struct {
	Name    string `json:"name"`
	ParentId string `json:"parent_id"`
}

type UpdateCategory struct {
	ID          string `json:"id"`
	Name    string `json:"name"`
	ParentId string `json:"parent_id"`
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
