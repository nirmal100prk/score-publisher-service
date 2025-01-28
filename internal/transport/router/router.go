package router

import (
	"score-publisher-svc/internal/transport/websockets"

	"github.com/gin-gonic/gin"
)

func NewRouter(router *gin.Engine, wsHandler *websockets.WebSocketHandler) {

	router.POST("/ws", wsHandler.HandleWebSocket)

}
