package main

import (
	"context"
	"fmt"

	"log"

	"github.com/apache/pulsar/pulsar-client-go/pulsar"
)

// Note: relace JWT token, tenant, namespace, and topic
func main() {
	fmt.Println("Pulsar Consumer")

	tokenStr := "{JWT token}"
	token := pulsar.NewAuthenticationToken(tokenStr)

	client, err := pulsar.NewClient(pulsar.ClientOptions{
		URL:                     "pulsar+ssl://{host}:6651",
		Authentication:          token,
		IOThreads:               3,
		OperationTimeoutSeconds: 5,
	})

	if err != nil {
		log.Fatalf("Could not instantiate Pulsar client: %v", err)
	}

	defer client.Close()

	log.Printf("creating producer...")

	// Use the client to instantiate a producer
	producer, err := client.CreateProducer(pulsar.ProducerOptions{
		Topic: "persistent://{tenant}/{namespace}/{topic}",
	})

	log.Printf("checking error of producer creation...")
	if producer == nil {
		log.Print("producer is null")
	}
	if err != nil {
		log.Fatal(err)
	}

	defer producer.Close()

	ctx := context.Background()

	// Send 3 messages synchronously and 3 messages asynchronously
	for i := 0; i < 3; i++ {
		// Create a message
		msg := pulsar.ProducerMessage{
			Payload: []byte(fmt.Sprintf("messageId-%d", i)),
		}

		// Attempt to send the message
		if err := producer.Send(ctx, msg); err != nil {
			log.Fatal(err)
		}

		// Create a different message to send asynchronously
		asyncMsg := pulsar.ProducerMessage{
			Payload: []byte(fmt.Sprintf("asyncMessageId-%d", i)),
		}

		// Attempt to send the message asynchronously and handle the response
		producer.SendAsync(ctx, asyncMsg, func(msg pulsar.ProducerMessage, err error) {
			if err != nil {
				log.Fatal(err)
			}

			fmt.Printf("the %s successfully published\n", string(msg.Payload))
		})
	}
}
