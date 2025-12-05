package infrastructure

import (
	"context"
	"fmt"
	"log"

	"github.com/IBM/sarama"
	"github.com/ezflow-me/identity-management-service/src/contexts/shared/domain"
)

// KafkaEventBus implements EventBus using Kafka as the message broker
type KafkaEventBus struct {
	producer sarama.SyncProducer
}

// KafkaConfig holds the configuration for Kafka connection
type KafkaConfig struct {
	// Host is the Kafka server address (e.g., "192.168.1.100:9092" or "kafka.example.com:9092")
	Host     string
	ClientID string
}

// NewKafkaEventBus creates a new KafkaEventBus connected to the specified Kafka server
func NewKafkaEventBus(config KafkaConfig) (*KafkaEventBus, error) {
	if config.Host == "" {
		return nil, fmt.Errorf("kafka host is required")
	}

	saramaConfig := sarama.NewConfig()
	saramaConfig.ClientID = config.ClientID
	saramaConfig.Producer.RequiredAcks = sarama.WaitForAll
	saramaConfig.Producer.Retry.Max = 3
	saramaConfig.Producer.Return.Successes = true

	brokers := []string{config.Host}
	producer, err := sarama.NewSyncProducer(brokers, saramaConfig)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to Kafka at %s: %w", config.Host, err)
	}

	log.Printf("Connected to Kafka at %s", config.Host)

	return &KafkaEventBus{
		producer: producer,
	}, nil
}

// Publish publishes domain events to Kafka
func (bus *KafkaEventBus) Publish(ctx context.Context, events ...domain.DomainEvent) error {
	for _, event := range events {
		jsonData, err := event.ToJSON()
		if err != nil {
			return fmt.Errorf("failed to serialize event %s: %w", event.EventName(), err)
		}

		message := &sarama.ProducerMessage{
			Topic: event.Topic(),
			Key:   sarama.StringEncoder(event.AggregateID()),
			Value: sarama.ByteEncoder(jsonData),
			Headers: []sarama.RecordHeader{
				{
					Key:   []byte("event_id"),
					Value: []byte(event.EventID()),
				},
				{
					Key:   []byte("event_name"),
					Value: []byte(event.EventName()),
				},
				{
					Key:   []byte("occurred_on"),
					Value: []byte(event.OccurredOn().Format("2006-01-02T15:04:05.999999999Z07:00")),
				},
			},
		}

		partition, offset, err := bus.producer.SendMessage(message)
		if err != nil {
			return fmt.Errorf("failed to publish event %s to topic %s: %w", event.EventName(), event.Topic(), err)
		}

		log.Printf("Event %s published to topic %s [partition=%d, offset=%d]",
			event.EventName(), event.Topic(), partition, offset)
	}

	return nil
}

// Close closes the Kafka producer connection
func (bus *KafkaEventBus) Close() error {
	if bus.producer != nil {
		return bus.producer.Close()
	}
	return nil
}
