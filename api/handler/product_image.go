package handler

import (
	"context"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"time"

	"comics/models"

	"github.com/gin-gonic/gin"
)

// Define the directory to save uploaded images
const uploadDir = "./uploads/products/"

func (h *Handler) CreateProductImage(c *gin.Context) {
	// Parse multipart form to get the file
	file, fileHeader, err := c.Request.FormFile("image")
	if err != nil {
		c.JSON(http.StatusBadRequest, models.DefaultError{
			Message: "Invalid request body: " + err.Error(),
		})
		return
	}
	defer file.Close()

	// Ensure upload directory exists
	if err := os.MkdirAll(uploadDir, os.ModePerm); err != nil {
		c.JSON(http.StatusInternalServerError, models.DefaultError{
			Message: "Failed to create upload directory: " + err.Error(),
		})
		return
	}

	// Sanitize and generate a unique file name
	sanitizedFileName := filepath.Base(fileHeader.Filename)
	uniqueFileName := strconv.FormatInt(time.Now().UnixNano(), 10) + filepath.Ext(sanitizedFileName)
	filePath := filepath.Join(uploadDir, uniqueFileName)

	// Save the file to the server
	out, err := os.Create(filePath)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.DefaultError{
			Message: "Failed to save file: " + err.Error(),
		})
		return
	}
	defer out.Close()

	if _, err := io.Copy(out, file); err != nil {
		c.JSON(http.StatusInternalServerError, models.DefaultError{
			Message: "Failed to write file: " + err.Error(),
		})
		return
	}

	// Parse and validate product_id
	productIdStr := c.PostForm("product_id")
	productId, err := strconv.Atoi(productIdStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.DefaultError{
			Message: "Invalid product_id: " + err.Error(),
		})
		return
	}

	// Prepare the entity for database insertion
	relativeFilePath := "/uploads/products/" + uniqueFileName
	entity := &models.CreateProductImage{
		ProductID: productId,
		ImageUrl:  relativeFilePath,
		IsPrimary: c.PostForm("is_primary") == "true",
	}

	// Insert into database
	id, err := h.strg.ProductImage().Create(context.Background(), entity)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.DefaultError{
			Message: "Failed to create ProductImage: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, models.SuccessResponse{
		Message: "ProductImage has been created",
		Data:    id,
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

	// Update the ProductImage in storage
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
	// Retrieve the list of ProductImages
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
	intId, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.DefaultError{
			Message: "Invalid ProductImage ID: " + err.Error(),
		})
		return
	}

	// Retrieve the ProductImage by ID
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
	intId, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.DefaultError{
			Message: "Invalid ProductImage ID: " + err.Error(),
		})
		return
	}

	// Delete the ProductImage
	deletedProductImage, err := h.strg.ProductImage().Delete(context.Background(), &models.PrimaryKey{Id: intId})
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.DefaultError{
			Message: "Failed to delete ProductImage: " + err.Error(),
		})
		return
	}

	// Also delete the file from the server
	// filePath := deletedProductImage.ImageUrl
	// if err := os.Remove(filePath); err != nil && !os.IsNotExist(err) {
	// 	c.JSON(http.StatusInternalServerError, models.DefaultError{
	// 		Message: "Failed to delete file: " + err.Error(),
	// 	})
	// 	return
	// }

	c.JSON(http.StatusOK, models.SuccessResponse{
		Message: "ProductImage has been deleted",
		Data:    deletedProductImage,
	})
}
