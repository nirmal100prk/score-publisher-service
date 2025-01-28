package websockets

import (
	"log"
	"net/http"
	"score-publisher-svc/internal/service"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

// WebSocketHandler handles WebSocket connections
type WebSocketHandler struct {
	messageService *service.MessageService
}

func NewWebSocketHandler(messageService *service.MessageService) *WebSocketHandler {
	return &WebSocketHandler{
		messageService: messageService,
	}
}

// HandleWebSocket upgrades the HTTP connection to a WebSocket connection
func (h *WebSocketHandler) HandleWebSocket(c *gin.Context) {
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Printf("Failed to upgrade WebSocket connection: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"message": "failed to initiate connection"})
		return
	}

	log.Println("Client connected")

	go h.handleConnection(conn)
}

func (h *WebSocketHandler) handleConnection(conn *websocket.Conn) {
	defer func() {
		conn.Close()
		log.Println("Client disconnected")
	}()

	for {
		messageType, message, err := conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("WebSocket error: %v", err)
			}
			break
		}

		log.Printf("Message received: %s", string(message))

		// publish message to kafka
		// Publish the message to Kafka using MessageService
		err = h.messageService.PublishMessage(string(message))
		if err != nil {
			log.Printf("Failed to publish message to Kafka: %v", err)
			continue
		}

		err = conn.WriteMessage(messageType, message)
		if err != nil {
			log.Printf("Failed to write message: %v", err)
			break
		}
	}
}
