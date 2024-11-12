package handler

import (
	"comics/models"
	"context"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) CreateUser(c *gin.Context) {
	var entity *models.CreateUser
	if err := c.BindJSON(&entity); err != nil {
		c.JSON(http.StatusBadRequest, models.DefaultError{
			Message: "Invalid request body: " + err.Error(),
		})
		return
	}

	fmt.Println(entity)

	// Create the user in the storage
	id, err := h.strg.User().Create(context.Background(), entity)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.DefaultError{
			Message: "Failed to create user: " + err.Error(),
		})
		return
	}

	// Get the user by ID
	user, err := h.strg.User().GetByID(context.Background(), &models.UserPrimaryKey{Id: id.Id})
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.DefaultError{
			Message: "Failed to retrieve created user: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, models.SuccessResponse{
		Message: "User has been created",
		Data:    user,
	})
}

func (h *Handler) UpdateUser(c *gin.Context) {
	var entity models.UpdateUser
	if err := c.BindJSON(&entity); err != nil {
		c.JSON(http.StatusBadRequest, models.DefaultError{
			Message: "Invalid request body: " + err.Error(),
		})
		return
	}

	fmt.Println(entity)

	// Update the user
	if _, err := h.strg.User().Update(context.Background(), &entity); err != nil {
		c.JSON(http.StatusInternalServerError, models.DefaultError{
			Message: "Failed to update user: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, models.SuccessResponse{
		Message: "User has been updated",
		Data:    "ok",
	})
}

func (h *Handler) GetUsersList(c *gin.Context) {
	// Retrieve the list of users (offset and limit can be implemented later)
	resp, err := h.strg.User().GetList(context.Background(), &models.GetListUserRequest{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.DefaultError{
			Message: "Failed to retrieve user list: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, resp)
}

func (h *Handler) GetUsersByIDHandler(c *gin.Context) {
	id := c.Param("id")
	// Get the user by ID
	user, err := h.strg.User().GetByID(context.Background(), &models.UserPrimaryKey{Id: id})
	if err != nil {
		c.JSON(http.StatusNotFound, models.DefaultError{
			Message: "User not found: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, models.SuccessResponse{
		Message: "OK",
		Data:    user,
	})
}

func (h *Handler) DeleteUser(c *gin.Context) {
	id := c.Param("id")
	// Delete the user by ID
	deletedUser, err := h.strg.User().Delete(context.Background(), &models.UserPrimaryKey{Id: id})
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.DefaultError{
			Message: "Failed to delete user: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, models.SuccessResponse{
		Message: "User has been deleted",
		Data:    deletedUser,
	})
}
