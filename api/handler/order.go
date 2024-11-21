package handler

import (
	"comics/models"
	"context"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *Handler) CreateOrder(c *gin.Context) {
	var entity *models.CreateOrder
	if err := c.BindJSON(&entity); err != nil {
		c.JSON(http.StatusBadRequest, models.DefaultError{
			Message: "Invalid request body: " + err.Error(),
		})
		return
	}

	fmt.Println(entity)

	// Create the Order in the storage
	id, err := h.strg.Order().Create(context.Background(), entity)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.DefaultError{
			Message: "Failed to create Order: " + err.Error(),
		})
		return
	}

	// Get the Order by ID
	Order, err := h.strg.Order().GetByID(context.Background(), &models.PrimaryKey{Id: id.Id})
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.DefaultError{
			Message: "Failed to retrieve created Order: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, models.SuccessResponse{
		Message: "Order has been created",
		Data:    Order,
	})
}

func (h *Handler) UpdateOrder(c *gin.Context) {
	var entity models.UpdateOrder
	if err := c.BindJSON(&entity); err != nil {
		c.JSON(http.StatusBadRequest, models.DefaultError{
			Message: "Invalid request body: " + err.Error(),
		})
		return
	}

	fmt.Println(entity)

	// Update the Order
	if _, err := h.strg.Order().Update(context.Background(), &entity); err != nil {
		c.JSON(http.StatusInternalServerError, models.DefaultError{
			Message: "Failed to update Order: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, models.SuccessResponse{
		Message: "Order has been updated",
		Data:    "ok",
	})
}

func (h *Handler) GetOrdersList(c *gin.Context) {
	// Retrieve the list of Orders (offset and limit can be implemented later)
	resp, err := h.strg.Order().GetList(context.Background(), &models.GetListOrderRequest{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.DefaultError{
			Message: "Failed to retrieve Order list: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, resp)
}

func (h *Handler) GetOrdersByIDHandler(c *gin.Context) {
	id := c.Param("id")
	idInt, err := strconv.Atoi(id)
	// Get the Order by ID
	Order, err := h.strg.Order().GetByID(context.Background(), &models.PrimaryKey{Id: idInt})
	if err != nil {
		c.JSON(http.StatusNotFound, models.DefaultError{
			Message: "Order not found: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, models.SuccessResponse{
		Message: "OK",
		Data:    Order,
	})
}

func (h *Handler) DeleteOrder(c *gin.Context) {
	id := c.Param("id")
	idInt, err := strconv.Atoi(id)
	// Delete the Order by ID
	deletedOrder, err := h.strg.Order().Delete(context.Background(), &models.PrimaryKey{Id: idInt})
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.DefaultError{
			Message: "Failed to delete Order: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, models.SuccessResponse{
		Message: "Comics has been deleted",
		Data:    deletedOrder,
	})
}
