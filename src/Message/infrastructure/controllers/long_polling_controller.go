package controllers

import (
	"chat/src/Message/application"
	"net/http"
	"strconv"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
)

type LongPollingController struct {
	messageViewer *application.ViewMessages
}

var (
	messageChannels = make(map[int]chan []application.ViewMessages)
	mutex           = sync.Mutex{}
)

func NewLongPollingController(useCase *application.ViewMessages) *LongPollingController {
	return &LongPollingController{messageViewer: useCase}
}

func (lpc *LongPollingController) Run(c *gin.Context) {
	userID, err := strconv.Atoi(c.Query("user_id"))

	// Si no hay `user_id`, tratamos como Long Polling general
	if err != nil || userID <= 0 {
		userID = 0
	}

	mutex.Lock()
	if _, exists := messageChannels[userID]; !exists {
		messageChannels[userID] = make(chan []application.ViewMessages, 1)
	}
	messageChan := messageChannels[userID]
	mutex.Unlock()

	select {
	case messages := <-messageChan:
		c.JSON(http.StatusOK, gin.H{"messages": messages})
	case <-time.After(30 * time.Second): // Timeout de 30 segundos
		c.JSON(http.StatusNoContent, gin.H{"message": "No hay mensajes nuevos"})
	}
}

// Función para notificar nuevos mensajes
func NotifyNewMessage(userID int, messages []application.ViewMessages) {
	mutex.Lock()
	if ch, exists := messageChannels[userID]; exists {
		ch <- messages
	}
	if ch, exists := messageChannels[0]; exists { // Notificar también a los de la versión general
		ch <- messages
	}
	mutex.Unlock()
}
