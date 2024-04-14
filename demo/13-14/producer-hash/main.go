package main

import (
	"encoding/json"
	"github.com/Shopify/sarama"
	"log"
	"strconv"
)

func main() {
	config := sarama.NewConfig()
	config.Consumer.Retry.Backoff = 0
	config.Producer.Retry.Backoff = 0
	config.Version = sarama.V2_0_0_0
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Return.Successes = true
	config.Producer.Retry.Max = 5
	config.Net.MaxOpenRequests = 1
	//设置消息分区策略
	config.Producer.Partitioner = sarama.NewHashPartitioner
	brokers := []string{"localhost:9092"}
	producer, err := sarama.NewSyncProducer(brokers, config)
	if err != nil {
		log.Fatal("Error creating producer: ", err)
	}

	user := struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	}{
		ID:   1,
		Name: "张三",
	}
	userBytes, _ := json.Marshal(user)

	msg := &sarama.ProducerMessage{
		Topic: "test-topic",
		Value: sarama.ByteEncoder(userBytes),
		//指定消息的key
		Key: sarama.StringEncoder(strconv.Itoa(user.ID)),
	}
	partition, offset, err := producer.SendMessage(msg)
	if err != nil {
		log.Println(err, partition, offset)
	}

}
