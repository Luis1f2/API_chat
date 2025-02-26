package infrastructure

import (
	"chat/src/Message/infrastructure/controllers"

	"github.com/gin-gonic/gin"
)

func SetupMessageRoutes(router *gin.Engine, saveMessageController *controllers.SaveMessageController, viewMessageController *controllers.ViewMessageController) {
	messageGroup := router.Group("/messages")
	{
		messageGroup.POST("", saveMessageController.Run)
		messageGroup.GET("/:id", viewMessageController.Run)
	}
}
