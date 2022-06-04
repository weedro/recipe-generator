package infrastructure

import (
	"encoding/json"
	"fmt"
	"reciper/recipe-generator/internal/domain"
	"reciper/recipe-generator/internal/env"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func ProduceRecipeCreateEvent(recipe domain.Recipe) {

	broker := env.GetEnv().KafkaBroker
	topic := env.GetEnv().ReceptCreateTopic

	producer, err := kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": broker})
	if err != nil {
		fmt.Printf("Failed to create producer: %s\n", err)
		return
	}

	value, _ := json.Marshal(recipe)

	deliveryChan := make(chan kafka.Event)

	_ = producer.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
		Value:          value,
	}, deliveryChan)

	e := <-deliveryChan
	m := e.(*kafka.Message)

	if m.TopicPartition.Error != nil {
		fmt.Printf("Delivery failed: %v\n", m.TopicPartition.Error)
	} else {
		fmt.Printf("Delivered message to topic %s [%d] at offset %v\n",
			*m.TopicPartition.Topic, m.TopicPartition.Partition, m.TopicPartition.Offset)
	}

	close(deliveryChan)
}
