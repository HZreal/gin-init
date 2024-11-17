package mq

/**
 * @Author nico
 * @Date 2024-11-16
 * @File: rabbitMQ.go
 * @Description:
 */

import (
	"gin-init/config"
	amqp "github.com/rabbitmq/amqp091-go"
	"log"
)

var Conn *amqp.Connection

// var Channel *amqp.Channel

func InitRabbitMQ() {
	var err error
	Conn, err = amqp.Dial(config.Conf.RabbitMQ.GetUrl())
	if err != nil {
		log.Fatalf("[ERROR] Failed to connect to RabbitMQ: %v", err)
	}

	// Channel, err = Conn.Channel()
	// if err != nil {
	// 	log.Fatalf("[ERROR] Failed to open a channel: %v", err)
	// }
}
