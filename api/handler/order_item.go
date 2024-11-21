package handler

import (
	"comics/models"
	"context"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *Handler) CreateOrderItem(c *gin.Context) {
	var entity *models.CreateOrderItem
	if err := c.BindJSON(&entity); err != nil {
		c.JSON(http.StatusBadRequest, models.DefaultError{
			Message: "Invalid request body: " + err.Error(),
		})
		return
	}

	fmt.Println(entity)

	// Create the OrderItem in the storage
	id, err := h.strg.OrderItem().Create(context.Background(), entity)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.DefaultError{
			Message: "Failed to create OrderItem: " + err.Error(),
		})
		return
	}

	// Get the OrderItem by ID
	OrderItem, err := h.strg.OrderItem().GetByID(context.Background(), &models.PrimaryKey{Id: id.Id})
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.DefaultError{
			Message: "Failed to retrieve created OrderItem: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, models.SuccessResponse{
		Message: "OrderItem has been created",
		Data:    OrderItem,
	})
}

func (h *Handler) UpdateOrderItem(c *gin.Context) {
	var entity models.UpdateOrderItem
	if err := c.BindJSON(&entity); err != nil {
		c.JSON(http.StatusBadRequest, models.DefaultError{
			Message: "Invalid request body: " + err.Error(),
		})
		return
	}

	fmt.Println(entity)

	// Update the OrderItem
	if _, err := h.strg.OrderItem().Update(context.Background(), &entity); err != nil {
		c.JSON(http.StatusInternalServerError, models.DefaultError{
			Message: "Failed to update OrderItem: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, models.SuccessResponse{
		Message: "OrderItem has been updated",
		Data:    "ok",
	})
}

func (h *Handler) GetOrderItemsList(c *gin.Context) {
	// Retrieve the list of OrderItems (offset and limit can be implemented later)
	resp, err := h.strg.OrderItem().GetList(context.Background(), &models.GetListOrderItemRequest{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.DefaultError{
			Message: "Failed to retrieve OrderItem list: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, resp)
}

func (h *Handler) GetOrderItemsByIDHandler(c *gin.Context) {
	id := c.Param("id")
	idInt, err := strconv.Atoi(id)
	// Get the OrderItem by ID
	OrderItem, err := h.strg.OrderItem().GetByID(context.Background(), &models.PrimaryKey{Id: idInt})
	if err != nil {
		c.JSON(http.StatusNotFound, models.DefaultError{
			Message: "OrderItem not found: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, models.SuccessResponse{
		Message: "OK",
		Data:    OrderItem,
	})
}

func (h *Handler) DeleteOrderItem(c *gin.Context) {
	id := c.Param("id")
	idInt, err := strconv.Atoi(id)
	// Delete the OrderItem by ID
	deletedOrderItem, err := h.strg.OrderItem().Delete(context.Background(), &models.PrimaryKey{Id: idInt})
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.DefaultError{
			Message: "Failed to delete OrderItem: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, models.SuccessResponse{
		Message: "OrderItem has been deleted",
		Data:    deletedOrderItem,
	})
}
