package routes

import (
	"github.com/gabriel-tama/be-week-1/api/v1/controllers"
	"github.com/gabriel-tama/be-week-1/api/v1/middlewares"
	"github.com/gabriel-tama/be-week-1/api/v1/services"
	"github.com/gin-gonic/gin"
)

func SetupProductRoutes(router *gin.RouterGroup, productController *controllers.ProductController, jwtService *services.JWTService) {
	productRouter := router.Group("/product")
	productRouter.Use(middlewares.AuthorizeJWT(*jwtService))

	{
		productRouter.POST("/",productController.CreateProduct)
		productRouter.PATCH("/:productId",productController.UpdateProduct)
		productRouter.DELETE("/:productId",productController.DeleteProduct)
	}
	
}
