package models

type UserRole struct {
	UserId string `json:"user_id"`
	RoleId int `json:"role_id"`
}

type CreateUserRole struct {
	UserId string `json:"user_id"`
	RoleId int `json:"role_id"`
}

type DeleteUserRole struct {
	UserId string `json:"user_id"`
	RoleId int `json:"role_id"`
}

type UpdateUserRole struct {
	UserId string `json:"user_id"`
	RoleId int `json:"role_id"`
}

type GetListUserRoleRequest struct {
	Offset int    `json:"offset"`
	Limit  int    `json:"limit"`
	Search string `json:"search"`
}

type GetListUserRoleResponse struct {
	Count     int     `json:"count"`
	UserRoles []*UserRole `json:"roles"`
}
