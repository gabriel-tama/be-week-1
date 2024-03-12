// main.go

package main

import (
	"context"
	"log"
	"os"

	"github.com/gabriel-tama/be-week-1/api/v1/controllers"
	"github.com/gabriel-tama/be-week-1/api/v1/routes"
	"github.com/gabriel-tama/be-week-1/api/v1/services"
	"github.com/gabriel-tama/be-week-1/models"
	"github.com/jackc/pgx/v5"
)

func main() {
    // Connect to the database
	// dbName := os.Getenv("DB_NAME")
    // dbPort := os.Getenv("DB_PORT")
    // dbHost := os.Getenv("DB_HOST")
    // dbUsername := os.Getenv("DB_USERNAME")
    // dbPassword := os.Getenv("DB_PASSWORD")
	secretKey := os.Getenv("JWT_SECRET")

    // Construct the connection string
    // connString := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", dbUsername, dbPassword, dbHost, dbPort, dbName)
	connString:= "postgres://postgres:gia777gia@localhost:5432/shopifyx"
    conn, err := pgx.Connect(context.Background(), connString)
    if err != nil {
        log.Fatal("Unable to connect to the database:", err)
    }
    defer conn.Close(context.Background())


	userModel := models.NewUserModel(conn)

    // Initialize services
    userService := services.NewUserService(userModel)
	jwtService := services.NewJWTService(secretKey)

    // Initialize controllers
    userController := controllers.NewUserController(userService,jwtService)

    // Setup Gin router
    router := routes.SetupRouter(userController)

    // Start the server
    if err := router.Run(":5000"); err != nil {
        log.Fatal("Server error:", err)
    }
}
