package infrastructure

import (
	"chat/src/Message/application"
	"chat/src/Message/infrastructure/controllers"
	"chat/src/core"
	"log"

	"github.com/gin-gonic/gin"
)

func InitMessages(db *core.ConnMySQL, router *gin.Engine) {
	log.Println("CARGANDO LAS DEPENDENCIAS PARA LOS MENSAGES")
	messageRepo := NewMessageRepository(db)

	// Crear casos de uso (use cases)
	saveMessageUseCase := application.NewSaveMessage(messageRepo)
	viewMessageUseCase := application.NewViewMessage(messageRepo)
	viewMessagesUseCase := application.NewViewMessages(messageRepo)

	// Crear controladores
	saveMessageController := controllers.NewSaveMessageController(saveMessageUseCase)
	viewMessageController := controllers.NewViewMessageController(viewMessageUseCase)
	shortPollingController := controllers.NewShortPollingController(viewMessagesUseCase)
	longPollingController := controllers.NewLongPollingController(viewMessagesUseCase)

	SetupMessageRoutes(router, saveMessageController, viewMessageController, shortPollingController, longPollingController)
}
