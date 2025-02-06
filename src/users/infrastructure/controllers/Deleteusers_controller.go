package controllers

import (
	"chat/src/users/application"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type RemoveUserController struct {
	UserRemover *application.DeleteUser
}

func NewRemoveUserController(useCase *application.DeleteUser) *RemoveUserController {
	return &RemoveUserController{UserRemover: useCase}
}

func (ur *RemoveUserController) Run(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid User ID"})
		return
	}
	err = ur.UserRemover.Execute(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "User removed successfully"})
}
