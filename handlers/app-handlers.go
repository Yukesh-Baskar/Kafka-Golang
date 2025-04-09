package handlers

import (
	"encoding/json"
	"kafka-golang/types"
	"net/http"
	"time"

	"github.com/IBM/sarama"
	"github.com/gin-gonic/gin"
)

func HandleRoutes(incomingRouterGroup *gin.RouterGroup) {
	incomingRouterGroup.POST("/publish-message", HandlePublishMessage)
}

func HandlePublishMessage(c *gin.Context) {
	var (
		user types.User
		err  error
	)

	if err = c.ShouldBindJSON(&user); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": err.Error(),
		})
		return
	}
	val, err := json.Marshal(user)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": err.Error(),
		})
		return
	}
	mes := &sarama.ProducerMessage{
		Topic:     types.Config.KafkaBrokerConfig.TopicConfigs.Topic_Name,
		Key:       sarama.StringEncoder(types.Config.KafkaBrokerConfig.TopicConfigs.TopicKey),
		Value:     sarama.StringEncoder(val),
		Partition: 0,
		Timestamp: time.Now(),
	}

	partition, offset, err := types.KafkaProducer.SendMessage(mes)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status":    http.StatusOK,
		"message":   "Data published successfully",
		"partition": partition,
		"offset":    offset,
	})
}
