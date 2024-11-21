package storage

import (
	"comics/models"
	"context"
)

type StorageRepoI interface {
	User() UserRepoI
	Role() RoleRepoI
	Permission() PermissionRepoI
	RolePermission() RolePermissionRepoI
	UserRole() UserRoleRepoI
	Comics() ComicRepoI
	Order() OrderRepoI
	OrderItem() OrderItemRepoI
	ComicReview() ComicReviewRepoI
	ComicsPages() ComicPageRepoI 
	CloseDB()
}

type UserRepoI interface {
	Create(ctx context.Context, req *models.CreateUser) (*models.UserPrimaryKey, error)
	GetByID(ctx context.Context, req *models.UserPrimaryKey) (*models.User, error)
	GetByPhone(ctx context.Context, req *models.Login) (*models.User, error)
	GetList(ctx context.Context, req *models.GetListUserRequest) (resp *models.GetListUserResponse, err error)
	Update(ctx context.Context, req *models.UpdateUser) (int64, error)
	Delete(ctx context.Context, req *models.UserPrimaryKey) (int64, error)
}

type PermissionRepoI interface {
	Create(ctx context.Context, req *models.CreatePermission) (*models.PrimaryKey, error)
	GetByID(ctx context.Context, req *models.PrimaryKey) (*models.Permission, error)
	GetList(ctx context.Context, req *models.GetListPermissionRequest) (resp *models.GetListPermissionResponse, err error)
	Update(ctx context.Context, req *models.UpdatePermission) (int64, error)
	Delete(ctx context.Context, req *models.PrimaryKey) (int64, error)
}

type RolePermissionRepoI interface {
	Create(ctx context.Context, req *models.CreateRolePermission) (*models.PrimaryKey, error)
	GetByID(ctx context.Context, req *models.PrimaryKey) (*models.RolePermission, error)
	GetList(ctx context.Context, req *models.GetListRolePermissionRequest) (resp *models.GetListRolePermissionResponse, err error)
	Update(ctx context.Context, req *models.UpdateRolePermission) (int64, error)
	Delete(ctx context.Context, req *models.PrimaryKey) (int64, error)
}

type UserRoleRepoI interface {
	Create(ctx context.Context, req *models.CreateUserRole) error
	GetByID(ctx context.Context, req string) (*models.UserRole, error)
	GetList(ctx context.Context, req *models.GetListUserRoleRequest) (resp *models.GetListUserRoleResponse, err error)
	Update(ctx context.Context, req *models.UpdateUserRole) (int64, error)
	Delete(ctx context.Context, req *models.PrimaryKey) (int64, error)
}

type RoleRepoI interface {
	Create(ctx context.Context, req *models.CreateRole) (*models.PrimaryKey, error)
	GetByID(ctx context.Context, req *models.PrimaryKey) (*models.Role, error)
	GetList(ctx context.Context, req *models.GetListRoleRequest) (resp *models.GetListRoleResponse, err error)
	Update(ctx context.Context, req *models.UpdateRole) (int64, error)
	Delete(ctx context.Context, req *models.PrimaryKey) (int64, error)
}

type ComicRepoI interface {
	Create(ctx context.Context, req *models.CreateComics) (*models.PrimaryKey, error)
	GetByID(ctx context.Context, req *models.PrimaryKey) (*models.Comics, error)
	GetList(ctx context.Context, req *models.GetListComicsRequest) (resp *models.GetListComicsResponse, err error)
	Update(ctx context.Context, req *models.UpdateComics) (int64, error)
	Delete(ctx context.Context, req *models.PrimaryKey) (int64, error)
}

type ComicPageRepoI interface {
	Create(ctx context.Context, req *models.CreateComicsPages) (*models.PrimaryKey, error)
	GetByID(ctx context.Context, req *models.PrimaryKey) (*models.ComicsPages, error)
	GetList(ctx context.Context, req *models.GetListComicsPagesRequest) (resp *models.GetListComicsPagesResponse, err error)
	Update(ctx context.Context, req *models.UpdateComicsPages) (int64, error)
	Delete(ctx context.Context, req *models.PrimaryKey) (int64, error)
}

type ComicReviewRepoI interface {
	Create(ctx context.Context, req *models.CreateComicsReview) (*models.PrimaryKey, error)
	GetByID(ctx context.Context, req *models.PrimaryKey) (*models.ComicsReview, error)
	GetList(ctx context.Context, req *models.GetListComicsReviewRequest) (resp *models.GetListComicsReviewResponse, err error)
	Update(ctx context.Context, req *models.UpdateComicsReview) (int64, error)
	Delete(ctx context.Context, req *models.PrimaryKey) (int64, error)
}

type OrderRepoI interface {
	Create(ctx context.Context, req *models.CreateOrder) (*models.PrimaryKey, error)
	GetByID(ctx context.Context, req *models.PrimaryKey) (*models.Order, error)
	GetList(ctx context.Context, req *models.GetListOrderRequest) (resp *models.GetListOrderResponse, err error)
	Update(ctx context.Context, req *models.UpdateOrder) (int64, error)
	Delete(ctx context.Context, req *models.PrimaryKey) (int64, error)
}

type OrderItemRepoI interface {
	Create(ctx context.Context, req *models.CreateOrderItem) (*models.PrimaryKey, error)
	GetByID(ctx context.Context, req *models.PrimaryKey) (*models.OrderItem, error)
	GetList(ctx context.Context, req *models.GetListOrderItemRequest) (resp *models.GetListOrderItemResponse, err error)
	Update(ctx context.Context, req *models.UpdateOrderItem) (int64, error)
	Delete(ctx context.Context, req *models.PrimaryKey) (int64, error)
}

type PurchaseRepoI interface {
	Create(ctx context.Context, req *models.CreateRole) (*models.PrimaryKey, error)
	GetByID(ctx context.Context, req *models.PrimaryKey) (*models.Role, error)
	GetList(ctx context.Context, req *models.GetListRoleRequest) (resp *models.GetListRoleResponse, err error)
	Update(ctx context.Context, req *models.UpdateRole) (int64, error)
	Delete(ctx context.Context, req *models.PrimaryKey) (int64, error)
}

type DiscountRepoI interface {
	Create(ctx context.Context, req *models.CreateRole) (*models.PrimaryKey, error)
	GetByID(ctx context.Context, req *models.PrimaryKey) (*models.Role, error)
	GetList(ctx context.Context, req *models.GetListRoleRequest) (resp *models.GetListRoleResponse, err error)
	Update(ctx context.Context, req *models.UpdateRole) (int64, error)
	Delete(ctx context.Context, req *models.PrimaryKey) (int64, error)
}

type PromotionRepoI interface {
	Create(ctx context.Context, req *models.CreateRole) (*models.PrimaryKey, error)
	GetByID(ctx context.Context, req *models.PrimaryKey) (*models.Role, error)
	GetList(ctx context.Context, req *models.GetListRoleRequest) (resp *models.GetListRoleResponse, err error)
	Update(ctx context.Context, req *models.UpdateRole) (int64, error)
	Delete(ctx context.Context, req *models.PrimaryKey) (int64, error)
}

type PaymentRepoI interface {
	Create(ctx context.Context, req *models.CreateRole) (*models.PrimaryKey, error)
	GetByID(ctx context.Context, req *models.PrimaryKey) (*models.Role, error)
	GetList(ctx context.Context, req *models.GetListRoleRequest) (resp *models.GetListRoleResponse, err error)
	Update(ctx context.Context, req *models.UpdateRole) (int64, error)
	Delete(ctx context.Context, req *models.PrimaryKey) (int64, error)
}
