// routes/routes.go

package routes

import (
	"log"
	"time"

	"github.com/fatih/color"
	"github.com/gabriel-tama/be-week-1/api/v1/controllers"
	"github.com/gabriel-tama/be-week-1/api/v1/services"
	"github.com/gin-gonic/gin"

	"go.uber.org/ratelimit"
)

var (
	limit ratelimit.Limiter
)

type RouteParam struct {
	UserController    *controllers.UserController
	BankController    *controllers.BankController
	ProductController *controllers.ProductController
	ImageController   *controllers.ImageController
	JwtService        *services.JWTService
	S3Service         *services.S3Service
}

func leakBucket() gin.HandlerFunc {
	prev := time.Now()
	return func(ctx *gin.Context) {
		now := limit.Take()
		log.Print(color.CyanString("%v", now.Sub(prev)))
		prev = now
	}
}

func SetupRouter(param RouteParam) *gin.Engine {
	limit = ratelimit.New(100)

	router := gin.Default()

	router.Use(leakBucket())
	router.SetTrustedProxies([]string{"::1"}) // This is for reverse proxy

	router.Use(gin.Recovery())

	// Setup API version 1 routes
	v1 := router.Group("/api/v1")
	{
		// Setup  routes
		SetupUserRoutes(v1, param.UserController)
		SetupBankRoutes(v1, param.BankController, param.JwtService)
		SetupProductRoutes(v1, param.ProductController, param.JwtService)
		SetupImageRoutes(v1, param.ImageController, param.JwtService, param.S3Service)
	}
	router.GET("/rate", func(ctx *gin.Context) {
		ctx.JSON(200, "rate limiting test")
	})

	return router
}
