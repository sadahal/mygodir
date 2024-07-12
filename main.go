package main

import (
	"fmt"
	"log"
	"os"

	"github.com/sadahal/mygodir/kafka"
)

func main() {
	fmt.Println(os.Getenv("KAFKA_BOOTSTRAP_SERVERS"))
	kafkaConfig := kafka.LoadKafkaConfig()
	log.Println("The kafka config is ", kafkaConfig)
    // configMap := kafka.CreateKafkaConfigMap(kafkaConfig)

    // adminClient, err := kafka.InitKafkaAdminClient(configMap)
    // if err != nil {
    //     panic(err)
    // }
    // defer adminClient.Close()

    // topics, err := kafka.ListKafkaTopics(adminClient)
    // if err != nil {
    //     panic(err)
    // }

    // for _, topic := range topics {
    //     fmt.Println(topic)
    // }
}
