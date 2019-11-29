package main

import (
	"context"
	"fmt"
	"log"

	"github.com/apache/pulsar/pulsar-client-go/pulsar"
)

// Note: relace JWT token, host, tenant, namespace, and topic
func main() {
	fmt.Println("Pulsar Consumer")

	tokenStr := "{JWT token}"
	token := pulsar.NewAuthenticationToken(tokenStr)

	// Pulsar client
	client, err := pulsar.NewClient(pulsar.ClientOptions{
		URL: "pulsar+ssl://{host}:6651",
		Authentication: token,
	})

	if err != nil {
		log.Fatal(err)
	}

	defer client.Close()

	consumer, err := client.Subscribe(pulsar.ConsumerOptions{
		Topic:            "persistent://{tenant}/{namespace}/{topic}",
		SubscriptionName: "my-subscription",
	})

	if err != nil {
		log.Fatal(err)
	}

	defer consumer.Close()

	ctx := context.Background()

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
