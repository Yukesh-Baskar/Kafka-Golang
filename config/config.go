package config

import (
	"fmt"
	"kafka-golang/handlers"
	"kafka-golang/types"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/segmentio/kafka-go"
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

	types.KafkaProducer = ConnectProducer(types.Config.KafkaBrokerConfig.KafkaBrokerPorts)
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
func ConnectProducer(brokersUrl []string) *kafka.Writer {
	return &kafka.Writer{
		Addr:        kafka.TCP(types.Config.KafkaBrokerConfig.KafkaBrokerPorts...),
		Topic:       types.Config.KafkaBrokerConfig.TopicConfigs.Topic_Name,
		Balancer:    &kafka.LeastBytes{},
		MaxAttempts: 5,
	}
}
