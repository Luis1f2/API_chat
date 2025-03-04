package infrastructure

import (
	"log"

	"chat/src/Users/infrastructure/controllers"

	"github.com/gin-gonic/gin"
)

func SetupUserRoutes(
	router *gin.Engine,
	saveUserController *controllers.SaveUserController,
	deleteUserController *controllers.DeleteUserController,
	viewUsersController *controllers.ViewAllUsersController,
	viewOneUserController *controllers.ViewOneUserController,
	loginUserscontroller *controllers.LoginUserController,
) {
	userGroup := router.Group("/users")
	{
		userGroup.POST("", saveUserController.Run)
		userGroup.GET("", viewUsersController.Run)
		userGroup.GET("/:id", viewOneUserController.Run)
		userGroup.DELETE("/:id", deleteUserController.Run)
		userGroup.POST("/login", loginUserscontroller.Run)

	}

	log.Println("Rutas de usuario configuradas")
}
