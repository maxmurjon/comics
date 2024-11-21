package handler

import (
	"comics/models"
	"context"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *Handler) CreateComics(c *gin.Context) {
	var entity *models.CreateComics
	if err := c.BindJSON(&entity); err != nil {
		c.JSON(http.StatusBadRequest, models.DefaultError{
			Message: "Invalid request body: " + err.Error(),
		})
		return
	}

	fmt.Println(entity)

	// Create the Comics in the storage
	id, err := h.strg.Comics().Create(context.Background(), entity)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.DefaultError{
			Message: "Failed to create Comics: " + err.Error(),
		})
		return
	}

	// Get the Comics by ID
	comics, err := h.strg.Comics().GetByID(context.Background(), &models.PrimaryKey{Id: id.Id})
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.DefaultError{
			Message: "Failed to retrieve created Comics: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, models.SuccessResponse{
		Message: "Comics has been created",
		Data:    comics,
	})
}

func (h *Handler) UpdateComics(c *gin.Context) {
	var entity models.UpdateComics
	fmt.Println(entity)
	if err := c.BindJSON(&entity); err != nil {
		c.JSON(http.StatusBadRequest, models.DefaultError{
			Message: "Invalid request body: " + err.Error(),
		})
		return
	}

	fmt.Println(entity)

	// Update the Comics
	if _, err := h.strg.Comics().Update(context.Background(), &entity); err != nil {
		c.JSON(http.StatusInternalServerError, models.DefaultError{
			Message: "Failed to update Comics: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, models.SuccessResponse{
		Message: "Comics has been updated",
		Data:    "ok",
	})
}

func (h *Handler) GetComicssList(c *gin.Context) {
	// Retrieve the list of Comicss (offset and limit can be implemented later)
	resp, err := h.strg.Comics().GetList(context.Background(), &models.GetListComicsRequest{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.DefaultError{
			Message: "Failed to retrieve Comics list: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, resp)
}

func (h *Handler) GetComicssByIDHandler(c *gin.Context) {
	id := c.Param("id")
	idInt, err := strconv.Atoi(id)
	// Get the Comics by ID
	comics, err := h.strg.Comics().GetByID(context.Background(), &models.PrimaryKey{Id: idInt})
	if err != nil {
		c.JSON(http.StatusNotFound, models.DefaultError{
			Message: "Comics not found: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, models.SuccessResponse{
		Message: "OK",
		Data:    comics,
	})
}

func (h *Handler) DeleteComics(c *gin.Context) {
	id := c.Param("id")
	idInt, err := strconv.Atoi(id)
	// Delete the Comics by ID
	deletedComics, err := h.strg.Comics().Delete(context.Background(), &models.PrimaryKey{Id: idInt})
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.DefaultError{
			Message: "Failed to delete Comics: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, models.SuccessResponse{
		Message: "Comics has been deleted",
		Data:    deletedComics,
	})
}
