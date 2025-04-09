package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"kafka-golang/types"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/segmentio/kafka-go"
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
	fmt.Println("types.Config.KafkaBrokerConfig.TopicConfigs.TopicKey", (types.Config.KafkaBrokerConfig.TopicConfigs))
	fmt.Println(val)
	err = types.KafkaProducer.WriteMessages(context.Background(), kafka.Message{
		Partition: 0,
		Key:       []byte(types.Config.KafkaBrokerConfig.TopicConfigs.TopicKey),
		Value:     val,
		Time:      time.Now(),
	})

	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "Data published successfully",
	})
}
