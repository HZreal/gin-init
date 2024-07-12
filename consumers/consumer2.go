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

func HandleMessage2(d amqp.Delivery) {
	log.Printf("Received a message from queue2: %s", d.Body)
}
