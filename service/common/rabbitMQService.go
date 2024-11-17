package common

/**
 * @Author elastic·H
 * @Date 2024-10-09
 * @File: rabbitMQService.go
 * @Description:
 */

import (
	"gin-init/common/constant"
	"gin-init/core/initialize/mq/rabbitMQ"
	amqp "github.com/rabbitmq/amqp091-go"
	"log"
	"sync"
)

type RabbitMQService struct {
	conn    *amqp.Connection
	channel *amqp.Channel
	once    sync.Once
}

// NewRabbitMQService 初始化 RabbitMQService
func NewRabbitMQService() *RabbitMQService {
	var rabbitMQService = &RabbitMQService{
		conn: rabbitMQ.Conn,
	}

	Channel, err := rabbitMQService.conn.Channel()
	if err != nil {
		log.Printf("[ERROR] Failed to open a channel: %v", err)
	}
	rabbitMQService.channel = Channel

	return rabbitMQService
}

func (r *RabbitMQService) connect() error {
	// var err error
	// r.once.Do(func() {
	// 	r.conn, err = amqp.Dial(r.url)
	// 	if err != nil {
	// 		log.Fatalf("Failed to connect to RabbitMQ: %v", err)
	// 	}
	// 	r.channel, err = r.conn.Channel()
	// 	if err != nil {
	// 		log.Fatalf("Failed to open a channel: %v", err)
	// 	}
	// })
	// return err
	rabbitMQ.InitRabbitMQ()
	return nil
}

func (r *RabbitMQService) reconnect() {
	r.once = sync.Once{} // 重新设置 once 以便重新连接
	if r.channel != nil {
		r.channel.Close()
	}
	if r.conn != nil {
		r.conn.Close()
	}
	if err := r.connect(); err != nil {
		log.Fatalf("Failed to reconnect to RabbitMQ: %v", err)
	}
}

func (r *RabbitMQService) GetChannel() {
	if r.conn == nil {
		rabbitMQ.InitRabbitMQ()
		r.conn = rabbitMQ.Conn
	}

	channel, err := r.conn.Channel()
	if err != nil {
		log.Printf("Failed to open a channel: %v", err)
	}
	r.channel = channel
}

func (r *RabbitMQService) Publish(exchange, routingKey string, body []byte) error {
	// 获取复用 channel，若没有新建
	// 生产者发送消息可以复用 channel
	if r.channel == nil {
		r.GetChannel()
	}

	// 声明交换机，指定 direct 模式
	err := r.channel.ExchangeDeclare(
		exchange, // exchange name
		"direct", // type
		true,     // durable
		false,    // auto-deleted
		false,    // internal
		false,    // no-wait
		nil,      // arguments
	)
	if err != nil {
		log.Printf("failed to declare exchange: %v", err)
		return err
	}

	// 发布消息
	err = r.channel.Publish(
		exchange,   // exchange
		routingKey, // routing key
		false,      // mandatory
		false,      // immediate
		amqp.Publishing{
			ContentType:  "text/plain",
			Body:         body,
			DeliveryMode: amqp.Persistent, // 持久化消息
			Priority:     0,               // 消息优先级
		},
	)
	if err != nil {
		log.Printf("Failed to publish a message: %v", err)
		return err
	}

	log.Printf("Message sent: %s", body)
	return nil
}

// ////////////////////////////////////// 以下是业务相关的发送函数 //////////////////////////////////////////////
func (r *RabbitMQService) SendMQMsq111(body []byte) {
	topic := constant.DemoTopic
	_ = r.Publish(topic.Exchange, topic.Route, body)
}

func (r *RabbitMQService) SendMQMsg222(body []byte) {
	topic := constant.FirstTopic
	_ = r.Publish(topic.Exchange, topic.Route, body)
}
