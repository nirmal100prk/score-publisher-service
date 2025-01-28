package kafka

import (
	"fmt"
	"time"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

type KafkaProducer struct {
	Producer *kafka.Producer
	Topic    string
}

func NewKafkaProducer(brokers []string, topic string) (*KafkaProducer, error) {

	producer, err := kafka.NewProducer(&kafka.ConfigMap{
		"bootstrap.servers":                     brokers,
		"acks":                                  "all",
		"retries":                               0,
		"batch.num.messages":                    100,
		"linger.ms":                             1000,
		"compression.codec":                     "gzip",
		"max.in.flight.requests.per.connection": 1,
		"enable.idempotence":                    true,
		"transactional.id":                      "myid",
	})
	if err != nil {
		return nil, err
	}
	return &KafkaProducer{Producer: producer,
		Topic: topic}, nil
}

type MessageRepository interface {
	PublishMessage(message []byte) error
}

func NewMessageRepository(kp *KafkaProducer) MessageRepository {
	return kp
}

func (kp *KafkaProducer) PublishMessage(message []byte) error {

	msg := &kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &kp.Topic,
			Partition: int32(kafka.PartitionAny),
		},
		Value: message,
	}

	// Send the message
	deliveryChan := make(chan kafka.Event, 1)
	defer close(deliveryChan)

	err := kp.Producer.Produce(msg, deliveryChan)
	if err != nil {
		return fmt.Errorf("failed to produce message: %w", err)
	}

	select {
	case ev := <-deliveryChan:
		m := ev.(*kafka.Message)
		if m.TopicPartition.Error != nil {
			return fmt.Errorf("delivery failed: %w", m.TopicPartition.Error)
		}
		fmt.Printf("Message delivered to %s [%d] at offset %v\n",
			*m.TopicPartition.Topic, m.TopicPartition.Partition, m.TopicPartition.Offset)
	case <-time.After(10 * time.Second):
		return fmt.Errorf("delivery timeout")
	}

	return nil

}

func (kp *KafkaProducer) CloseConnection() {
	kp.Producer.Close()
}
