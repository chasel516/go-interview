package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"sync"

	"github.com/Shopify/sarama"
)

func main() {
	config := sarama.NewConfig()
	config.Consumer.Return.Errors = true
	config.Version = sarama.V2_4_0_0
	config.Consumer.Offsets.Initial = sarama.OffsetOldest
	// Disable auto commit
	config.Consumer.Offsets.AutoCommit.Enable = false

	brokers := []string{"localhost:9092"}
	topic := "test-topic"
	group := "test-group"

	consumer, err := sarama.NewConsumerGroup(brokers, group, config)
	if err != nil {
		log.Fatal("Error creating consumer group: ", err)
	}

	wg := &sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			err := consumer.Consume(context.Background(), []string{topic}, &exampleConsumerGroupHandler{})
			if err != nil {
				log.Println("Error in consumer loop: ", err)
			}
		}
	}()

	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt)
	<-signalChan

	err = consumer.Close()
	if err != nil {
		log.Fatal("Error closing consumer: ", err)
	}

	wg.Wait()
}

type exampleConsumerGroupHandler struct{}

func (exampleConsumerGroupHandler) Setup(_ sarama.ConsumerGroupSession) error {
	return nil
}

func (exampleConsumerGroupHandler) Cleanup(_ sarama.ConsumerGroupSession) error {
	return nil
}

func (h exampleConsumerGroupHandler) ConsumeClaim(sess sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for msg := range claim.Messages() {
		fmt.Printf("Received message: value = %s, timestamp = %v, topic = %s, partition = %d, offset = %d\n",
			string(msg.Value), msg.Timestamp, msg.Topic, msg.Partition, msg.Offset)
		// Manually commit the offset
		sess.MarkMessage(msg, "")
	}
	return nil
}
