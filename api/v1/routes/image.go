// routes/users.go

package routes

import (
	"github.com/gabriel-tama/be-week-1/api/v1/controllers"
	"github.com/gabriel-tama/be-week-1/api/v1/middlewares"
	"github.com/gabriel-tama/be-week-1/api/v1/services"
	"github.com/gin-gonic/gin"
)

func SetupImageRoutes(router *gin.RouterGroup, imageController *controllers.ImageController, jwtService *services.JWTService, s3Service *services.S3Service) {
	imgRouter := router.Group("/image")
	imgRouter.Use(middlewares.AuthorizeJWT(*jwtService))

	{
		imgRouter.POST("/", imageController.UploadImage)
	}
}
