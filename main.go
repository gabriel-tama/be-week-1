// main.go

package main

import (
	"context"
	"fmt"
	"log"

	"github.com/gabriel-tama/be-week-1/api/v1/controllers"
	"github.com/gabriel-tama/be-week-1/api/v1/routes"
	"github.com/gabriel-tama/be-week-1/api/v1/services"
	C "github.com/gabriel-tama/be-week-1/config"
	psql "github.com/gabriel-tama/be-week-1/lib"
	"github.com/gabriel-tama/be-week-1/models"
)

func main() {

	env, err := C.Get()

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	secretKey := env.JWTSecret

	dbErr := psql.Init(context.Background())

	if dbErr != nil {
		log.Fatal(dbErr)
	}
	defer psql.Close(context.Background())

	userModel := models.NewUserModel(psql.PostgresConn)
	bankModel := models.NewBankAccountModel(psql.PostgresConn)
	productModel := models.NewProductModel(psql.PostgresConn)

	// Initialize services
	s3Service := services.NewS3Service(env.S3Region, env.S3ID, env.S3Secret, env.S3Bucket, env.S3Url)
	userService := services.NewUserService(userModel)
	bankService := services.NewBankService(bankModel)
	productService := services.NewProductService(productModel)
	jwtService := services.NewJWTService(secretKey)

	// Initialize controllers
	userController := controllers.NewUserController(userService, jwtService)
	bankController := controllers.NewBankController(bankService, jwtService)
	productController := controllers.NewProductController(productService, jwtService)
	imageController := controllers.NewImageController(jwtService, s3Service)

	// Setup Gin router
	router := routes.SetupRouter(routes.RouteParam{
		JwtService:        &jwtService,
		S3Service:         &s3Service,
		UserController:    userController,
		BankController:    bankController,
		ProductController: productController,
		ImageController:   imageController,
	})

	// Start the server
	if err := router.Run(fmt.Sprintf("%s:%s", env.Host, env.Port)); err != nil {
		log.Fatal("Server error:", err)
	}
}
