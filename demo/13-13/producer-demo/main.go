package main

import (
	"fmt"
	"log"
	"time"

	"github.com/IBM/sarama"
)

func main() {
	config := sarama.NewConfig()
	// 等待ISR（In-Sync Replicas）列表中的全部副本同步完成才返回客户端
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Retry.Max = 5
	//启用幂等生产者（SDK版本>=1.20）
	config.Producer.Idempotent = true

	brokers := []string{"localhost:9092"}
	producer, err := sarama.NewSyncProducer(brokers, config)
	if err != nil {
		log.Fatal("Error creating producer: ", err)
	}

	defer func() {
		if err := producer.Close(); err != nil {
			log.Fatal("Error closing producer: ", err)
		}
	}()

	topic := "test-topic"
	message := &sarama.ProducerMessage{
		Topic: topic,
		Value: sarama.StringEncoder(fmt.Sprintf("Hello Kafka, it's %s", time.Now().String())),
	}

	partition, offset, err := producer.SendMessage(message)
	if err != nil {
		log.Fatal("Error sending message: ", err)
	}

	fmt.Printf("Message sent to partition %d at offset %d\n", partition, offset)
}
