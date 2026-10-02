package main

import (
	"fmt"

	gamelogic "github.com/absoluteKeven/go-pub-sub/internal/gamelogic"
	"github.com/absoluteKeven/go-pub-sub/internal/pubsub"
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

	name, err := gamelogic.ClientWelcome()
	
	pubsub.DeclareAndBind(conn, "peril_direct", pause.username)
}
