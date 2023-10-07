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

type UpdateProductCmdHandler interface {
	Handle(ctx context.Context, command *UpdateProductCommand) error
}

type updateProductCmdHandler struct {
	log           logger.Logger
	cfg           *config.Config
	kafkaProducer kafkaClient.Producer
}

func NewUpdateProductHandler(log logger.Logger, cfg *config.Config, kafkaProducer kafkaClient.Producer) *updateProductCmdHandler {
	return &updateProductCmdHandler{log: log, cfg: cfg, kafkaProducer: kafkaProducer}
}

func (c *updateProductCmdHandler) Handle(ctx context.Context, command *UpdateProductCommand) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "updateProductCmdHandler.Handle")
	defer span.Finish()

	updateDto := &kafkaMessages.ProductUpdate{
		ProductID:   command.UpdateDto.ProductID.String(),
		Name:        command.UpdateDto.Name,
		Description: command.UpdateDto.Description,
		Price:       command.UpdateDto.Price,
	}

	dtoBytes, err := proto.Marshal(updateDto)
	if err != nil {
		return err
	}

	return c.kafkaProducer.PublishMessage(ctx, kafka.Message{
		Topic: c.cfg.KafkaTopics.ProductUpdate.TopicName,
		Value: dtoBytes,
		Time:  time.Now().UTC(),
		//Headers: tracing.GetKafkaTracingHeadersFromSpanCtx(span.Context()),
	})
}
