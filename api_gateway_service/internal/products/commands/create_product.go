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

type CreateProductCmdHandler interface {
	Handle(ctx context.Context, command *CreateProductCommand) error
}

type createProductHandler struct {
	log           logger.Logger
	cfg           *config.Config
	kafkaProducer kafkaClient.Producer
}

func NewCreateProductHandler(log logger.Logger, cfg *config.Config, kafkaProducer kafkaClient.Producer) CreateProductCmdHandler {
	return &createProductHandler{log: log, cfg: cfg, kafkaProducer: kafkaProducer}
}

func (c *createProductHandler) Handle(ctx context.Context, command *CreateProductCommand) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "createProductHandler.Handle")
	defer span.Finish()

	createDto := &kafkaMessages.ProductCreate{
		ProductID:   command.CreateDto.ProductID.String(),
		Name:        command.CreateDto.Name,
		Description: command.CreateDto.Description,
		Price:       command.CreateDto.Price,
	}

	dtoBytes, err := proto.Marshal(createDto)
	if err != nil {
		return err
	}

	return c.kafkaProducer.PublishMessage(ctx, kafka.Message{
		Topic: c.cfg.KafkaTopics.ProductCreate.TopicName,
		Value: dtoBytes,
		Time:  time.Now().UTC(),
		//Headers: tracing.GetKafkaTracingHeadersFromSpanCtx(span.Context()),
	})
}
