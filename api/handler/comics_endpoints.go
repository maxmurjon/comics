package handler

import (
	"comics/models"
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
)



func (h *Handler) GetProductList(c *gin.Context) {
	resp, err := h.strg.Product().GetList(context.Background(), &models.GetListProductRequest{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.DefaultError{
			Message: "Failed to retrieve Product list: " + err.Error(),
		})
		return
	}
	resp_2, err := h.strg.Attribute().GetList(context.Background(), &models.GetListAttributeRequest{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.DefaultError{
			Message: "Failed to retrieve Attribute list: " + err.Error(),
		})
		return
	}

	resp_3, err := h.strg.ProductImage().GetList(context.Background(), &models.GetListProductImageRequest{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.DefaultError{
			Message: "Failed to retrieve ProductImage list: " + err.Error(),
		})
		return
	}

	resp_4, err := h.strg.Category().GetList(context.Background(), &models.GetListCategoryRequest{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.DefaultError{
			Message: "Failed to retrieve Category list: " + err.Error(),
		})
		return
	}

	products := models.ProductList{}
	products.Products=resp.Products
	products.Attributes=resp_2.Attributes
	products.ImageURLs=resp_3.ProductImages
	products.Categories=resp_4.Categories

	c.JSON(http.StatusOK, resp)
}
