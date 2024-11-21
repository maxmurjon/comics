package handler

import (
	"comics/models"
	"context"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *Handler) CreatePermission(c *gin.Context) {
	var entity *models.CreatePermission
	fmt.Println("1111")
	if err := c.BindJSON(&entity); err != nil {
		c.JSON(http.StatusBadRequest, models.DefaultError{
			Message: "Invalid request body: " + err.Error(),
		})
		return
	}

	fmt.Println(entity)

	// Create the Permission in the storage
	id, err := h.strg.Permission().Create(context.Background(), entity)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.DefaultError{
			Message: "Failed to create Permission: " + err.Error(),
		})
		return
	}

	// Get the Permission by ID
	permission, err := h.strg.Permission().GetByID(context.Background(), &models.PrimaryKey{Id: id.Id})
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.DefaultError{
			Message: "Failed to retrieve created Permission: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, models.SuccessResponse{
		Message: "Permission has been created",
		Data:    permission,
	})
}

func (h *Handler) UpdatePermission(c *gin.Context) {
	var entity models.UpdatePermission
	if err := c.BindJSON(&entity); err != nil {
		c.JSON(http.StatusBadRequest, models.DefaultError{
			Message: "Invalid request body: " + err.Error(),
		})
		return
	}

	fmt.Println(entity)

	// Update the Permission
	if _, err := h.strg.Permission().Update(context.Background(), &entity); err != nil {
		c.JSON(http.StatusInternalServerError, models.DefaultError{
			Message: "Failed to update Permission: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, models.SuccessResponse{
		Message: "Permission has been updated",
		Data:    "ok",
	})
}

func (h *Handler) GetPermissionsList(c *gin.Context) {
	// Retrieve the list of Permissions (offset and limit can be implemented later)
	resp, err := h.strg.Permission().GetList(context.Background(), &models.GetListPermissionRequest{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.DefaultError{
			Message: "Failed to retrieve Permission list: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, resp)
}

func (h *Handler) GetPermissionsByIDHandler(c *gin.Context) {
	id := c.Param("id")
	intId,err:=strconv.Atoi(id)
	if err!=nil{
		c.JSON(http.StatusNotFound, models.DefaultError{
			Message: "Permission not found: " + err.Error(),
		})
		return
	}
	// Get the Permission by ID
	permission, err := h.strg.Permission().GetByID(context.Background(), &models.PrimaryKey{Id: intId})
	if err != nil {
		c.JSON(http.StatusNotFound, models.DefaultError{
			Message: "Permission not found: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, models.SuccessResponse{
		Message: "OK",
		Data:    permission,
	})
}

func (h *Handler) DeletePermission(c *gin.Context) {
	id := c.Param("id")

	intId,err:=strconv.Atoi(id)
	if err!=nil{
		c.JSON(http.StatusNotFound, models.DefaultError{
			Message: "Permission not found: " + err.Error(),
		})
		return
	}
	// Delete the Permission by ID
	deletedPermission, err := h.strg.Permission().Delete(context.Background(), &models.PrimaryKey{Id: intId})
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.DefaultError{
			Message: "Failed to delete Permission: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, models.SuccessResponse{
		Message: "Permission has been deleted",
		Data:    deletedPermission,
	})
}
