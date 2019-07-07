package rabbitmq

import (
	"fmt"
	"log"

	"github.com/sirnarsh/gotelltherabbit/readconf"
	"github.com/streadway/amqp"
)

/*
Send a message to exchange with body
If exchange is not declared yet this will also declare the exchange
*/
func Send(exchangeName string, body []byte) {

	conf := readconf.GetGeneral()

	conn, err := amqp.Dial(conf.RabbitMQServer)
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	// Declare Exchange to send message to, ideally exchange should be already declared and binded to queues
	{
		err := ch.ExchangeDeclare(exchangeName, "fanout", true, false, false, false, nil)
		failOnError(err, fmt.Sprintf("Failed to declare exchange '%s'", exchangeName))
		log.Printf("Declared or found exchange %s", exchangeName)
	}

	err = ch.Publish(
		exchangeName, // exchange
		"",           // routing key
		false,        // mandatory
		false,        // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        body,
		})
	log.Printf(" [x] Sent msg to exchange %s", exchangeName)
	failOnError(err, "Failed to publish a message")
}
