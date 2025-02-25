package controllers

import (
	"chat/src/users/application"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ViewAllUsersController struct {
	userViewer *application.ViewUsers
}

func NewViewAllUsersController(useCase *application.ViewUsers) *ViewAllUsersController {
	return &ViewAllUsersController{userViewer: useCase}
}

func (vc *ViewAllUsersController) Run(c *gin.Context) {
	users, err := vc.userViewer.Execute()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al obtener usuarios"})
		return
	}

	// Si no hay usuarios, devolver 204 No Content
	if len(users) == 0 {
		c.JSON(http.StatusNoContent, gin.H{"message": "No hay usuarios registrados"})
		return
	}

	// Asegurar que no se devuelvan contraseñas
	response := make([]gin.H, len(users))
	for i, user := range users {
		response[i] = gin.H{
			"id":       user.ID,
			"username": user.Username,
		}
	}

	c.JSON(http.StatusOK, gin.H{"users": response})
}
