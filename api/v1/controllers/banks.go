package controllers

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gabriel-tama/be-week-1/api/v1/services"
	"github.com/gabriel-tama/be-week-1/types"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/jackc/pgx/v5"
)

type BankController struct {
	bankService services.BankService
	jwtService  services.JWTService

}


func NewBankController(bankService services.BankService, jwtService services.JWTService) *BankController {
    return &BankController{bankService: bankService, jwtService: jwtService}
}

func (bc *BankController) Check(c *gin.Context) {
    err:=bc.bankService.CheckInBank()
	if err!=nil{
		c.JSON(http.StatusBadRequest, gin.H{"message": "server error"})
		return
	}
	c.JSON(http.StatusOK,gin.H{"meesage":"ur ok!"})
}

func (bc *BankController) GetUserIDByToken(token string) string {
	aToken, err := bc.jwtService.ValidateToken(token)
	if err != nil {
		panic(err.Error())
	}
	claims := aToken.Claims.(jwt.MapClaims)
	return fmt.Sprintf("%v", claims["user_id"])
}

func (bc *BankController) CreateBankAccount (c *gin.Context){
	var bank types.RegisterBankData
	if err := c.ShouldBindJSON(&bank); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "required details are missing or invalid"})
		return
	}
	authHeader := c.GetHeader("Authorization")
	
	const BEARER_SCHEMA = "BEARER "
	
	tokenString:= authHeader[len(BEARER_SCHEMA):]
	
	userID := bc.GetUserIDByToken(tokenString)

	convertedUserID, err := strconv.ParseUint(userID, 10, 64)
	if err!=nil{
		c.JSON(http.StatusInternalServerError,gin.H{"message":err})
		return
	}

	_,err = bc.bankService.CreateBankAccount(int(convertedUserID),bank.BankName,bank.BankAccountName,bank.BankAccountNumber)
	
	if err!=nil{
		c.JSON(http.StatusInternalServerError,gin.H{"message":err})
	}

	c.JSON(http.StatusOK, gin.H{"message":"account added sucessfully","data": bank})

}

func (bc *BankController) GetBankAccount (c *gin.Context){

	authHeader := c.GetHeader("Authorization")
	
	const BEARER_SCHEMA = "BEARER "
	
	tokenString:= authHeader[len(BEARER_SCHEMA):]
	
	userID := bc.GetUserIDByToken(tokenString)

	convertedUserID, err := strconv.ParseUint(userID, 10, 64)
	if err!=nil{
		c.JSON(http.StatusInternalServerError,gin.H{"message":"server error"})
		return
	}
	var bankResponses []types.GetBankData

	bank,err := bc.bankService.GetBankAccount(int(convertedUserID))
	
	if err!=nil{
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError,gin.H{"message":"server error"})
		return
	}
	for _, account := range bank {
		bankResponse := types.GetBankData{
			BankAccountId:      account.BankAccountId,
			BankName:           account.BankName,
			BankAccountName:    account.BankAccountName,
			BankAccountNumber:  account.BankAccountNumber,
		}
		bankResponses = append(bankResponses, bankResponse)
	}
	
	c.JSON(http.StatusOK,gin.H{"message":"success","data":bankResponses})


}

func (bc *BankController) DeleteBankAccount(c *gin.Context){
	
	acc_id, err := strconv.Atoi(c.Param("bankAccountId"))
	if err != nil {
		// If the parameter cannot be converted to an integer, return a Bad Request response
		c.JSON(http.StatusNotFound, gin.H{"message": "account request"})
		return
	}
	
	
	authHeader := c.GetHeader("Authorization")

	const BEARER_SCHEMA = "BEARER "
	
	tokenString:= authHeader[len(BEARER_SCHEMA):]
	
	userID := bc.GetUserIDByToken(tokenString)

	convertedUserID, err := strconv.ParseUint(userID, 10, 64)
	if err!=nil{
		c.JSON(http.StatusInternalServerError,gin.H{"message":"server error"})
		return
	}

	err= bc.bankService.DeleteBankAccount(acc_id,int(convertedUserID))

	if err!=nil{
		if errors.Is(err,pgx.ErrNoRows){
		c.JSON(http.StatusNotFound,gin.H{"message":"account not found"})
		return
		}
		c.JSON(http.StatusInternalServerError,gin.H{"message":"server error"})
	}
	c.JSON(http.StatusOK,gin.H{"message":"account deleted successfully"})
}