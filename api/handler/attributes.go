package handler

import (
	"comics/models"
	"context"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *Handler) CreateAttribute(c *gin.Context) {
	var entity *models.CreateAttribute
	if err := c.BindJSON(&entity); err != nil {
		c.JSON(http.StatusBadRequest, models.DefaultError{
			Message: "Invalid request body: " + err.Error(),
		})
		return
	}

	fmt.Println(entity)

	// Create the Attribute in the storage
	id, err := h.strg.Attribute().Create(context.Background(), entity)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.DefaultError{
			Message: "Failed to create Attribute: " + err.Error(),
		})
		return
	}

	// Get the Attribute by ID
	Attribute, err := h.strg.Attribute().GetByID(context.Background(), &models.PrimaryKey{Id: id.Id})
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.DefaultError{
			Message: "Failed to retrieve created Attribute: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, models.SuccessResponse{
		Message: "Attribute has been created",
		Data:    Attribute,
	})
}

func (h *Handler) UpdateAttribute(c *gin.Context) {
	var entity models.UpdateAttribute
	if err := c.BindJSON(&entity); err != nil {
		c.JSON(http.StatusBadRequest, models.DefaultError{
			Message: "Invalid request body: " + err.Error(),
		})
		return
	}

	// Update the Attribute
	if _, err := h.strg.Attribute().Update(context.Background(), &entity); err != nil {
		c.JSON(http.StatusInternalServerError, models.DefaultError{
			Message: "Failed to update Attribute: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, models.SuccessResponse{
		Message: "Attribute has been updated",
		Data:    "ok",
	})
}

func (h *Handler) GetAttributesList(c *gin.Context) {
	// Retrieve the list of Attributes (offset and limit can be implemented later)
	resp, err := h.strg.Attribute().GetList(context.Background(), &models.GetListAttributeRequest{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.DefaultError{
			Message: "Failed to retrieve Attribute list: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, resp)
}

func (h *Handler) GetAttributesByIDHandler(c *gin.Context) {
	id := c.Param("id")
	intId,err:=strconv.Atoi(id)
	if err!=nil {
		c.JSON(http.StatusInternalServerError, models.DefaultError{
			Message: "Failed to retrieve Attribute list: " + err.Error(),
		})
		return
	}
	// Get the Attribute by ID
	Attribute, err := h.strg.Attribute().GetByID(context.Background(), &models.PrimaryKey{Id: intId})
	if err != nil {
		c.JSON(http.StatusNotFound, models.DefaultError{
			Message: "Attribute not found: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, models.SuccessResponse{
		Message: "OK",
		Data:    Attribute,
	})
}

func (h *Handler) DeleteAttribute(c *gin.Context) {
	id := c.Param("id")
	intId,err:=strconv.Atoi(id)
	if err!=nil {
		c.JSON(http.StatusInternalServerError, models.DefaultError{
			Message: "Failed to retrieve Attribute list: " + err.Error(),
		})
		return
	}
	
	// Delete the Attribute by ID
	deletedAttribute, err := h.strg.Attribute().Delete(context.Background(), &models.PrimaryKey{Id: intId})
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.DefaultError{
			Message: "Failed to delete Attribute: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, models.SuccessResponse{
		Message: "Attribute has been deleted",
		Data:    deletedAttribute,
	})
}
