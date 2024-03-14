package routes

import (
	"github.com/gabriel-tama/be-week-1/api/v1/controllers"
	"github.com/gabriel-tama/be-week-1/api/v1/middlewares"
	"github.com/gabriel-tama/be-week-1/api/v1/services"
	"github.com/gin-gonic/gin"
)

func SetupBankRoutes(router *gin.RouterGroup, bankController *controllers.BankController,jwtService *services.JWTService) {
	// Define routes related to user management
	bankRouter := router.Group("/bank")
	bankRouter.Use(middlewares.AuthorizeJWT(*jwtService))
	{
		bankRouter.GET("/check", bankController.Check)
		bankRouter.POST("/account", bankController.CreateBankAccount)
		bankRouter.GET("/account", bankController.GetBankAccount)
		bankRouter.DELETE("/account/:bankAccountId",bankController.DeleteBankAccount)
		bankRouter.PATCH("/account/:bankAccountId",bankController.UpdateBankInfo)

	}
}