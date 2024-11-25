package handler

import (
	"comics/models"
	"context"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *Handler) CreateProductImage(c *gin.Context) {
	var entity *models.CreateProductImage
	if err := c.BindJSON(&entity); err != nil {
		c.JSON(http.StatusBadRequest, models.DefaultError{
			Message: "Invalid request body: " + err.Error(),
		})
		return
	}

	fmt.Println(entity)

	// Create the ProductImage in the storage
	id, err := h.strg.ProductImage().Create(context.Background(), entity)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.DefaultError{
			Message: "Failed to create ProductImage: " + err.Error(),
		})
		return
	}

	// Get the ProductImage by ID
	productImage, err := h.strg.ProductImage().GetByID(context.Background(), &models.PrimaryKey{Id: id.Id})
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.DefaultError{
			Message: "Failed to retrieve created ProductImage: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, models.SuccessResponse{
		Message: "ProductImage has been created",
		Data:    productImage,
	})
}

func (h *Handler) UpdateProductImage(c *gin.Context) {
	var entity models.UpdateProductImage
	if err := c.BindJSON(&entity); err != nil {
		c.JSON(http.StatusBadRequest, models.DefaultError{
			Message: "Invalid request body: " + err.Error(),
		})
		return
	}

	// Update the ProductImage
	if _, err := h.strg.ProductImage().Update(context.Background(), &entity); err != nil {
		c.JSON(http.StatusInternalServerError, models.DefaultError{
			Message: "Failed to update ProductImage: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, models.SuccessResponse{
		Message: "ProductImage has been updated",
		Data:    "ok",
	})
}

func (h *Handler) GetProductImagesList(c *gin.Context) {
	// Retrieve the list of ProductImages (offset and limit can be implemented later)
	resp, err := h.strg.ProductImage().GetList(context.Background(), &models.GetListProductImageRequest{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.DefaultError{
			Message: "Failed to retrieve ProductImage list: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, resp)
}

func (h *Handler) GetProductImagesByIDHandler(c *gin.Context) {
	id := c.Param("id")
	intId,err:=strconv.Atoi(id)
	if err!=nil {
		c.JSON(http.StatusInternalServerError, models.DefaultError{
			Message: "Failed to retrieve ProductImage list: " + err.Error(),
		})
		return
	}
	// Get the ProductImage by ID
	productImage, err := h.strg.ProductImage().GetByID(context.Background(), &models.PrimaryKey{Id: intId})
	if err != nil {
		c.JSON(http.StatusNotFound, models.DefaultError{
			Message: "ProductImage not found: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, models.SuccessResponse{
		Message: "OK",
		Data:    productImage,
	})
}

func (h *Handler) DeleteProductImage(c *gin.Context) {
	id := c.Param("id")
	intId,err:=strconv.Atoi(id)
	if err!=nil {
		c.JSON(http.StatusInternalServerError, models.DefaultError{
			Message: "Failed to retrieve ProductImage list: " + err.Error(),
		})
		return
	}
	
	// Delete the ProductImage by ID
	deletedProductImage, err := h.strg.ProductImage().Delete(context.Background(), &models.PrimaryKey{Id: intId})
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.DefaultError{
			Message: "Failed to delete ProductImage: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, models.SuccessResponse{
		Message: "ProductImage has been deleted",
		Data:    deletedProductImage,
	})
}
