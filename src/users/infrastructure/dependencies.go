package infrastructure

import (
	"log"

	"chat/src/Users/application"
	"chat/src/Users/infrastructure/controllers"
	"chat/src/core"

	"github.com/gin-gonic/gin"
)

func InitUsers(db *core.ConnMySQL, router *gin.Engine) {
	log.Println("CARGANDO DEPENDENCIAS DE USUARIOS")

	userRepo := NewUserRepository(db)

	userSaver := application.NewSaveUser(userRepo)
	userRemover := application.NewDeleteUser(userRepo)
	userViewer := application.NewViewUsers(userRepo)
	userView := application.NewViewUser(userRepo)
	loginUser := application.NewLoginUser(userRepo)

	addUserController := controllers.NewSaveUserController(userSaver)
	deleteUserController := controllers.NewDeleteUserController(userRemover)
	viewUsersController := controllers.NewViewAllUsersController(userViewer)
	viewUserController := controllers.NewViewOneUserController(userView)
	loginUserController := controllers.NewLoginUserController(loginUser)

	SetupUserRoutes(router, addUserController, deleteUserController, viewUsersController, viewUserController, loginUserController)
}
