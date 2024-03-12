// routes/routes.go

package routes

import (
	"github.com/gabriel-tama/be-week-1/api/v1/controllers"
	"github.com/gin-gonic/gin"
)

func SetupRouter(userController *controllers.UserController) *gin.Engine {
    router := gin.Default()

    // Setup API version 1 routes
    v1 := router.Group("/api/v1")
    {

        // Setup user routes
        SetupUserRoutes(v1, userController)
    }

    return router
}
