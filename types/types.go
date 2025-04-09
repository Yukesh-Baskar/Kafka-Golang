package types

import (
	"github.com/IBM/sarama"
)

var KafkaProducer sarama.SyncProducer
var Config Configs

type Configs struct {
	App               Application        `yaml:"App"`
	KafkaBrokerConfig KafkaBrokerConfigs `yaml:"KafkaBrokerConfigs"`
}

type Application struct {
	Port string `yaml:"Port"`
}

type KafkaBrokerConfigs struct {
	KafkaBrokerPorts []string           `yaml:"KafkaBrokerPorts"`
	TopicConfigs     TopicConfiguration `yaml:"TopicConfigs"`
}

type TopicConfiguration struct {
	TopicKey   string `yaml:"TopicKey"`
	Topic_Name string `yaml:"Topic_Name"`
	Partitions int    `yaml:"partitions"`
}

type User struct { // just for passing the data via broken
	UserId string `json:"user_id"`
	Name   string `json:"name"`
	Age    int    `json:"age"`
	Email  string `json:"email"`
}
