package controllers

import (
	"chat/src/users/application"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AddUserController struct {
	userSaver *application.SaveUser
}

func NewSaveUserController(useCase *application.SaveUser) *AddUserController {
	return &AddUserController{userSaver: useCase}
}

func (uc *AddUserController) Run(c *gin.Context) {

	var body struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := uc.userSaver.Execute(body.Username, body.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "User saved successfully"})
}
