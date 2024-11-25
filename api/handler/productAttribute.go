package handler

import (
	"comics/models"
	"context"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *Handler) CreateProductAttribute(c *gin.Context) {
	var entity *models.CreateProductAttribute
	if err := c.BindJSON(&entity); err != nil {
		c.JSON(http.StatusBadRequest, models.DefaultError{
			Message: "Invalid request body: " + err.Error(),
		})
		return
	}

	fmt.Println(entity)

	// Create the ProductAttribute in the storage
	id, err := h.strg.ProductAttribute().Create(context.Background(), entity)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.DefaultError{
			Message: "Failed to create ProductAttribute: " + err.Error(),
		})
		return
	}

	// Get the ProductAttribute by ID
	productAttribute, err := h.strg.ProductAttribute().GetByID(context.Background(), &models.PrimaryKey{Id: id.Id})
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.DefaultError{
			Message: "Failed to retrieve created ProductAttribute: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, models.SuccessResponse{
		Message: "ProductAttribute has been created",
		Data:    productAttribute,
	})
}

func (h *Handler) UpdateProductAttribute(c *gin.Context) {
	var entity models.UpdateProductAttribute
	if err := c.BindJSON(&entity); err != nil {
		c.JSON(http.StatusBadRequest, models.DefaultError{
			Message: "Invalid request body: " + err.Error(),
		})
		return
	}

	// Update the ProductAttribute
	if _, err := h.strg.ProductAttribute().Update(context.Background(), &entity); err != nil {
		c.JSON(http.StatusInternalServerError, models.DefaultError{
			Message: "Failed to update ProductAttribute: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, models.SuccessResponse{
		Message: "ProductAttribute has been updated",
		Data:    "ok",
	})
}

func (h *Handler) GetProductAttributesList(c *gin.Context) {
	// Retrieve the list of ProductAttributes (offset and limit can be implemented later)
	resp, err := h.strg.ProductAttribute().GetList(context.Background(), &models.GetListProductAttributeRequest{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.DefaultError{
			Message: "Failed to retrieve ProductAttribute list: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, resp)
}

func (h *Handler) GetProductAttributesByIDHandler(c *gin.Context) {
	id := c.Param("id")
	intId,err:=strconv.Atoi(id)
	if err!=nil {
		c.JSON(http.StatusInternalServerError, models.DefaultError{
			Message: "Failed to retrieve ProductAttribute list: " + err.Error(),
		})
		return
	}
	// Get the ProductAttribute by ID
	productAttribute, err := h.strg.ProductAttribute().GetByID(context.Background(), &models.PrimaryKey{Id: intId})
	if err != nil {
		c.JSON(http.StatusNotFound, models.DefaultError{
			Message: "ProductAttribute not found: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, models.SuccessResponse{
		Message: "OK",
		Data:    productAttribute,
	})
}

func (h *Handler) DeleteProductAttribute(c *gin.Context) {
	id := c.Param("id")
	intId,err:=strconv.Atoi(id)
	if err!=nil {
		c.JSON(http.StatusInternalServerError, models.DefaultError{
			Message: "Failed to retrieve ProductAttribute list: " + err.Error(),
		})
		return
	}
	
	// Delete the ProductAttribute by ID
	deletedProductAttribute, err := h.strg.ProductAttribute().Delete(context.Background(), &models.PrimaryKey{Id: intId})
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.DefaultError{
			Message: "Failed to delete ProductAttribute: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, models.SuccessResponse{
		Message: "ProductAttribute has been deleted",
		Data:    deletedProductAttribute,
	})
}
