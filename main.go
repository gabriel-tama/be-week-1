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

	// Initialize services
	// s3Service := services.NewS3Service(env.S3ID, env.S3Secret, env.S3Bucket, "file.txt")
	userService := services.NewUserService(userModel)
	bankService := services.NewBankService(bankModel)
	jwtService := services.NewJWTService(secretKey)

	// s3Service.UploadFile(nil)

	// Initialize controllers
	userController := controllers.NewUserController(userService, jwtService)
	bankController := controllers.NewBankController(bankService, jwtService)

	// Setup Gin router
	router := routes.SetupRouter(userController, bankController, &jwtService)

	// Start the server
	if err := router.Run(fmt.Sprintf("%s:%s", env.Host, env.Port)); err != nil {
		log.Fatal("Server error:", err)
	}
}
