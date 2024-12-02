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
	Category() CategoryRepoI
	Product() ProductRepoI
	ProductImage() ProductImageRepoI
	ProductAttribute() ProductAttributeRepoI
	Attribute() AttributeRepoI
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

type CategoryRepoI interface {
	Create(ctx context.Context, req *models.CreateCategory) (*models.PrimaryKey, error)
	GetByID(ctx context.Context, req *models.PrimaryKey) (*models.Category, error)
	GetList(ctx context.Context, req *models.GetListCategoryRequest) (resp *models.GetListCategoryResponse, err error)
	Update(ctx context.Context, req *models.UpdateCategory) (int64, error)
	Delete(ctx context.Context, req *models.PrimaryKey) (int64, error)
}

type ProductRepoI interface {
	Create(ctx context.Context, req *models.CreateProduct) (*models.PrimaryKey, error)
	GetByID(ctx context.Context, req *models.PrimaryKey) (*models.Product, error)
	GetList(ctx context.Context, req *models.GetListProductRequest) (resp *models.GetListProductResponse, err error)
	Update(ctx context.Context, req *models.UpdateProduct) (int64, error)
	Delete(ctx context.Context, req *models.PrimaryKey) (int64, error)
}

type ProductImageRepoI interface {
	Create(ctx context.Context, req *models.CreateProductImage) (*models.PrimaryKey, error)
	GetByID(ctx context.Context, req *models.PrimaryKey) (*models.ProductImage, error)
	GetList(ctx context.Context, req *models.GetListProductImageRequest) (resp *models.GetListProductImageResponse, err error)
	Update(ctx context.Context, req *models.UpdateProductImage) (int64, error)
	Delete(ctx context.Context, req *models.PrimaryKey) (int64, error)
}

type AttributeRepoI interface {
	Create(ctx context.Context, req *models.CreateAttribute) (*models.PrimaryKey, error)
	GetByID(ctx context.Context, req *models.PrimaryKey) (*models.Attribute, error)
	GetList(ctx context.Context, req *models.GetListAttributeRequest) (resp *models.GetListAttributeResponse, err error)
	Update(ctx context.Context, req *models.UpdateAttribute) (int64, error)
	Delete(ctx context.Context, req *models.PrimaryKey) (int64, error)
}

type ProductAttributeRepoI interface {
	Create(ctx context.Context, req *models.CreateProductAttribute) (*models.PrimaryKey, error)
	GetByID(ctx context.Context, req *models.PrimaryKey) (*models.ProductAttribute, error)
	GetList(ctx context.Context, req *models.GetListProductAttributeRequest) (resp *models.GetListProductAttributeResponse, err error)
	GetByProductID(ctx context.Context, req *models.PrimaryKey) (*models.ProductAttribute, error)
	Update(ctx context.Context, req *models.UpdateProductAttribute) (int64, error)
	Delete(ctx context.Context, req *models.PrimaryKey) (int64, error)
}