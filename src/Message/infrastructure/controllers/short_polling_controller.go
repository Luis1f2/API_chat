package controllers

import (
	"chat/src/Message/application"
	"chat/src/Message/domain/entities"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ShortPollingController struct {
	messageViewer *application.ViewMessages
}

func NewShortPollingController(useCase *application.ViewMessages) *ShortPollingController {
	return &ShortPollingController{messageViewer: useCase}
}

func (spc *ShortPollingController) Run(c *gin.Context) {
	userID, err := strconv.Atoi(c.Query("user_id"))

	var messages []entities.Message
	if err == nil && userID > 0 {
		messages, err = spc.messageViewer.ExecuteForUser(userID)
	} else {
		messages, err = spc.messageViewer.ExecuteAll()
	}

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al obtener mensajes"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"messages": messages})
}
