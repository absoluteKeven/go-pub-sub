package pubsub

import (
	"context"
	"encoding/json"
	"fmt"

	amqp "github.com/rabbitmq/amqp091-go"
)

func PublishJSON[T any](ch *amqp.Channel, exchange, key string, val T) error {
	data, err := json.Marshal(val)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	e := ch.PublishWithContext(context.Background(), exchange, key, false, false, amqp.Publishing{ContentType: "aplication/json", Body: data})

	return e
}
