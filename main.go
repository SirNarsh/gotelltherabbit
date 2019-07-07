package main

import (
	"log"
	"sync"

	"github.com/sirnarsh/gotelltherabbit/http"
	"github.com/sirnarsh/gotelltherabbit/rabbitmq"
	"github.com/sirnarsh/gotelltherabbit/readconf"
)

func main() {
	log.Println("Starting...")
	readconf.CheckAllRequiredFiles()
	conf := readconf.GetGeneral()

	var wg sync.WaitGroup

	// RabbitMQ Listner thread
	if conf.EnableRabbit2HTTP {
		log.Println("Starting RabbitMQ Listener")
		wg.Add(1)
		go func() {
			rabbitmq.Receive()
		}()
	} else {
		log.Println("Rabbit 2 HTTP is not enabled in general.json")
	}

	// HTTP Listner thread
	if conf.EnableHTTP2Rabbit {
		log.Printf("Starting HTTP Listener at %s \n", conf.ServerBind)
		wg.Add(1)
		go func() {
			http.StartServer()
		}()
	} else {
		log.Println("HTTP to Rabbit is not enabled in general.json")
	}

	wg.Wait()

}
