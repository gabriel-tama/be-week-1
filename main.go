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

	userModel := models.NewUserModel(conn)
	bankModel := models.NewBankAccountModel(conn)

    // Initialize services
    userService := services.NewUserService(userModel)
	bankService := services.NewBankService(bankModel)
	jwtService := services.NewJWTService(secretKey)

    // Initialize controllers
    userController := controllers.NewUserController(userService,jwtService)
	bankController := controllers.NewBankController(bankService)

    // Setup Gin router
    router := routes.SetupRouter(userController,bankController,&jwtService)

	// Start the server
	if err := router.Run(fmt.Sprintf("%s:%s", env.Host, env.Port)); err != nil {
		log.Fatal("Server error:", err)
	}
}
