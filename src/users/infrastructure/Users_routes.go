package infrastructure

import (
	"log"

	"chat/src/Users/infrastructure/controllers"

	"github.com/gin-gonic/gin"
)

// SetupUserRoutes configura las rutas para la entidad "Users"
func SetupUserRoutes(
	router *gin.Engine,
	saveUserController *controllers.SaveUserController,
	deleteUserController *controllers.DeleteUserController,
	viewUsersController *controllers.ViewAllUsersController,
	viewOneUserController *controllers.ViewOneUserController,
) {
	userGroup := router.Group("/users")
	{
		userGroup.POST("", saveUserController.Run)         // Crear usuario
		userGroup.GET("", viewUsersController.Run)         // Obtener todos los usuarios
		userGroup.GET("/:id", viewOneUserController.Run)   // Obtener un usuario por ID
		userGroup.DELETE("/:id", deleteUserController.Run) // Eliminar usuario por ID
	}

	log.Println("Rutas de usuario configuradas")
}
