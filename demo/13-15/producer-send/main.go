package main

import (
	"encoding/json"
	"github.com/Shopify/sarama"
	"log"
	"strconv"
	"sync"
	"time"
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
	//实例化一个异步生产者
	producer, err := sarama.NewAsyncProducer(brokers, config)
	if err != nil {
		log.Fatal("Error creating producer: ", err)
	}
	defer func() {
		closeProducer(producer)
	}()
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

	wg := &sync.WaitGroup{}
	for i := 0; i < 10; i++ {
		//异步发送消息
		wg.Add(1)
		producer.Input() <- msg
	}
	for i := 0; i < 10; i++ {
		select {
		case msg := <-producer.Errors():
			log.Println(msg.Err)
		case msg := <-producer.Successes():
			log.Printf("Message sent to partition %d, offset %d\n", msg.Partition, msg.Offset)
		case <-time.After(time.Second):
			log.Printf("Timeout waiting for msg #%d", i)
		}
		wg.Done()
	}

	//等待所有消息发送完成
	wg.Wait()

}
func closeProducer(p sarama.AsyncProducer) {
	var wg sync.WaitGroup

	//AsyncClose触发生产者的关闭。当Errors和Successes通道都关闭时，关闭过程才算完成。
	//当调用AsyncClose时，必须继续从这些通道中读取，以便清空所有正在传输中的消息。
	p.AsyncClose()

	wg.Add(2)
	go func() {
		for range p.Successes() {
			log.Println("Unexpected message on Successes()")
		}
		wg.Done()
	}()
	go func() {
		for msg := range p.Errors() {
			log.Println(msg.Err)
		}
		wg.Done()
	}()
	wg.Wait()
}
