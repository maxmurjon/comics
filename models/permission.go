package models

type Permission struct {
	Id             string `json:"id"`
	PermissionName string `json:"permission_name"`
	Description string `json:"description"`
}

type CreatePermission struct {
	PermissionName string `json:"permission_name"`
	Description string `json:"description"`
}

type UpdatePermission struct {
	Id             string `json:"id"`
	PermissionName string `json:"permission_name"`
	Description string `json:"description"`
}

type GetListPermissionRequest struct {
	Offset int    `json:"offset"`
	Limit  int    `json:"limit"`
	Search string `json:"search"`
}

type GetListPermissionResponse struct {
	Count       int           `json:"count"`
	Permissions []*Permission `json:"permissions"`
}
