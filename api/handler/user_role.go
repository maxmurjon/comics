package handler

import (
	"comics/models"
	"context"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *Handler) CreateUserRole(c *gin.Context) {
	var entity *models.CreateUserRole
	if err := c.BindJSON(&entity); err != nil {
		c.JSON(http.StatusBadRequest, models.DefaultError{
			Message: "Invalid request body: " + err.Error(),
		})
		return
	}

	fmt.Println(entity)

	// Create the UserRole in the storage
	err := h.strg.UserRole().Create(context.Background(), entity)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.DefaultError{
			Message: "Failed to create UserRole: " + err.Error(),
		})
		return
	}

	// Get the UserRole by ID
	userRole, err := h.strg.UserRole().GetByID(context.Background(), entity.UserId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.DefaultError{
			Message: "Failed to retrieve created UserRole: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, models.SuccessResponse{
		Message: "UserRole has been created",
		Data:    userRole,
	})
}

func (h *Handler) UpdateUserRole(c *gin.Context) {
	var entity models.UpdateUserRole
	if err := c.BindJSON(&entity); err != nil {
		c.JSON(http.StatusBadRequest, models.DefaultError{
			Message: "Invalid request body: " + err.Error(),
		})
		return
	}

	fmt.Println(entity)

	// Update the UserRole
	if _, err := h.strg.UserRole().Update(context.Background(), &entity); err != nil {
		c.JSON(http.StatusInternalServerError, models.DefaultError{
			Message: "Failed to update UserRole: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, models.SuccessResponse{
		Message: "UserRole has been updated",
		Data:    "ok",
	})
}

func (h *Handler) GetUserRolesList(c *gin.Context) {
	// Retrieve the list of UserRoles (offset and limit can be implemented later)
	resp, err := h.strg.UserRole().GetList(context.Background(), &models.GetListUserRoleRequest{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.DefaultError{
			Message: "Failed to retrieve UserRole list: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, resp)
}

func (h *Handler) GetUserRolesByIDHandler(c *gin.Context) {
	id := c.Param("id")
	// Get the UserRole by ID
	userRole, err := h.strg.UserRole().GetByID(context.Background(), id)
	if err != nil {
		c.JSON(http.StatusNotFound, models.DefaultError{
			Message: "UserRole not found: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, models.SuccessResponse{
		Message: "OK",
		Data:    userRole,
	})
}

func (h *Handler) DeleteUserRole(c *gin.Context) {
	id := c.Param("id")

	intId, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusNotFound, models.DefaultError{
			Message: "UserRole not found: " + err.Error(),
		})
		return
	}
	// Delete the UserRole by ID
	deletedUserRole, err := h.strg.UserRole().Delete(context.Background(), &models.PrimaryKey{Id: intId})
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.DefaultError{
			Message: "Failed to delete UserRole: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, models.SuccessResponse{
		Message: "UserRole has been deleted",
		Data:    deletedUserRole,
	})
}
