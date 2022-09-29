package main

import (
	"fmt"
	"log"
	"runtime"
	"strings"

	"github.com/nats-io/nats.go"
)

func main() {
	nc, err := nats.Connect("demo.nats.io:4222")
	if err != nil {
		log.Fatalln(err)
	}

	nc.QueueSubscribe("greet.*", "greet", func(msg *nats.Msg) {
		log.Println("Received message on " + msg.Subject)
		msg.Respond([]byte(fmt.Sprintf("Hello, " + strings.Split(msg.Subject, ".")[1])))
	})

	log.Println("Listening on greet.*...")
	runtime.Goexit()
}
