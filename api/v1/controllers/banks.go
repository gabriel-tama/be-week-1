package controllers

import (
	"net/http"

	"github.com/gabriel-tama/be-week-1/api/v1/services"
	"github.com/gin-gonic/gin"
)

type BankController struct {
	bankService services.BankService
}


func NewBankController(bankService services.BankService) *BankController {
    return &BankController{bankService: bankService}
}

func (bc *BankController) Check(c *gin.Context) {
    err:=bc.bankService.CheckInBank()
	if err!=nil{
		c.JSON(http.StatusBadRequest, gin.H{"message": "server error"})
		return
	}
	c.JSON(http.StatusOK,gin.H{"meesage":"ur ok!"})
	return
}