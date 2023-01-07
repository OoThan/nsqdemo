package main

import (
	"flag"
	"github.com/nsqio/go-nsq"
	"log"
	"os"
)

var (
	addr    = flag.String("addr", "localhost:4150", "NSQ addr")
	topic   = flag.String("topic", "", "NSQ topic")
	message = flag.String("message", "", "Message body")
)

func main() {
	flag.Parse()
	if *topic == "" || *message == "" {
		flag.PrintDefaults()
		os.Exit(1)
	}

	config := nsq.NewConfig()
	producer, err := nsq.NewProducer(*addr, config)
	if err != nil {
		log.Fatal(err)
	}

	err = producer.Publish(*topic, []byte(*message))
	if err != nil {
		log.Fatal(err)
	}

	producer.Stop()
}
