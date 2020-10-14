package main

import (
	"fmt"
	"log"
	"os"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

func publish(topic, data string) {
	go listen(topic)

	client := connect("pub")
	client.Publish(topic, 0, true, data)

}

func listen(topic string) {
	// client := connect("sub")
	setTopic("teste", "teste")

	// client.Subscribe(topic, 0, func(client mqtt.Client, msg mqtt.Message) {

	// })

}

func connect(clientID string) mqtt.Client {
	opts := createClientOptions(clientID)
	client := mqtt.NewClient(opts)
	token := client.Connect()
	for !token.WaitTimeout(3 * time.Second) {
	}
	if err := token.Error(); err != nil {
		log.Fatal(err)
	}

	return client
}

func createClientOptions(clientID string) *mqtt.ClientOptions {
	opts := mqtt.NewClientOptions()
	opts.AddBroker(fmt.Sprintf("tcp://%s", os.Getenv("AWS_PATH")))
	opts.SetUsername(os.Getenv("AWS_USERNAME"))
	opts.SetPassword(os.Getenv("AWS_PASSWORD"))
	opts.SetClientID(clientID)
	return opts
}
