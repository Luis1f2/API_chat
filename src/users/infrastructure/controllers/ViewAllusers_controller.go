package controllers

import (
	"chat/src/users/application"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ViewAllUsersController struct {
	usersViewer *application.ViewUsers
}

func NewViewAllUsersController(useCase *application.ViewUsers) *ViewAllUsersController {
	return &ViewAllUsersController{usersViewer: useCase}
}

func (vu *ViewAllUsersController) Run(c *gin.Context) {
	users, err := vu.usersViewer.Execute()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"users": users})
}
