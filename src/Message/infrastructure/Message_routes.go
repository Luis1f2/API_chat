package infrastructure

import (
	"chat/src/Message/infrastructure/controllers"

	"github.com/gin-gonic/gin"
)

func SetupMessageRoutes(
	router *gin.Engine,
	saveMessageController *controllers.SaveMessageController,
	viewMessageController *controllers.ViewMessageController,
	shortPollingController *controllers.ShortPollingController,
	longPollingController *controllers.LongPollingController,
) {
	messageGroup := router.Group("/messages")
	{
		messageGroup.POST("", saveMessageController.Run)
		messageGroup.GET("/:id", viewMessageController.Run)
		messageGroup.GET("/poll", shortPollingController.Run)
		messageGroup.GET("/longpoll", longPollingController.Run)
	}
}
