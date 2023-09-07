package main

import (
	"rabbitmq/rabbitmq"

	"github.com/streadway/amqp"
)

func main() {

	ch := make(chan struct{})
	rabbit := rabbitmq.NewRabbitMQ("_exchange", "_route", "_queue")
	defer rabbit.Close()
	rabbit.SendMessage(rabbitmq.Message{Body: "这是一条普通消息"})
	rabbit.SendDelayMessage(rabbitmq.Message{Body: "这是一条延时5秒的消息", DelayTime: 5})

	go registerRabbitMQConsumer()
	<-ch
}

func registerRabbitMQConsumer() {
	// 新建连接
	rabbit := rabbitmq.NewRabbitMQ("yoyo_exchange", "yoyo_route", "yoyo_queue")
	// 一般来说消费者不关闭，常驻进程进行消息消费处理
	// defer rabbit.Close()

	// 执行消费
	rabbit.Consume(func(d amqp.Delivery) {
		//logger.Info("rabbitmq", zap.String("rabbitmq", string(d.Body)))
	})
}
