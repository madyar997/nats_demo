package main

import (
	"github.com/nats-io/nats.go"
	"log"
	"time"
)

func main() {
	nc, err := nats.Connect("localhost:4222", nats.Timeout(1 * time.Minute))
	if err != nil {
		log.Fatal(err)
	}
	log.Print("subscriber: established connection with nats cluster")
	defer nc.Close()

	if _, err = nc.QueueSubscribe("test_subject", "test_subject", MessagesConsumer); err != nil {
		log.Fatal("could not subscribe to queue")
	}
	log.Printf("subscribed to queue")
	
	time.Sleep(5 * time.Minute)
}

func MessagesConsumer(m *nats.Msg) {
	strMessage := string(m.Data)
	log.Printf("Async message data: %s", strMessage)
}
