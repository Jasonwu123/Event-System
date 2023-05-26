package event

import (
	"context"

	"github.com/dapr/go-sdk/client"
)

type Publisher struct {
	client client.Client
}

func NewPublisher(c client.Client) *Publisher {
	return &Publisher{client: c}
}

func (p *Publisher) Publish(ctx context.Context, pubsubName, topic string, data interface{}) error {
	err := p.client.PublishEvent(ctx, pubsubName, topic, data)
	return err
}
