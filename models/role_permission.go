package models

type RolePermission struct {
	RoleId       string `json:"role_id"`
	PermissionId string `json:"permission_id"`
}

type CreateRolePermission struct {
	RoleId       string `json:"role_id"`
	PermissionId string `json:"permission_id"`
}

type DeleteRolePermission struct {
	RoleId       string `json:"role_id"`
	PermissionId string `json:"permission_id"`
}


type UpdateRolePermission struct {
	RoleId       string `json:"role_id"`
	PermissionId string `json:"permission_id"`
}

type GetListRolePermissionRequest struct {
	Offset int    `json:"offset"`
	Limit  int    `json:"limit"`
	Search string `json:"search"`
}

type GetListRolePermissionResponse struct {
	Count int     `json:"count"`
	RolePermission []*RolePermission `json:"roles"`
}
