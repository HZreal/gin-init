package rabbitMQ

/**
 * @Author huang
 * @Date 2024-07-12
 * @File: init.go
 * @Description:
 */

import (
	"gin-init/consumers"
	"github.com/streadway/amqp"
	"log"
)

var Conn *amqp.Connection

var producer *RabbitMQProducer

var consumer *RabbitMQConsumer

/*
*
 */
func createGlobalProducer() {
	producer = NewProducer(Conn)
}

/*
*
 */
func startConsumers() {
	consumer = NewConsumer(Conn)

	// 启动消费者1
	err := consumer.Listen("queue1", consumers.HandleMessage1)
	if err != nil {
		log.Fatalf("Failed to start consumer for queue1: %v", err)
	}

	// 启动消费者2
	err = consumer.Listen("queue2", consumers.HandleMessage2)
	if err != nil {
		log.Fatalf("Failed to start consumer for queue2: %v", err)
	}
}

func Start() {
	var err error
	Conn, err = amqp.Dial("amqp://admin:root123456@localhost:5672/%2Flocal")
	if err != nil {
		log.Fatalf("Failed to connect to RabbitMQ: %v", err)
	}
	log.Println("Connected to RabbitMQ")

	//
	createGlobalProducer()

	//
	startConsumers()
}
