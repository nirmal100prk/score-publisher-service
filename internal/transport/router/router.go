package router

import (
	"score-publisher-svc/internal/transport/handler"
	"score-publisher-svc/internal/transport/websockets"
	"score-publisher-svc/pkg/middleware"

	"github.com/gin-gonic/gin"
)

func NewRouter(router *gin.Engine, wsHandler *websockets.WebSocketHandler) {

	router.GET("/ws", wsHandler.HandleWebSocket)

	router.POST("/login", handler.Login)
	router.Use(middleware.AuthMiddleware())
	{
		router.GET("protected", handler.Protected)
	}

}
