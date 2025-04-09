package config

import (
	"fmt"
	"kafka-golang/handlers"
	"kafka-golang/types"
	"log"
	"os"

	"github.com/IBM/sarama"
	"github.com/gin-gonic/gin"
	"gopkg.in/yaml.v3"
)

func StartApp() {
	// var err error
	f, err := os.ReadFile("config/config.yaml")
	if err != nil {
		log.Fatalf("Error while reading config yaml file data: %v", err)
	}
	if err = yaml.Unmarshal(f, &types.Config); err != nil {
		log.Fatalf("Error while unmarshalling config yaml file data: %v", err)
	}

	types.KafkaProducer, err = ConnectProducer(types.Config.KafkaBrokerConfig.KafkaBrokerPorts)
	if err != nil {
		log.Fatalf("Error while connecting to producer with the given broken url's config: %v", err)
	}

	router := gin.Default()

	routerGroup := router.Group("/api/v1")
	handlers.HandleRoutes(routerGroup)

	if err = router.Run(types.Config.App.Port); err != nil {
		log.Fatalf("Error while starting server with port: %v", types.Config.App.Port)
	} else {
		fmt.Printf("server running on PORT: %s\n", types.Config.App.Port)
	}
}
func ConnectProducer(brokersUrl []string) (sarama.SyncProducer, error) {
	config := sarama.NewConfig()
	config.Producer.Return.Successes = true // Enable success notifications
	// config.Version = sarama.V4_0_0_0        // Specify the Kafka version you're using
	producer, err := sarama.NewSyncProducer(types.Config.KafkaBrokerConfig.KafkaBrokerPorts, config)
	if err != nil {
		return nil, err
	}
	return producer, nil
}
