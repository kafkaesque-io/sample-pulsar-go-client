package main

import (
	"context"
	"fmt"
	"log"

	"github.com/apache/pulsar-client-go/pulsar"
)

// Note: relace JWT token, host, tenant, namespace, and topic
func main() {
	fmt.Println("Pulsar Consumer")

	// Configuration variables pertaining to this consumer
	tokenStr := "{JWT token}"
	uri := "pulsar+ssl://{host}:6651"
	trustStore := "/etc/ssl/certs/ca-bundle.crt"
	topicName := "persistent://{tenant}/{namespace}/{topic}"
	subscriptionName := "my-subscription"

	token := pulsar.NewAuthenticationToken(tokenStr)

	// Pulsar client
	client, err := pulsar.NewClient(pulsar.ClientOptions{
		URL:                   uri,
		Authentication:        token,
		TLSTrustCertsFilePath: trustStore,
	})

	if err != nil {
		log.Fatal(err)
	}

	defer client.Close()

	consumer, err := client.Subscribe(pulsar.ConsumerOptions{
		Topic:            topicName,
		SubscriptionName: subscriptionName,
	})

	if err != nil {
		log.Fatal(err)
	}

	defer consumer.Close()

	ctx := context.Background()

	// infinite loop to receive messages
	for {
		msg, err := consumer.Receive(ctx)
		if err != nil {
			log.Fatal(err)
		} else {
			fmt.Printf("Received message : %v\n", string(msg.Payload()))
		}

		consumer.Ack(msg)
	}

}
