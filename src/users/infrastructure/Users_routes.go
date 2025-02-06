package infrastructure

import (
	"chat/src/users/infrastructure/controllers"

	"github.com/gin-gonic/gin"
)

// SetupUserRoutes configura las rutas para la entidad "Users"
func SetupUserRoutes(
	router *gin.Engine,
	addUserController *controllers.AddUserController,
	deleteUserController *controllers.RemoveUserController,
	viewUsersController *controllers.ViewAllUsersController,
	viewUserController *controllers.ViewOneUserController,
) {
	userGroup := router.Group("/users")
	{
		userGroup.POST("", addUserController.Run)
		userGroup.GET("", viewUsersController.Run)
		userGroup.GET("/:id", viewUserController.Run)
		userGroup.DELETE("/:id", deleteUserController.Run)
	}
}
