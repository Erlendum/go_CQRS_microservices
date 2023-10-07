package kafka

import (
	"context"
	"github.com/Erlendum/go_CQRS_microservices/pkg/logger"
	"github.com/segmentio/kafka-go"
)

type Producer interface {
	PublishMessage(ctx context.Context, msgs ...kafka.Message) error
	Close() error
}

type producer struct {
	log     logger.Logger
	brokers []string
	w       *kafka.Writer
}

func NewProducer(log logger.Logger, brokers []string) Producer {
	return &producer{log: log, brokers: brokers, w: NewWriter(brokers, kafka.LoggerFunc(log.Errorf))}
}

func (p *producer) PublishMessage(ctx context.Context, msgs ...kafka.Message) error {
	return p.w.WriteMessages(ctx, msgs...)
}

func (p *producer) Close() error {
	return p.w.Close()
}
