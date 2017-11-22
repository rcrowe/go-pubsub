package example

import (
	"context"
	"log"

	pubsub "github.com/utilitywarehouse/go-pubsub"
	"github.com/utilitywarehouse/go-pubsub/sqs"
)

func consumerExample() {
	// create a queue first
	queue := sqs.NewConsumerQueue("accessKey", "secretKey", "eu-west-1", "https://sqs.eu-west-1.amazonaws.com/123/queueName")
	// then the consumer
	consumer := sqs.NewConsumer(queue)

	// Optionally set how long you want to wait between API poll requests (in seconds).
	// Defaults to 0 (disabled) if not set.
	consumer.WaitSeconds = 1

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// handler will handle messages from sqs. Add your message handling logic here.
	handler := func(pubsub.ConsumerMessage) error { return nil }

	// errHandler will be called if handler fails. Add your error handling logic here.
	errHandler := func(pubsub.ConsumerMessage, error) error { return nil }

	// Will poll for messages and if you need to stop the loop, call cancel()
	// to cancel the context.
	if err := consumer.ConsumeMessages(ctx, handler, errHandler); err != nil {
		log.Fatalf("failed to consume from SQS: %v", err)
	}
}
