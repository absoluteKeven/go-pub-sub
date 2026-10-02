package pubsub

import (
	"context"
	"encoding/json"
	"fmt"

	amqp "github.com/rabbitmq/amqp091-go"
)

type SimpleQueueType int

const (
	durable SimpleQueueType = iota
	transient
)

func PublishJSON[T any](ch *amqp.Channel, exchange, key string, val T) error {
	json, err := json.Marshal(val)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	e := ch.PublishWithContext(context.Background(), exchange, key, false, false, amqp.Publishing{ContentType: "aplication/json", Body: json})

	return e
}

func DeclareAndBind(
	conn *amqp.Connection,
	exchange,
	queueName,
	key string,
	queueType SimpleQueueType,
) (*amqp.Channel, amqp.Queue, error) {
	ch, errc := conn.Channel()
	if errc != nil {
		return nil, amqp.Queue{}, errc
	}

	queue, errq := ch.QueueDeclare(queueName, queueType == durable, queueType == transient, queueType == transient, false, nil)
	if errq != nil {
		return nil, amqp.Queue{}, errc
	}

	err := ch.QueueBind(queueName, key, exchange, false, nil)
	if err != nil {
		return nil, amqp.Queue{}, errc
	}

	return ch, queue, err
}
