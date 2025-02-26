package controllers

import (
	"chat/src/Message/application"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ViewMessageController struct {
	messageViewer *application.ViewMessage
}

func NewViewMessageController(useCase *application.ViewMessage) *ViewMessageController {
	return &ViewMessageController{messageViewer: useCase}
}

func (vc *ViewMessageController) Run(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil || id <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	message, err := vc.messageViewer.Execute(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Mensaje no encontrado"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": message})
}
