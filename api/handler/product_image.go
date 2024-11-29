package handler

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strconv"

	"comics/models"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// Define the directory to save uploaded images
// const uploadDir = "./uploads/products/"

func (h *Handler) CreateProductImage(c *gin.Context) {
	file, fileHeader, err := c.Request.FormFile("image")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to get image"})
		return
	}
	defer file.Close()

	// Fayl hajmini tekshirish
	if fileHeader.Size > 5*1024*1024 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "File size exceeds 5MB"})
		return
	}

	// UUID asosida noyob nom yaratish
	fileName := fmt.Sprintf("%s%s", uuid.New().String(), filepath.Ext(fileHeader.Filename))
	filePath := filepath.Join("uploads", fileName)

	// Faylni saqlash
	out, err := os.Create(filePath)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save image"})
		return
	}
	defer out.Close()

	if _, err := io.Copy(out, file); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to write file"})
		return
	}

	// Qo'shimcha ma'lumotlarni olish
	productId := c.PostForm("product_id")
	productID, _ := strconv.Atoi(productId)
	isPrimary := c.PostForm("is_primary")
	isprimary, _ := strconv.ParseBool(isPrimary)

	// Ma'lumotlarni saqlash
	id, err := h.strg.ProductImage().Create(context.Background(), &models.CreateProductImage{
		ProductID: productID, ImageUrl: filePath, IsPrimary: isprimary,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to write file"})
		return
	}

	productImage, err := h.strg.ProductImage().GetByID(context.Background(), id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch product image"})
		return
	}

	// Javob qaytarish
	c.JSON(http.StatusOK, productImage)
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
