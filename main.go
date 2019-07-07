package main

import (
	"log"

	"github.com/sirnarsh/gotelltherabbit/rabbitmq"
	"github.com/sirnarsh/gotelltherabbit/readconf"
)

func main() {
	log.Println("Starting...")
	readconf.CheckAllRequiredFiles()
	conf := readconf.GetGeneral()

	if conf.EnableRabbit2HTTP {
		log.Println("Starting RabbitMQ Listener")
		rabbitmq.Receive()
	} else {
		log.Println("Rabbit 2 HTTP is not enabled in general.json")
	}

	if conf.EnableHTTP2Rabbit {
		log.Println("Starting HTTP Listener")
		log.Println("Warning, HTTP listener is not implemented yet")
		// @todo add server listening to HTTP to and use rabbitmq.Send()
	} else {
		log.Println("HTTP to Rabbit is not enabled in general.json")
	}
}
