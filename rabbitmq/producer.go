/*
*
This package provides a producer for RabbitMQ messages.
It will publish messages to a queue.
The function will be encapsulated in a struct.
*
*/

package rabbitmq

import (
	"encoding/json"
	"fmt"

	amqp "github.com/rabbitmq/amqp091-go"
)

// PublishMessage defines the structure of a message to be published
type PublishMessage struct {
	Exchange   string
	RoutingKey string
	Body       interface{} // This can be any type of data
}

// Producer defines the interface for message publishing
type Producer interface {
	Publish(msg PublishMessage) error
}

// RabbitMQ implements the Producer interface
func (r *RabbitMQ) Publish(msg PublishMessage) error {
	body, err := json.Marshal(msg.Body)
	if err != nil {
		return fmt.Errorf("failed to marshal message body: %v", err)
	}

	err = r.Channel.Publish(
		msg.Exchange,   // Exchange name
		msg.RoutingKey, // Routing Key
		false,          // Mandatory
		false,          // Immediate
		amqp.Publishing{
			ContentType: "application/json",
			Body:        body,
		},
	)

	if err != nil {
		return fmt.Errorf("failed to publish message: %v", err)
	}

	return nil
}
