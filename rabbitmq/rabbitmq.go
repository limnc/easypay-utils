/*
*

RabbitMQ is a message broker that implements the Advanced Message Queuing Protocol (AMQP)
This package provides a connection to RabbitMQ and a channel to send and receive messages
This will establish a connection to RabbitMQ and keep it alive

*
*/
package rabbitmq

import (
	"fmt"

	amqp "github.com/rabbitmq/amqp091-go"
)

//RabbitMQ struct

type RabbitMQ struct {
	Conn    *amqp.Connection
	Channel *amqp.Channel
}

// New RabbitMQ Connection initializes a connection that stay alive
func NewRabbitMQConnection(amqpURL string) (*RabbitMQ, error) {
	conn, err := amqp.Dial(amqpURL)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to RabbitMQ: %v", err)
	}

	ch, err := conn.Channel()
	if err != nil {
		conn.Close()
		return nil, fmt.Errorf("faile to open a channel :%v", err)
	}

	return &RabbitMQ{Conn: conn, Channel: ch}, nil
}

// Close closes the RabbitMQ connection and channel
func (r *RabbitMQ) Close() {
	if r.Channel != nil {
		r.Channel.Close()
	}

	if r.Conn != nil {
		r.Conn.Close()
	}
}

func (r *RabbitMQ) DeclareQueue(queueName string) error {
	channel, err := r.Conn.Channel()
	if err != nil {
		return err
	}
	defer channel.Close()

	_, err = channel.QueueDeclare(
		queueName, // Queue name
		true,      // Durable
		false,     // Auto-delete
		false,     // Exclusive
		false,     // No-wait
		nil,       // Arguments
	)
	return err
}
