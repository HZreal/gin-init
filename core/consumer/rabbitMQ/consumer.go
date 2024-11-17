package rabbitMQ

/**
 * @Author huang
 * @Date 2024-07-12
 * @File: consumer.go
 * @Description:
 */

import (
	"fmt"
	"gin-init/common/constant"
	"gin-init/core/initialize/mq"
	"gin-init/job/consumerHandler"
	amqp "github.com/rabbitmq/amqp091-go"
	"log"
)

// RabbitMQConsumer 消费者
type RabbitMQConsumer struct {
	Conn *amqp.Connection
}

func NewConsumer() *RabbitMQConsumer {
	return &RabbitMQConsumer{
		Conn: mq.Conn,
	}
}

func (c *RabbitMQConsumer) Listen(topic constant.Topic, handler func(body []byte)) error {
	queueName, exchange, routingKey := topic.Queue, topic.Exchange, topic.Route

	ch, err := c.Conn.Channel()
	if err != nil {
		log.Printf("Failed to open a channel: %v", err)
		return err
	}
	// defer ch.Close()

	// 声明交换机，指定 direct 模式
	err = ch.ExchangeDeclare(
		exchange, // exchange name
		"direct", // type
		true,     // durable
		false,    // auto-deleted
		false,    // internal
		false,    // no-wait
		nil,      // arguments
	)
	if err != nil {
		return fmt.Errorf("failed to declare exchange: %v", err)
	}

	_, err = ch.QueueDeclare(
		queueName,
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Printf("Failed to declare a queue: %v", err)
		return err
	}

	_ = ch.QueueBind(queueName, routingKey, exchange, false, nil)

	msgs, err := ch.Consume(
		queueName,
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Printf("Failed to register a consumer: %v", err)
		return err
	}

	go func() {
		for d := range msgs {
			handler(d.Body)
		}
	}()

	log.Printf("Listening for messages on queue: %s", queueName)
	return nil
}

// 注册器
var registers = []struct {
	topic   constant.Topic
	handler func(msg []byte)
}{
	{topic: constant.DemoTopic, handler: consumerHandler.HandleMessage1},
	{topic: constant.FirstTopic, handler: consumerHandler.HandleMessage2},
}

// 启动所有的消费者
func StartConsumer() {
	consumer := NewConsumer()

	//
	for _, register := range registers {
		// 启动消费者
		err := consumer.Listen(register.topic, register.handler)
		if err != nil {
			log.Fatalf("Failed to start consumer for %v: %v", register.topic, err)
		}
	}

	log.Println("[INFO] RabbitMQ 消费者成功启动并监听中！！！")
}
