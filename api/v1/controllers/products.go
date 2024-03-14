package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gabriel-tama/be-week-1/api/v1/services"
	"github.com/gabriel-tama/be-week-1/types"
	"github.com/gin-gonic/gin"
)

type ProductController struct {
	productService services.ProductService
	jwtService     services.JWTService
}

func NewProductController(productService services.ProductService, jwtService services.JWTService) *ProductController {
	return &ProductController{productService: productService, jwtService: jwtService}
}

func (pc *ProductController) CreateProduct (c *gin.Context){
	var product types.ProductCreate
	if err := c.ShouldBindJSON(&product); err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"message": "required details are missing or invalid"})
		return
	}
	authHeader := c.GetHeader("Authorization")
	
	const BEARER_SCHEMA = "BEARER "
	
	tokenString:= authHeader[len(BEARER_SCHEMA):]
	
	userID,err := pc.jwtService.GetUserIDByToken(tokenString)

	if err!=nil{
		c.JSON(http.StatusInternalServerError,gin.H{"message":"server error"})
	}

	convertedUserID, err := strconv.ParseUint(userID, 10, 64)
	if err!=nil{
		c.JSON(http.StatusInternalServerError,gin.H{"message":err})
		return
	}

	err = pc.productService.Create(int(convertedUserID),product)
	
	if err!=nil{
		c.JSON(http.StatusInternalServerError,gin.H{"message":err})
	}

	c.JSON(http.StatusOK, gin.H{"message":"account added sucessfully","data": product})

}
