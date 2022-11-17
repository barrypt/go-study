package main

import (
	"fmt"

	"github.com/Shopify/sarama"
)

func NewProducer()  {
	
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Partitioner = sarama.NewHashPartitioner
	config.Producer.Return.Successes = true


	
	msg := &sarama.ProducerMessage{
		Topic: "test_topic",
		Key: sarama.StringEncoder("1"),
		Value: sarama.StringEncoder("this is a test log"),

	}
	client, err := sarama.NewSyncProducer([]string{"localhost:9092"}, config)
	if err != nil {
		fmt.Println("send msg failed, err:", err)
		return
	}
	defer client.Close()

	pid, offset, err := client.SendMessage(msg)
	if err != nil {
		fmt.Println("send msg failed, err:", err)
		return
	}

	fmt.Printf("pid: %v offset:%v\n", pid, offset)

}