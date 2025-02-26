package controllers

import (
	"chat/src/Users/application"
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

	if len(users) == 0 {
		c.JSON(http.StatusOK, gin.H{"message": "No hay usuarios registrados"})
		return
	}

	response := make([]gin.H, len(users))
	for i, user := range users {
		response[i] = gin.H{
			"id":       user.ID,
			"username": user.Username,
		}
	}

	c.JSON(http.StatusOK, gin.H{"users": response})
}
