package handler

import (
	"comics/models"
	"context"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *Handler) CreateComicReview(c *gin.Context) {
	var entity *models.CreateComicsReview
	if err := c.BindJSON(&entity); err != nil {
		c.JSON(http.StatusBadRequest, models.DefaultError{
			Message: "Invalid request body: " + err.Error(),
		})
		return
	}

	fmt.Println(entity)

	// Create the ComicReview in the storage
	id, err := h.strg.ComicReview().Create(context.Background(), entity)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.DefaultError{
			Message: "Failed to create ComicReview: " + err.Error(),
		})
		return
	}

	// Get the ComicReview by ID
	comicReview, err := h.strg.ComicReview().GetByID(context.Background(), &models.PrimaryKey{Id: id.Id})
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.DefaultError{
			Message: "Failed to retrieve created ComicReview: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, models.SuccessResponse{
		Message: "ComicReview has been created",
		Data:    comicReview,
	})
}

func (h *Handler) UpdateComicReview(c *gin.Context) {
	var entity models.UpdateComicsReview
	fmt.Println(entity)
	if err := c.BindJSON(&entity); err != nil {
		c.JSON(http.StatusBadRequest, models.DefaultError{
			Message: "Invalid request body: " + err.Error(),
		})
		return
	}

	fmt.Println(entity)

	// Update the ComicReview
	if _, err := h.strg.ComicReview().Update(context.Background(), &entity); err != nil {
		c.JSON(http.StatusInternalServerError, models.DefaultError{
			Message: "Failed to update ComicReview: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, models.SuccessResponse{
		Message: "ComicReview has been updated",
		Data:    "ok",
	})
}

func (h *Handler) GetComicReviewsList(c *gin.Context) {
	// Retrieve the list of ComicReviews (offset and limit can be implemented later)
	resp, err := h.strg.ComicReview().GetList(context.Background(), &models.GetListComicsReviewRequest{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.DefaultError{
			Message: "Failed to retrieve ComicReview list: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, resp)
}

func (h *Handler) GetComicReviewsByIDHandler(c *gin.Context) {
	id := c.Param("id")
	idInt, err := strconv.Atoi(id)
	// Get the ComicReview by ID
	comicReview, err := h.strg.ComicReview().GetByID(context.Background(), &models.PrimaryKey{Id: idInt})
	if err != nil {
		c.JSON(http.StatusNotFound, models.DefaultError{
			Message: "ComicReview not found: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, models.SuccessResponse{
		Message: "OK",
		Data:    comicReview,
	})
}

func (h *Handler) DeleteComicReview(c *gin.Context) {
	id := c.Param("id")
	idInt, err := strconv.Atoi(id)
	// Delete the ComicReview by ID
	deletedComicReview, err := h.strg.ComicReview().Delete(context.Background(), &models.PrimaryKey{Id: idInt})
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.DefaultError{
			Message: "Failed to delete ComicReview: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, models.SuccessResponse{
		Message: "ComicReview has been deleted",
		Data:    deletedComicReview,
	})
}
