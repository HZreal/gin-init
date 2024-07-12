package consumers

/**
 * @Author huang
 * @Date 2024-07-12
 * @File: consumer1.go
 * @Description:
 */

import (
	"github.com/streadway/amqp"
	"log"
)

func HandleMessage1(d amqp.Delivery) {
	log.Printf("Received a message from queue1: %s", d.Body)
}
