package kafka

import (
	"context"
	"github.com/Erlendum/go_CQRS_microservices/pkg/logger"
	"github.com/segmentio/kafka-go"
	"github.com/segmentio/kafka-go/compress"
	"sync"
)

type MessageProcessor interface {
	ProcessMessages(ctx context.Context, r *kafka.Reader, wg *sync.WaitGroup, workerID int)
}

type Worker func(ctx context.Context, r *kafka.Reader, wg *sync.WaitGroup, workerID int)

type ConsumerGroup interface {
	ConsumeTopic(ctx context.Context, groupTopics []string, poolSize int, worker Worker)
	GetNewKafkaReader(kafkaURL []string, groupTopics []string, groupID string) *kafka.Reader
	GetNewKafkaWriter() *kafka.Writer
}

type consumerGroup struct {
	Brokers []string
	GroupID string
	log     logger.Logger
}

func NewConsumerGroup(brokers []string, groupID string, log logger.Logger) ConsumerGroup {
	return &consumerGroup{Brokers: brokers, GroupID: groupID, log: log}
}

func (c *consumerGroup) GetNewKafkaReader(kafkaURL []string, groupTopics []string, groupID string) *kafka.Reader {
	return kafka.NewReader(kafka.ReaderConfig{
		Brokers:                kafkaURL,
		GroupID:                groupID,
		GroupTopics:            groupTopics,
		MinBytes:               minBytes,
		MaxBytes:               maxBytes,
		QueueCapacity:          queueCapacity,
		HeartbeatInterval:      heartbeatInterval,
		CommitInterval:         commitInterval,
		PartitionWatchInterval: partitionWatchInterval,
		MaxAttempts:            maxAttempts,
		MaxWait:                maxWait,
		Dialer: &kafka.Dialer{
			Timeout: dialTimeout,
		},
	})
}

func (c *consumerGroup) GetNewKafkaWriter() *kafka.Writer {
	w := &kafka.Writer{
		Addr:         kafka.TCP(c.Brokers...),
		Balancer:     &kafka.LeastBytes{},
		RequiredAcks: writerRequiredAcks,
		MaxAttempts:  writerMaxAttempts,
		Compression:  compress.Snappy,
		ReadTimeout:  writerReadTimeout,
		WriteTimeout: writerWriteTimeout,
	}

	return w
}

func (c *consumerGroup) ConsumeTopic(ctx context.Context, groupTopics []string, poolSize int, worker Worker) {
	r := c.GetNewKafkaReader(c.Brokers, groupTopics, c.GroupID)

	defer func() {
		if err := r.Close(); err != nil {
			c.log.Warnf("consumerGroup.r.Close: %v", err)
		}
	}()

	c.log.Infof("Starting consumer groupID: %s, topic: %+v, pool size: %v", c.GroupID, groupTopics, poolSize)

	wg := &sync.WaitGroup{}
	for i := 0; i <= poolSize; i++ {
		wg.Add(1)
		go worker(ctx, r, wg, i)
	}
	wg.Wait()
}
