package kafka

import (
	"os"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)
type KafkaConfig struct {
    BootstrapServers string
    Username         string
    Password         string
    SecurityProtocol string
    SaslMechanisms   string
}

func LoadKafkaConfig() KafkaConfig {
    return KafkaConfig{
        BootstrapServers: os.Getenv("KAFKA_BOOTSTRAP_SERVERS"),
        Username:         os.Getenv("KAFKA_USERNAME"),
        Password:         os.Getenv("KAFKA_PASSWORD"),
        SecurityProtocol: os.Getenv("KAFKA_SECURITY_PROTOCOL"),
        SaslMechanisms:   os.Getenv("KAFKA_SASL_MECHANISMS"),
    }
}

func CreateKafkaConfigMap(config KafkaConfig) *kafka.ConfigMap {
    return &kafka.ConfigMap{
        "bootstrap.servers": config.BootstrapServers,
        "security.protocol": config.SecurityProtocol,
        "sasl.mechanisms":   config.SaslMechanisms,
        "sasl.username":     config.Username,
        "sasl.password":     config.Password,
    }
}