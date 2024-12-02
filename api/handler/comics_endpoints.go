package handler

import (
	"comics/models"
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) GetProductList(c *gin.Context) {
	resp:=[]models.ProductInfo{}

	// products
	
	_, err := h.strg.Product().GetList(context.Background(), &models.GetListProductRequest{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.DefaultError{
			Message: "Failed to retrieve Product list: " + err.Error(),
		})
		return
	}

	// _, err = h.strg.Category().GetByID(context.Background(), &models.GetListCategoryRequest{})
	// if err != nil {
	// 	c.JSON(http.StatusInternalServerError, models.DefaultError{
	// 		Message: "Failed to retrieve Category list: " + err.Error(),
	// 	})
	// 	return
	// }

	c.JSON(http.StatusOK, resp)
}
