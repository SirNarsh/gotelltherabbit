package rabbitmq

import (
	"fmt"
	"log"

	"github.com/sirnarsh/gotelltherabbit/readconf"
	"github.com/streadway/amqp"
)

// Receive messages from rabbit mq and send to http
func Receive() {

	var conf = readconf.GetGeneral()
	var r2hConf = readconf.GetR2H()

	conn, err := amqp.Dial(conf.RabbitMQServer)
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	// Declare our listener queue
	{
		_, err := ch.QueueDeclare(
			r2hConf.ListenQueue, // name
			true,                // durable
			false,               // delete when unused
			false,               // exclusive
			false,               // no-wait
			nil,                 // arguments
		)
		failOnError(err, "Failed to declare a queue")
	}

	// Bind the queue we are consuming with all exchanges we want to listen to.
	for _, exchange := range r2hConf.BindExchanges {
		{
			err := ch.ExchangeDeclare(exchange.ExchangeName, "fanout", true, false, false, false, nil)
			failOnError(err, fmt.Sprintf("Failed to bind Exchange '%s' to Queue '%s'", r2hConf.ListenQueue, exchange.ExchangeName))
			log.Printf("Declared or found exchange %s", exchange.ExchangeName)
		}
		{
			err := ch.QueueBind(
				r2hConf.ListenQueue,
				"",
				exchange.ExchangeName,
				false,
				nil,
			)
			failOnError(err, fmt.Sprintf("Failed to bind Exchange '%s' to Queue '%s'", r2hConf.ListenQueue, exchange.ExchangeName))
			log.Printf("Bind Exchange '%s' to listener queue '%s'", exchange.ExchangeName, r2hConf.ListenQueue)
		}
	}

	msgs, err := ch.Consume(
		r2hConf.ListenQueue, // queue
		"",                  // consumer
		false,               // auto-ack
		false,               // exclusive
		false,               // no-local
		false,               // no-wait
		nil,                 // args
	)
	failOnError(err, "Failed to register a consumer")

	listener := make(chan bool)

	go func() {
		for d := range msgs {
			log.Printf("Received a message: %s, from exchange %s with routing %s", d.Body, d.Exchange, d.RoutingKey)
			// @todo ACK Should be send after we send the HTTP request and get code 200
			// For now we ACK message after consuming directly
			d.Ack(false)
		}
	}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-listener
}
