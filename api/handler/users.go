package handler

import (
	"comics/models"
	"context"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (h *Handler) CreateUser(c *gin.Context) {
	var entity *models.CreateUser
	if err := c.BindJSON(&entity); err != nil {
		c.JSON(http.StatusBadRequest, models.DefaultError{
			Message: "Invalid request body: " + err.Error(),
		})
		return
	}

	fmt.Println(entity)

	// Create the user in the storage
	id, err := h.strg.User().Create(context.Background(), entity)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.DefaultError{
			Message: "Failed to create user: " + err.Error(),
		})
		return
	}

	// Get the user by ID
	user, err := h.strg.User().GetByID(context.Background(), &models.UserPrimaryKey{Id: id.Id})
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.DefaultError{
			Message: "Failed to retrieve created user: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, models.SuccessResponse{
		Message: "User has been created",
		Data:    user,
	})
}

func (h *Handler) UpdateUser(c *gin.Context) {
	var entity models.UpdateUser
	if err := c.BindJSON(&entity); err != nil {
		c.JSON(http.StatusBadRequest, models.DefaultError{
			Message: "Invalid request body: " + err.Error(),
		})
		return
	}

	fmt.Println(entity)

	// Update the user
	if _, err := h.strg.User().Update(context.Background(), &entity); err != nil {
		c.JSON(http.StatusInternalServerError, models.DefaultError{
			Message: "Failed to update user: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, models.SuccessResponse{
		Message: "User has been updated",
		Data:    "ok",
	})
}

func (h *Handler) GetUsersList(c *gin.Context) {
	// Retrieve the list of users (offset and limit can be implemented later)
	resp, err := h.strg.User().GetList(context.Background(), &models.GetListUserRequest{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.DefaultError{
			Message: "Failed to retrieve user list: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, resp)
}

func (h *Handler) GetUsersByIDHandler(c *gin.Context) {
	id := c.Param("id")
	// Get the user by ID
	user, err := h.strg.User().GetByID(context.Background(), &models.UserPrimaryKey{Id: id})
	if err != nil {
		c.JSON(http.StatusNotFound, models.DefaultError{
			Message: "User not found: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, models.SuccessResponse{
		Message: "OK",
		Data:    user,
	})
}

func (h *Handler) DeleteUser(c *gin.Context) {
	id := c.Param("id")
	// Delete the user by ID
	deletedUser, err := h.strg.User().Delete(context.Background(), &models.UserPrimaryKey{Id: id})
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.DefaultError{
			Message: "Failed to delete user: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, models.SuccessResponse{
		Message: "User has been deleted",
		Data:    deletedUser,
	})
}

func (h *Handler) CreateUserImage(c *gin.Context) {
	file, fileHeader, err := c.Request.FormFile("image")
	if err != nil {
		fmt.Println("err 1", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to get image"})
		return
	}
	defer file.Close()

	// Fayl hajmini tekshirish
	if fileHeader.Size > 5*1024*1024 {
		fmt.Println("err 2", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "File size exceeds 5MB"})
		return
	}

	// UUID asosida noyob nom yaratish
	fileName := fmt.Sprintf("%s%s", uuid.New().String(), filepath.Ext(fileHeader.Filename))
	filePath := filepath.Join("uploads", fileName)

	// Faylni saqlash
	out, err := os.Create(filePath)
	if err != nil {
		fmt.Println("err 3", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save image"})
		return
	}
	defer out.Close()

	if _, err := io.Copy(out, file); err != nil {
		fmt.Println("err 4", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to write file"})
		return
	}

	// Qo'shimcha ma'lumotlarni olish
	userId := c.PostForm("user_id")

	user, err := h.strg.User().GetByID(context.Background(), &models.UserPrimaryKey{Id: userId})
	if err != nil {
		fmt.Println("err 5", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to write file"})
		return
	}
	user.ImageUrl = &filePath

	productImage, err := h.strg.User().Update(context.Background(), &models.UpdateUser{
		Id: user.Id,
		ImageUrl: user.ImageUrl,
	 })
	if err != nil {
		fmt.Println("err 6", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch product image"})
		return
	}

	// Javob qaytarish
	c.JSON(http.StatusOK, productImage)
}
