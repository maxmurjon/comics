package api

import (
	"comics/api/handler"
	"comics/config"

	"github.com/gin-gonic/gin"
)

func SetUpAPI(r *gin.Engine, h handler.Handler, cfg config.Config) {
	
	r.Use(customCORSMiddleware()) 
	
	// Auth Endpoints
	r.POST("/register",h.Register)
	r.POST("/login",h.Login)
	// Users Endpoints
	r.POST("/createuser", h.CreateUser)
	r.PUT("/updateuser", h.UpdateUser)
	r.GET("/users", h.GetUsersList)
	r.GET("/user/:id", h.GetUsersByIDHandler)
	r.DELETE("/deleteuser/:id", h.DeleteUser)

	//Role Endpoints
	r.POST("/createrole", h.CreateRole)
	r.PUT("/updaterole", h.UpdateRole)
	r.GET("/roles", h.GetRolesList)
	r.GET("/role/:id", h.GetRolesByIDHandler)
	r.DELETE("/deleterole/:id", h.DeleteRole)

	//Permission Endpoints
	r.POST("/createpermission", h.CreatePermission)
	r.PUT("/updatepermission", h.UpdatePermission)
	r.GET("/permissions", h.GetPermissionsList)
	r.GET("/permission/:id", h.GetPermissionsByIDHandler)
	r.DELETE("/deletepermission/:id", h.DeletePermission)

	//RolePermission Endpoints
	r.POST("/createrolepermission", h.CreateRolePermission)
	r.PUT("/updaterolepermission", h.UpdateRolePermission)
	r.GET("/rolepermissions", h.GetRolePermissionsList)
	r.GET("/rolepermission/:id", h.GetRolePermissionsByIDHandler)
	r.DELETE("/deleterolepermission/:id", h.DeleteRolePermission)

	//UserRole Endpoints
	r.POST("/createuserrole", h.CreateUserRole)
	r.PUT("/updateuserrole", h.UpdateUserRole)
	r.GET("/userroles", h.GetUserRolesList)
	r.GET("/userrole/:id", h.GetUserRolesByIDHandler)
	r.DELETE("/deleteuserrole/:id", h.DeleteUserRole)

	//Category Endpoints
	r.POST("/createcategory", h.CreateCategory)
	r.PUT("/updatecategory", h.UpdateCategory)
	r.GET("/categorys", h.GetCategorysList)
	r.GET("/category/:id", h.GetCategorysByIDHandler)
	r.DELETE("/deletecategory/:id", h.DeleteCategory)

	//Product Endpoints
	r.POST("/createproduct", h.CreateProduct)
	r.PUT("/updateproduct", h.UpdateProduct)
	r.GET("/products", h.GetProductsList)
	r.GET("/product/:id", h.GetProductsByIDHandler)
	r.DELETE("/deleteproduct/:id", h.DeleteProduct)
	
	//ProductImage Endpoints
	r.POST("/createproductimage", h.CreateProductImage)
	r.PUT("/updateproductimage", h.UpdateProductImage)
	r.GET("/productimages", h.GetProductImagesList)
	r.GET("/productimage/:id", h.GetProductImagesByIDHandler)
	r.DELETE("/deleteproductimage/:id", h.DeleteProductImage)
	
	
}

func customCORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		c.Header("Acces-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, PATCH, DELETE, HEAD")
		c.Header("Access-Control-Allow-Headers", "Platform-Id, Content-Type, Content-Length, Accept-Encoding, X-CSF-TOKEN, Authorization, Cache-Control")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
