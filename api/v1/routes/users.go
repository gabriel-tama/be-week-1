// routes/users.go

package routes

import (
	"github.com/gabriel-tama/be-week-1/api/v1/controllers"
	"github.com/gin-gonic/gin"
)

func SetupUserRoutes(router *gin.RouterGroup, userController *controllers.UserController) {
    // Define routes related to user management
    userRouter := router.Group("/user")
    {
        userRouter.POST("/register", userController.Register)
        userRouter.POST("/login", userController.Login)
    }
}
