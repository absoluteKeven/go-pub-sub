package main

import (
	"fmt"
	"os"
	"os/signal"

	pubsub "github.com/absoluteKeven/go-pub-sub/internal/pubsub"
	routing "github.com/absoluteKeven/go-pub-sub/internal/routing"

	amqp "github.com/rabbitmq/amqp091-go"
)

func main() {
	connS := "amqp://guest:guest@localhost:5672/"

	conn, err := amqp.Dial(connS)
	if err != nil {
		fmt.Printf("%s\n", err)
		return
	}
	defer conn.Close()
	ch, err := conn.Channel()
	if err != nil {
		fmt.Printf("%s\n", err)
		return
	}

	pubsub.PublishJSON(ch, routing.ExchangePerilDirect, routing.PauseKey, routing.PlayingState{IsPaused: true})

	fmt.Println("Connection Successful...")

	c := make(chan os.Signal, 1)

	signal.Notify(c, os.Interrupt)
	<-c
	fmt.Println("Shutting gracefully.")
}
