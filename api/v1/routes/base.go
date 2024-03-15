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

func leakBucket() gin.HandlerFunc {
	prev := time.Now()
	return func(ctx *gin.Context) {
		now := limit.Take()
		log.Print(color.CyanString("%v", now.Sub(prev)))
		prev = now
	}
}

func SetupRouter(userController *controllers.UserController, bankController *controllers.BankController, productController *controllers.ProductController, jwtService *services.JWTService) *gin.Engine {
	limit = ratelimit.New(100)

	router := gin.Default()

	router.Use(leakBucket())
	router.SetTrustedProxies([]string{"::1"}) // This is for reverse proxy

	// Setup API version 1 routes
	v1 := router.Group("/api/v1")
	{
		// Setup  routes
		SetupUserRoutes(v1, userController)
		SetupBankRoutes(v1, bankController, jwtService)
		SetupProductRoutes(v1, productController, jwtService)
	}
	router.GET("/rate", func(ctx *gin.Context) {
		ctx.JSON(200, "rate limiting test")
	})

	return router
}
