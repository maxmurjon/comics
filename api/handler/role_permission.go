package handler

import (
	"comics/models"
	"context"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *Handler) CreateRolePermission(c *gin.Context) {
	var entity *models.CreateRolePermission
	if err := c.BindJSON(&entity); err != nil {
		c.JSON(http.StatusBadRequest, models.DefaultError{
			Message: "Invalid request body: " + err.Error(),
		})
		return
	}

	fmt.Println(entity)

	// Create the RoleRolePermission in the storage
	id, err := h.strg.RolePermission().Create(context.Background(), entity)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.DefaultError{
			Message: "Failed to create RolePermission: " + err.Error(),
		})
		return
	}

	// Get the RolePermission by ID
	rolePermission, err := h.strg.RolePermission().GetByID(context.Background(), &models.PrimaryKey{Id: id.Id})
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.DefaultError{
			Message: "Failed to retrieve created RolePermission: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, models.SuccessResponse{
		Message: "RolePermission has been created",
		Data:    rolePermission,
	})
}

func (h *Handler) UpdateRolePermission(c *gin.Context) {
	var entity models.UpdateRolePermission
	if err := c.BindJSON(&entity); err != nil {
		c.JSON(http.StatusBadRequest, models.DefaultError{
			Message: "Invalid request body: " + err.Error(),
		})
		return
	}

	fmt.Println(entity)

	// Update the RolePermission
	if _, err := h.strg.RolePermission().Update(context.Background(), &entity); err != nil {
		c.JSON(http.StatusInternalServerError, models.DefaultError{
			Message: "Failed to update RolePermission: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, models.SuccessResponse{
		Message: "RolePermission has been updated",
		Data:    "ok",
	})
}

func (h *Handler) GetRolePermissionsList(c *gin.Context) {
	// Retrieve the list of RolePermissions (offset and limit can be implemented later)
	resp, err := h.strg.RolePermission().GetList(context.Background(), &models.GetListRolePermissionRequest{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.DefaultError{
			Message: "Failed to retrieve RolePermission list: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, resp)
}

func (h *Handler) GetRolePermissionsByIDHandler(c *gin.Context) {
	id := c.Param("id")
	intId,err:=strconv.Atoi(id)
	if err!=nil{
		c.JSON(http.StatusNotFound, models.DefaultError{
			Message: "RolePermission not found: " + err.Error(),
		})
		return
	}
	// Get the RolePermission by ID
	rolePermission, err := h.strg.RolePermission().GetByID(context.Background(), &models.PrimaryKey{Id: intId})
	if err != nil {
		c.JSON(http.StatusNotFound, models.DefaultError{
			Message: "RolePermission not found: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, models.SuccessResponse{
		Message: "OK",
		Data:    rolePermission,
	})
}

func (h *Handler) DeleteRolePermission(c *gin.Context) {
	id := c.Param("id")

	intId,err:=strconv.Atoi(id)
	if err!=nil{
		c.JSON(http.StatusNotFound, models.DefaultError{
			Message: "RolePermission not found: " + err.Error(),
		})
		return
	}
	// Delete the RolePermission by ID
	deletedRolePermission, err := h.strg.RolePermission().Delete(context.Background(), &models.PrimaryKey{Id: intId})
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.DefaultError{
			Message: "Failed to delete RolePermission: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, models.SuccessResponse{
		Message: "RolePermission has been deleted",
		Data:    deletedRolePermission,
	})
}
