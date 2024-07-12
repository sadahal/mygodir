package kafka

import (
	"fmt"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)
func InitKafkaAdminClient(configMap *kafka.ConfigMap) (*kafka.AdminClient, error) {
    adminClient, err := kafka.NewAdminClient(configMap)
    if err != nil {
        return nil, fmt.Errorf("failed to create admin client: %w", err)
    }
    return adminClient, nil
}

func ListKafkaTopics(adminClient *kafka.AdminClient) ([]string, error) {
    metadata, err := adminClient.GetMetadata(nil, true, 10000)
    if err != nil {
        return nil, fmt.Errorf("failed to get metadata: %w", err)
    }

    var topics []string
    for topic := range metadata.Topics {
        topics = append(topics, topic)
    }
    return topics, nil
}