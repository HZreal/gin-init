package rabbitMQ

/**
 * @Author nico
 * @Date 2024-11-16
 * @File: connection.go
 * @Description:
 */

import (
	"gin-init/config"
	amqp "github.com/rabbitmq/amqp091-go"
	"log"
)

var Conn *amqp.Connection

func InitRabbitMQ() {
	var err error
	Conn, err = amqp.Dial(config.Conf.RabbitMQ.GetUrl())
	if err != nil {
		log.Fatalf("[ERROR] Failed to connect to RabbitMQ: %v", err)
	}
}