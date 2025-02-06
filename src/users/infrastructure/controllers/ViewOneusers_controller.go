package controllers

import (
	"net/http"
	"strconv"

	"chat/src/users/application"

	"github.com/gin-gonic/gin"
)

type ViewOneUserController struct {
	userView *application.ViewUser
}

func NewViewOneUserController(useCase *application.ViewUser) *ViewOneUserController {
	return &ViewOneUserController{userView: useCase}
}

func (vu *ViewOneUserController) Run(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid User ID"})
		return
	}

	user, err := vu.userView.Execute(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"user": user})
}