package main

import (
	"fmt"

	"github.com/absoluteKeven/go-pub-sub/internal/gamelogic"
	amqp "github.com/rabbitmq/amqp091-go"
)

func main() {
	fmt.Println("Starting Peril client...")

	connS := "amqp://guest:guest@localhost:5672/"
	conn, err := amqp.Dial(connS)
	if err != nil {
		fmt.Printf("%s\n", err)
		return
	}
	defer conn.Close()

	gamelogic.ClientWelcome()
}
