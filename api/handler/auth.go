package handler

import (
	"comics/config"
	"comics/models"
	"comics/pkg/helper/helper"
	"context"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func (h *Handler) Register(c *gin.Context) {
	var createUser models.CreateUser

	err := c.ShouldBindJSON(&createUser)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.DefaultError{
			Message: "Register User: " + err.Error(),
		})
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(createUser.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.DefaultError{
			Message: "Error Hashing User's password: " + err.Error(),
		})
		return
	}

	createUser.Password = string(hashedPassword)

	userId, err := h.strg.User().Create(context.Background(), &createUser)
	if err != nil {
		if err.Error() == `ERROR: duplicate key value violates unique constraint "users_login_key" (SQLSTATE 23505)` {
			c.JSON(http.StatusConflict, models.DefaultError{
				Message: "User already exists, please login!",
			})
			return
		}
		c.JSON(http.StatusInternalServerError, models.DefaultError{
			Message: "Error creating user: " + err.Error(),
		})
		return
	}
	// err = h.strg.UserRole().Create(context.Background(), &models.CreateUserRole{UserId: userId.Id, RoleId: 2})
	// if err != nil {
	// 	c.JSON(http.StatusInternalServerError, models.DefaultError{
	// 		Message: "Error assigning role to user: " + err.Error(),
	// 	})
	// 	return
	// }

	user, err := h.strg.User().GetByID(context.Background(), userId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.DefaultError{
			Message: "Error fetching user information: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, user)
}

func (h *Handler) Login(c *gin.Context) {

	var login models.Login

	err := c.ShouldBindJSON(&login) // parse req body to given type struct
	if err != nil {
		c.JSON(http.StatusBadRequest, models.DefaultError{Message: "Parsing data error"})
		return
	}

	resp, err := h.strg.User().GetByPhone(context.Background(), &login)
	if err != nil {
		if err.Error() == "no rows in result set" {
			c.JSON(http.StatusBadRequest, models.DefaultError{Message: "storage.user.getByID" + "\nuser not found please register first"})
			return
		}

		fmt.Println(c, "storage.user.getByID", http.StatusInternalServerError, err.Error())
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(resp.Password), []byte(login.Password))
	if err != nil {
		c.JSON(http.StatusBadRequest, models.DefaultError{Message: "storage.user.getByID \ncredentials are wrong"})
		return
	}

	roleId, err := h.strg.UserRole().GetByID(context.Background(), resp.Id)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.DefaultError{Message: "storage.user.getByID \ncredentials are wrong"})
		return
	}

	role, err := h.strg.Role().GetByID(context.Background(), &models.PrimaryKey{Id: roleId.RoleId})
	if err != nil {
		c.JSON(http.StatusBadRequest, models.DefaultError{Message: "storage.user.getByID \ncredentials are wrong"})
		return
	}

	data := map[string]interface{}{
		"id":           resp.Id,
		"first_name":   resp.FirstName,
		"login":        resp.LastName,
		"phone_number": resp.PhoneNumber,
		"password":     resp.Password,
		"created_at":   resp.CreatedAt,
		"updated_at":   resp.UpdatedAt,
		"role":         role,
	}

	token, err := helper.GenerateJWT(data, config.TimeExpiredAt, h.cfg.SekretKey)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.DefaultError{Message: "storage.user.getByID"})
		return
	}
	// var bearer = "Bearer " + token

	c.JSON(http.StatusCreated, models.LoginResponse{Token: token, UserData: resp})
}
