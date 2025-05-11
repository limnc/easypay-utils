/*
*

This package provides a consumer for RabbitMQ messages
It will consume messages from a queue and process them
The function will be encapsulated in a struct
*
*/
//VER : 0.2.0
package rabbitmq

import (
	"fmt"
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
)

// ConsumerConfig defines consumer parameters
type ConsumerConfig struct {
	QueueName  string
	AutoAck    bool
	WorkerFunc func(msg amqp.Delivery) error
}

// Start Consumer
func (r *RabbitMQ) StartConsumer(config ConsumerConfig) error {
	msgs, err := r.Channel.Consume(
		config.QueueName,
		"",
		config.AutoAck,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return fmt.Errorf("failed to consume messages: %v", err)
	}

	//Process messages
	go func() {
		for msg := range msgs {
			if err := config.WorkerFunc(msg); err != nil {
				log.Printf("failed to process message: %v", err)
			}
		}
	}()
	return nil
}
