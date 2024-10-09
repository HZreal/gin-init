package rabbitMQ

/**
 * @Author huang
 * @Date 2024-07-12
 * @File: producer.go
 * @Description:
 */

// import (
// 	"github.com/streadway/amqp"
// 	"log"
// )
//
// type Producer interface {
// 	Publish(queueName string, body string) error
// }
//
// type RabbitMQProducer struct {
// 	Conn *amqp.Connection
// }
//
// func NewProducer(conn *amqp.Connection) *RabbitMQProducer {
// 	return &RabbitMQProducer{Conn: conn}
// }
//
// func (p *RabbitMQProducer) Publish(queueName string, body string) error {
// 	ch, err := p.Conn.Channel()
// 	if err != nil {
// 		log.Printf("Failed to open a channel: %v", err)
// 		return err
// 	}
// 	defer ch.Close()
//
// 	_, err = ch.QueueDeclare(
// 		queueName,
// 		false,
// 		false,
// 		false,
// 		false,
// 		nil,
// 	)
// 	if err != nil {
// 		log.Printf("Failed to declare a queue: %v", err)
// 		return err
// 	}
//
// 	err = ch.Publish(
// 		"",
// 		queueName,
// 		false,
// 		false,
// 		amqp.Publishing{
// 			ContentType: "text/plain",
// 			Body:        []byte(body),
// 		})
// 	if err != nil {
// 		log.Printf("Failed to publish a message: %v", err)
// 		return err
// 	}
//
// 	log.Printf("Message sent to queue %s: %s", queueName, body)
// 	return nil
// }
