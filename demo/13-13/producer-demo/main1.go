package main

import (
	"fmt"
	"github.com/IBM/sarama"
	"log"
	"time"
)

func main() {
	config := sarama.NewConfig()
	config.Consumer.Retry.Backoff = 0
	config.Producer.Retry.Backoff = 0
	config.Version = sarama.V0_8_2_0
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Return.Successes = true
	//开启幂等生产者
	config.Producer.Idempotent = true
	//设置事务id (SDK>=1.40.0)
	config.Producer.Transaction.ID = "test"
	config.Producer.Retry.Max = 5
	config.Net.MaxOpenRequests = 1

	brokers := []string{"localhost:9092"}
	producer, err := sarama.NewSyncProducer(brokers, config)
	if err != nil {
		log.Fatal("Error creating producer: ", err)
	}

	if !producer.IsTransactional() {
		log.Println("producer is not transactional")
	}

	defer func() {
		if err := producer.Close(); err != nil {
			log.Fatal("Error closing producer: ", err)
		}
	}()

	topic := "test-topic"
	message1 := &sarama.ProducerMessage{
		Topic: topic,
		Value: sarama.StringEncoder(fmt.Sprintf("Hello Kafka, it's %s", time.Now().String())),
	}
	message2 := &sarama.ProducerMessage{
		Topic: topic,
		Value: sarama.StringEncoder(fmt.Sprintf("Hello Kafka again, it's %s", time.Now().String())),
	}

	err = producer.BeginTxn()
	if err != nil {
		return
	}
	producer.SendMessage(message1)
	producer.SendMessage(message2)
	err = producer.CommitTxn()
	if err != nil {
		log.Fatal("Error committing transaction: ", err)
	}

	fmt.Println("Messages sent in a transaction")
}
