package service

import (
	"encoding/json"
	"log"
	"score-publisher-svc/internal/repository/kafka"
)

type MessageService struct {
	messageRepo kafka.MessageRepository
}

func NewMessageRepository(messageRepo kafka.MessageRepository) *MessageService {
	return &MessageService{
		messageRepo: messageRepo,
	}
}

func (s *MessageService) PublishMessage(message any) error {
	msgBytes, err := json.Marshal(message)
	if err != nil {
		return err
	}
	err = s.messageRepo.PublishMessage(msgBytes)
	if err != nil {
		log.Printf("Failed to publish message: %v", err)
		return err
	}
	return nil
}
