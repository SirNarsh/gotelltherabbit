package main

import (
	"fmt"

	"github.com/sirnarsh/gotelltherabbit/rabbitmq"
)

func main() {
	fmt.Print("Testing")
	rabbitmq.Send()
	rabbitmq.Receive()
}
