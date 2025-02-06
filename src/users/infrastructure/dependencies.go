package infrastructure

import (
	"chat/src/users/application"
	"chat/src/users/infrastructure/controllers"

	"github.com/gin-gonic/gin"
)

func InitUsers(db *MySQL, router *gin.Engine) {
	println("CARGANDO DEPENDENCIAS DE USUARIOS")

	userSaver := application.NewSaveUser(db)
	userRemover := application.NewDeleteUser(db)
	userViewer := application.NewViewUsers(db)
	userView := application.NewViewUser(db)

	addUserController := controllers.NewSaveUserController(userSaver)
	deleteUserController := controllers.NewRemoveUserController(userRemover)
	viewUsersController := controllers.NewViewAllUsersController(userViewer)
	viewUserController := controllers.NewViewOneUserController(userView)

	SetupUserRoutes(router, addUserController, deleteUserController, viewUsersController, viewUserController)
}
