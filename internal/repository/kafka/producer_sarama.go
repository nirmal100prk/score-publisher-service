package kafka

import (
	"fmt"
	"time"

	"github.com/IBM/sarama"
)

type KafkaProducer struct {
	Producer sarama.SyncProducer
	Topic    string
}

func NewKafkaProducer(brokers []string, topic string) (*KafkaProducer, error) {
	config := sarama.NewConfig()

	// waits for acks from all in sync replicas
	config.Producer.RequiredAcks = sarama.WaitForAll

	//  number of retry for failed attempts
	config.Producer.Retry.Max = 1

	// maximum number of messages in batch
	config.Producer.Flush.Messages = 100

	// how long producer wait before sending a batch
	config.Producer.Flush.Frequency = 1 * time.Second

	// -- GZIP compression
	config.Producer.Compression = sarama.CompressionGZIP

	//  for strict ordering
	config.Net.MaxOpenRequests = 1

	config.Producer.Return.Successes = true

	// -- Idempotent producer
	config.Producer.Idempotent = true

	producer, err := sarama.NewSyncProducer(brokers, config)
	if err != nil {
		return nil, fmt.Errorf("failed to create Sarama producer: %w", err)
	}

	return &KafkaProducer{Producer: producer, Topic: topic}, nil
}

type MessageRepository interface {
	PublishMessage(message []byte) error
}

func NewMessageRepository(kp *KafkaProducer) MessageRepository {
	return kp
}

func (kp *KafkaProducer) PublishMessage(message []byte) error {

	msg := &sarama.ProducerMessage{
		Topic:     kp.Topic,
		Partition: 0,
		Value:     sarama.ByteEncoder(message),
	}

	// Send the message and wait for ack
	partition, offset, err := kp.Producer.SendMessage(msg)
	if err != nil {
		return fmt.Errorf("failed to produce message: %w", err)
	}

	fmt.Printf("Message delivered to %s [partition=%d] offset=%d\n",
		kp.Topic, partition, offset)
	return nil
}

func (kp *KafkaProducer) CloseConnection() {
	if err := kp.Producer.Close(); err != nil {
		fmt.Printf("Error closing producer: %v\n", err)
	}
}
