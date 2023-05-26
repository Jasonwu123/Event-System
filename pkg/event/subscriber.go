package event

import (
	"context"

	"github.com/dapr/go-sdk/service/common"
	"github.com/dapr/go-sdk/service/http"
)

type Subscriber struct {
	srv common.Service
	pubsubName string
}

func NewSubscriber(addr string, pubsubName string) *Subscriber {
	srv := http.NewService(addr)
	return &Subscriber{srv: srv, pubsubName: pubsubName}
}

func (s *Subscriber) Subscribe(topic string, handler func(ctx context.Context, e *common.TopicEvent) (retry bool, err error)) error  {
	sub := &common.Subscription{
		PubsubName: s.pubsubName,
		Topic: topic,
	}
	return s.srv.AddTopicEventHandler(sub, handler)
}