package routes

import (
	"github.com/gabriel-tama/be-week-1/api/v1/controllers"
	"github.com/gabriel-tama/be-week-1/api/v1/middlewares"
	"github.com/gabriel-tama/be-week-1/api/v1/services"
	"github.com/gin-gonic/gin"
)

func SetupProductRoutes(router *gin.RouterGroup, productController *controllers.ProductController, paymentController *controllers.PaymentController,jwtService *services.JWTService) {
	productRouter := router.Group("/product")
	// productRouter.Use(middlewares.AuthorizeJWT(*jwtService))

	{
		// protected route
   		productRouter.POST("/",middlewares.AuthorizeJWT(*jwtService),productController.CreateProduct)
		productRouter.PATCH("/:productId",middlewares.AuthorizeJWT(*jwtService),productController.UpdateProduct)
		productRouter.DELETE("/:productId",middlewares.AuthorizeJWT(*jwtService),productController.DeleteProduct)
		productRouter.POST("/:productId/buy",middlewares.AuthorizeJWT(*jwtService),paymentController.CreatePayment)
		productRouter.POST("/:productId/stock",middlewares.AuthorizeJWT(*jwtService),productController.UpdateStock)

		// public route
		productRouter.GET("/", productController.FindAll)
		productRouter.GET("/:productId", productController.FindById)


	}

}
