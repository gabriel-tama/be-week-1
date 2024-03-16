package controllers

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gabriel-tama/be-week-1/api/v1/services"
	"github.com/gabriel-tama/be-week-1/types"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgconn"
)

type PaymentController struct {
	paymentService services.PaymentService
	jwtService     services.JWTService
}

func NewPaymentController(paymentService services.PaymentService, jwtService services.JWTService) *PaymentController {
	return &PaymentController{paymentService: paymentService, jwtService: jwtService}
}

func (pc *PaymentController) CreatePayment(c *gin.Context) {
	var pgErr *pgconn.PgError
	var paymentCreate types.PaymentCreate
	productId, err := strconv.Atoi(c.Param("productId"))

	if err !=nil{
		c.JSON(http.StatusNotFound,gin.H{"message":"product not found"})
		return
	}

	if err := c.ShouldBindJSON(&paymentCreate); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "required details are missing or invalid"})
		return
	}

	authHeader := c.GetHeader("Authorization")
	
	const BEARER_SCHEMA = "BEARER "
	
	tokenString:= authHeader[len(BEARER_SCHEMA):]
	userID,err := pc.jwtService.GetUserIDByToken(tokenString)	

	if err!=nil{
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError,gin.H{"message":"server error"})
		return
	}

	convertedUserID, err := strconv.ParseUint(userID, 10, 64)
	if err!=nil{
		c.JSON(http.StatusInternalServerError,gin.H{"message":"server error"})
		return
	}

	exists, err := pc.paymentService.CreatePayment(int(convertedUserID),productId,paymentCreate)

	if err!=nil{
		if errors.As(err, &pgErr) && pgErr.Code == "23503" {
			c.JSON(http.StatusBadRequest, gin.H{"message": "payments details invalid"})
			return
		}
		c.JSON(http.StatusInternalServerError,gin.H{"message":"server error"})
		return
	}

	if !exists{
		c.JSON(http.StatusNotFound,gin.H{"message":"product not found"})
		return
	}	
	
	c.JSON(http.StatusOK,gin.H{"message":"product updated successfully"})

}