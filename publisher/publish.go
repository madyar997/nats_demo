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
	log.Print("established connection with nats cluster")
	defer nc.Close()
	
	for i := 1; i <= 20000; i++ {
		msg := []byte(time.Now().String())
		if err = nc.Publish("test_subject", msg); err != nil {
			log.Fatal(err)
		}
	}
	log.Printf("finished sending to nats")

	
	time.Sleep(5 * time.Minute)
}
