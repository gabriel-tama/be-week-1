// controllers/user_controller.go

package controllers

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/gabriel-tama/be-week-1/api/v1/services"
	"github.com/gabriel-tama/be-week-1/types"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
	"golang.org/x/crypto/bcrypt"
)

type UserController struct {
	userService services.UserService
	jwtService  services.JWTService
}

func NewUserController(userService services.UserService, jwtService services.JWTService) *UserController {
	return &UserController{userService: userService, jwtService: jwtService}
}

func (uc *UserController) Register(c *gin.Context) {
	var user types.RegisterData
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "name /password / username is too short or long"})
		return
	}

	exist, err := uc.userService.CheckUserExists(user.Username)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "server error"})
		return
	}

	if exist {
		c.JSON(http.StatusConflict, gin.H{"message": "user already exist"})
		return
	}

	data, err := uc.userService.Register(user.Username, user.Name, user.Password)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "server error"})
		return
	}

	tokenString, err := uc.jwtService.CreateToken(user.Username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "server error"})
	}

	authResponse := types.AuthResponse{
		Username:    data.Username,
		Name:        data.Name,
		AccessToken: tokenString,
	}

	c.JSON(http.StatusCreated, gin.H{"message": "User registered successfully", "data": authResponse})
}

func (uc *UserController) Login(c *gin.Context) {
	var loginData types.LoginData
	if err := c.ShouldBindJSON(&loginData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "password / username is too short or long"})
		return
	}

	data, err := uc.userService.GetUserExist(loginData.Username)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			c.JSON(http.StatusNotFound, gin.H{"message": "user not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"message": "server error"})
		return
	}
	fmt.Println(data.Password)
	if err := bcrypt.CompareHashAndPassword([]byte(data.Password), []byte(loginData.Password)); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "wrong password"})
		return
	}

	tokenString, err := uc.jwtService.CreateToken(loginData.Username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "server error"})
	}

	authResponse := types.AuthResponse{
		Username:    data.Username,
		Name:        data.Name,
		AccessToken: tokenString,
	}

	c.JSON(http.StatusOK, gin.H{"message": "Login successful", "data": authResponse})
}
