package infrastructure

import (
	"log"

	"chat/src/core"
	"chat/src/users/application"
	"chat/src/users/infrastructure/controllers"

	"github.com/gin-gonic/gin"
)

func InitUsers(db *core.ConnMySQL, router *gin.Engine) {
	log.Println("CARGANDO DEPENDENCIAS DE USUARIOS")

	userRepo := NewUserRepository(db)

	userSaver := application.NewSaveUser(userRepo)
	userRemover := application.NewDeleteUser(userRepo)
	userViewer := application.NewViewUsers(userRepo)
	userView := application.NewViewUser(userRepo)

	addUserController := controllers.NewSaveUserController(userSaver)
	deleteUserController := controllers.NewDeleteUserController(userRemover)
	viewUsersController := controllers.NewViewAllUsersController(userViewer)
	viewUserController := controllers.NewViewOneUserController(userView)

	SetupUserRoutes(router, addUserController, deleteUserController, viewUsersController, viewUserController)
}
