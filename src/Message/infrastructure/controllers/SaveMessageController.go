package controllers

import (
	"chat/src/Message/application"
	"net/http"

	"github.com/gin-gonic/gin"
)

type SaveMessageController struct {
	messageSaver *application.SaveMessage
}

func NewSaveMessageController(useCase *application.SaveMessage) *SaveMessageController {
	return &SaveMessageController{messageSaver: useCase}
}

func (mc *SaveMessageController) Run(c *gin.Context) {
	var body struct {
		SenderID   int    `json:"sender_id"`
		ReceiverID int    `json:"receiver_id"`
		Content    string `json:"content"`
	}

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Formato JSON inválido"})
		return
	}

	err := mc.messageSaver.Execute(body.SenderID, body.ReceiverID, body.Content)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al guardar el mensaje"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Mensaje guardado correctamente"})
}
