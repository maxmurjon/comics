package handler

import (
	"comics/models"
	"context"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *Handler) CreateComicsPages(c *gin.Context) {
	var entity *models.CreateComicsPages
	if err := c.BindJSON(&entity); err != nil {
		c.JSON(http.StatusBadRequest, models.DefaultError{
			Message: "Invalid request body: " + err.Error(),
		})
		return
	}

	fmt.Println(entity)

	// Create the ComicsPages in the storage
	id, err := h.strg.ComicsPages().Create(context.Background(), entity)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.DefaultError{
			Message: "Failed to create ComicsPages: " + err.Error(),
		})
		return
	}

	// Get the ComicsPages by ID
	comicsPages, err := h.strg.ComicsPages().GetByID(context.Background(), &models.PrimaryKey{Id: id.Id})
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.DefaultError{
			Message: "Failed to retrieve created ComicsPages: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, models.SuccessResponse{
		Message: "ComicsPages has been created",
		Data:    comicsPages,
	})
}

func (h *Handler) UpdateComicsPages(c *gin.Context) {
	var entity models.UpdateComicsPages
	fmt.Println(entity)
	if err := c.BindJSON(&entity); err != nil {
		c.JSON(http.StatusBadRequest, models.DefaultError{
			Message: "Invalid request body: " + err.Error(),
		})
		return
	}

	fmt.Println(entity)

	// Update the ComicsPages
	if _, err := h.strg.ComicsPages().Update(context.Background(), &entity); err != nil {
		c.JSON(http.StatusInternalServerError, models.DefaultError{
			Message: "Failed to update ComicsPages: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, models.SuccessResponse{
		Message: "ComicsPages has been updated",
		Data:    "ok",
	})
}

func (h *Handler) GetComicsPagessList(c *gin.Context) {
	// Retrieve the list of ComicsPagess (offset and limit can be implemented later)
	resp, err := h.strg.ComicsPages().GetList(context.Background(), &models.GetListComicsPagesRequest{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.DefaultError{
			Message: "Failed to retrieve ComicsPages list: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, resp)
}

func (h *Handler) GetComicsPagessByIDHandler(c *gin.Context) {
	id := c.Param("id")
	idInt, err := strconv.Atoi(id)
	// Get the ComicsPages by ID
	comicsPages, err := h.strg.ComicsPages().GetByID(context.Background(), &models.PrimaryKey{Id: idInt})
	if err != nil {
		c.JSON(http.StatusNotFound, models.DefaultError{
			Message: "ComicsPages not found: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, models.SuccessResponse{
		Message: "OK",
		Data:    comicsPages,
	})
}

func (h *Handler) DeleteComicsPages(c *gin.Context) {
	id := c.Param("id")
	idInt, err := strconv.Atoi(id)
	// Delete the ComicsPages by ID
	deletedComicsPages, err := h.strg.ComicsPages().Delete(context.Background(), &models.PrimaryKey{Id: idInt})
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.DefaultError{
			Message: "Failed to delete ComicsPages: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, models.SuccessResponse{
		Message: "Comics has been deleted",
		Data:    deletedComicsPages,
	})
}
