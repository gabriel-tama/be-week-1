// controllers/user_controller.go

package controllers

import (
	"net/http"

	"github.com/gabriel-tama/be-week-1/api/v1/services"
	"github.com/gabriel-tama/be-week-1/types"
	"github.com/gin-gonic/gin"
)

type UserController struct {
    userService services.UserService
    jwtService  services.JWTService
}


func NewUserController(userService services.UserService,jwtService services.JWTService) *UserController {
    return &UserController{userService: userService,jwtService: jwtService}
}

func (uc *UserController) Register(c *gin.Context) {
    var user types.RegisterData
    if err := c.ShouldBindJSON(&user); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
        return
    }

    data,err := uc.userService.Register(user.Username,user.Name,user.Password)
   
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
        return
    }

    tokenString,err:= uc.jwtService.CreateToken(user.Username)
    if err !=nil{
        c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
    }

    authResponse := types.AuthResponse{
        Username: data.Username,
        Name: data.Name,
        AccessToken: tokenString,
    }

    c.JSON(http.StatusCreated, gin.H{"message": "User registered successfully","data":authResponse})
}

func (uc *UserController) Login(c *gin.Context) {
    var loginData types.LoginData
    if err := c.ShouldBindJSON(&loginData); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"message": err})
        return
    }

     data,err := uc.userService.Authenticate(loginData.Username, loginData.Password)
    
     if err != nil {
        c.JSON(http.StatusUnauthorized, gin.H{"message": err})
        return
    }
    tokenString,err:= uc.jwtService.CreateToken(loginData.Username)
    if err !=nil{
        c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
    }

    authResponse := types.AuthResponse{
        Username: data.Username,
        Name: data.Name,
        AccessToken: tokenString,
    }

    c.JSON(http.StatusOK, gin.H{"message": "Login successful","data":authResponse})
}