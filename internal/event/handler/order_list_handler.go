package handler

import (
	"encoding/json"
	"fmt"
	"github.com/LeoAntunesBrombilla/clean_arch/pkg/events"
	"sync"

	"github.com/streadway/amqp"
)

type OrderListHandler struct {
	RabbitMQChannel *amqp.Channel
}

func NewOrderListHandler(rabbitMQChannel *amqp.Channel) *OrderListHandler {
	return &OrderListHandler{
		RabbitMQChannel: rabbitMQChannel,
	}
}

func (h *OrderListHandler) Handle(event events.EventInterface, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("List Orders: %v", event.GetPayload())
	jsonOutput, _ := json.Marshal(event.GetPayload())

	msgRabbitmq := amqp.Publishing{
		ContentType: "application/json",
		Body:        jsonOutput,
	}

	h.RabbitMQChannel.Publish(
		"amq.direct", // exchange
		"",           // key name
		false,        // mandatory
		false,        // immediate
		msgRabbitmq,  // message to publish
	)
}
