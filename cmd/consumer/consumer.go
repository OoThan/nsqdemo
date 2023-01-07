package main

import (
	"flag"
	"github.com/nsqio/go-nsq"
	"log"
	"os"
	"os/signal"
	"syscall"
)

var (
	addr    = flag.String("addr", "nsqlookupd:4161", "NSQ lookup addr")
	topic   = flag.String("topic", "", "NSQ topic")
	channel = flag.String("channel", "", "NSQ channel")
)

type MyHandler struct{}

func (h *MyHandler) HandleMessage(message *nsq.Message) error {
	log.Printf("Got a message: %s", string(message.Body))
	return nil
}

func main() {
	flag.Parse()
	if *topic == "" || *channel == "" {
		flag.PrintDefaults()
		os.Exit(1)
	}

	config := nsq.NewConfig()
	consumer, err := nsq.NewConsumer(*topic, *channel, config)
	if err != nil {
		log.Fatal(err)
	}
	consumer.AddHandler(&MyHandler{})

	err = consumer.ConnectToNSQLookupd(*addr)
	if err != nil {
		log.Fatal(err)
	}

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	<-sigChan

	consumer.Stop()
}
