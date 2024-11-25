package handler

import (
	"comics/models"
	"context"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *Handler) CreateCategory(c *gin.Context) {
	var entity *models.CreateCategory
	fmt.Println("1111")
	if err := c.BindJSON(&entity); err != nil {
		c.JSON(http.StatusBadRequest, models.DefaultError{
			Message: "Invalid request body: " + err.Error(),
		})
		return
	}

	fmt.Println(entity)

	// Create the Category in the storage
	id, err := h.strg.Category().Create(context.Background(), entity)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.DefaultError{
			Message: "Failed to create Category: " + err.Error(),
		})
		return
	}

	// Get the Category by ID
	category, err := h.strg.Category().GetByID(context.Background(), &models.PrimaryKey{Id: id.Id})
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.DefaultError{
			Message: "Failed to retrieve created Category: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, models.SuccessResponse{
		Message: "Category has been created",
		Data:    category,
	})
}

func (h *Handler) UpdateCategory(c *gin.Context) {
	var entity models.UpdateCategory
	if err := c.BindJSON(&entity); err != nil {
		c.JSON(http.StatusBadRequest, models.DefaultError{
			Message: "Invalid request body: " + err.Error(),
		})
		return
	}

	fmt.Println(entity)

	// Update the Category
	if _, err := h.strg.Category().Update(context.Background(), &entity); err != nil {
		c.JSON(http.StatusInternalServerError, models.DefaultError{
			Message: "Failed to update Category: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, models.SuccessResponse{
		Message: "Category has been updated",
		Data:    "ok",
	})
}

func (h *Handler) GetCategorysList(c *gin.Context) {
	// Retrieve the list of Categorys (offset and limit can be implemented later)
	resp, err := h.strg.Category().GetList(context.Background(), &models.GetListCategoryRequest{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.DefaultError{
			Message: "Failed to retrieve Category list: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, resp)
}

func (h *Handler) GetCategorysByIDHandler(c *gin.Context) {
	id := c.Param("id")
	intId,err:=strconv.Atoi(id)
	if err!=nil{
		c.JSON(http.StatusNotFound, models.DefaultError{
			Message: "Category not found: " + err.Error(),
		})
		return
	}
	// Get the Category by ID
	category, err := h.strg.Category().GetByID(context.Background(), &models.PrimaryKey{Id: intId})
	if err != nil {
		c.JSON(http.StatusNotFound, models.DefaultError{
			Message: "Category not found: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, models.SuccessResponse{
		Message: "OK",
		Data:    category,
	})
}

func (h *Handler) DeleteCategory(c *gin.Context) {
	id := c.Param("id")

	intId,err:=strconv.Atoi(id)
	if err!=nil{
		c.JSON(http.StatusNotFound, models.DefaultError{
			Message: "Category not found: " + err.Error(),
		})
		return
	}
	// Delete the Category by ID
	deletedCategory, err := h.strg.Category().Delete(context.Background(), &models.PrimaryKey{Id: intId})
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.DefaultError{
			Message: "Failed to delete Category: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, models.SuccessResponse{
		Message: "Category has been deleted",
		Data:    deletedCategory,
	})
}
