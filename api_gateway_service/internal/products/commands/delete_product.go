package commands

import (
	"context"
	"github.com/Erlendum/go_CQRS_microservices/api_gateway_service/config"
	kafkaClient "github.com/Erlendum/go_CQRS_microservices/pkg/kafka"
	"github.com/Erlendum/go_CQRS_microservices/pkg/logger"
	kafkaMessages "github.com/Erlendum/go_CQRS_microservices/proto/kafka"
	"github.com/opentracing/opentracing-go"
	"github.com/segmentio/kafka-go"
	"google.golang.org/protobuf/proto"
	"time"
)

type DeleteProductCmdHandler interface {
	Handle(ctx context.Context, command *DeleteProductCommand) error
}

type deleteProductHandler struct {
	log           logger.Logger
	cfg           *config.Config
	kafkaProducer kafkaClient.Producer
}

func NewDeleteProductHandler(log logger.Logger, cfg *config.Config, kafkaProducer kafkaClient.Producer) *deleteProductHandler {
	return &deleteProductHandler{log: log, cfg: cfg, kafkaProducer: kafkaProducer}
}

func (c *deleteProductHandler) Handle(ctx context.Context, command *DeleteProductCommand) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "deleteProductHandler.Handle")
	defer span.Finish()

	createDto := &kafkaMessages.ProductDelete{ProductID: command.ProductID.String()}

	dtoBytes, err := proto.Marshal(createDto)
	if err != nil {
		return err
	}

	return c.kafkaProducer.PublishMessage(ctx, kafka.Message{
		Topic: c.cfg.KafkaTopics.ProductDelete.TopicName,
		Value: dtoBytes,
		Time:  time.Now().UTC(),
		//Headers: tracing.GetKafkaTracingHeadersFromSpanCtx(span.Context()),
	})
}
