package rabbitMQ

import (
	"github.com/streadway/amqp"
	"log"
)

/**
 * @Author huang
 * @Date 2024-07-12
 * @File: consumer.go
 * @Description:
 */

type Consumer interface {
	Listen(queueName string, handler func(amqp.Delivery)) error
}

type RabbitMQConsumer struct {
	Conn *amqp.Connection
}

func NewConsumer(conn *amqp.Connection) *RabbitMQConsumer {
	return &RabbitMQConsumer{Conn: conn}
}

func (c *RabbitMQConsumer) Listen(queueName string, handler func(amqp.Delivery)) error {
	ch, err := c.Conn.Channel()
	if err != nil {
		log.Printf("Failed to open a channel: %v", err)
		return err
	}
	defer ch.Close()

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
			handler(d)
		}
	}()

	log.Printf("Listening for messages on queue: %s", queueName)
	return nil
}
