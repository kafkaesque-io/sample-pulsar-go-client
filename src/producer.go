package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/apache/pulsar-client-go/pulsar"
)

// Note: relace JWT token, tenant, namespace, and topic
func main() {
	log.Println("Pulsar Producer")

	// Configuration variables pertaining to this producer
	tokenStr := "{JWT token}"
	uri := "pulsar+ssl://{host}:6651"
	trustStore := "/etc/ssl/certs/ca-bundle.crt"
	topicName := "persistent://{tenant}/{namespace}/{topic}"

	token := pulsar.NewAuthenticationToken(tokenStr)

	client, err := pulsar.NewClient(pulsar.ClientOptions{
		URL:                   uri,
		Authentication:        token,
		TLSTrustCertsFilePath: trustStore,
		OperationTimeout:      30 * time.Second,
		ConnectionTimeout:     30 * time.Second,
	})

	if err != nil {
		log.Fatalf("Could not instantiate Pulsar client: %v", err)
	}

	defer client.Close()

	log.Printf("creating producer...")

	// Use the client to instantiate a producer
	producer, err := client.CreateProducer(pulsar.ProducerOptions{
		Topic: topicName,
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
		messageID, err := producer.Send(ctx, &msg)
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("the %s successfully published with the message ID %v\n", string(msg.Payload), messageID)

		// Create a different message to send asynchronously
		asyncMsg := pulsar.ProducerMessage{
			Payload: []byte(fmt.Sprintf("asyncMessageId-%d", i)),
		}

		// Attempt to send the message asynchronously and handle the response
		producer.SendAsync(ctx, &asyncMsg, func(msgID pulsar.MessageID, msg *pulsar.ProducerMessage, err error) {
			if err != nil {
				log.Fatal(err)
			}

			log.Printf("the %s successfully published with the message ID %v\n", string(msg.Payload), msgID)
		})
	}
}
