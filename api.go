package main

import (
	"fmt"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"net/url"
	"os"
	"time"
)

// api simulates an amount of sensors to feed our broker
func api(uri *url.URL) {

	r := gin.Default()

	r.GET("/temperatura/:temp", func(c *gin.Context) {
		param := c.Param("temp")
		publish(uri, "temperatura", param)
		c.String(http.StatusOK, "temperatura %s", param)
	})

	r.GET("/tomar/:algo", func(c *gin.Context) {
		param := c.Param("algo")
		publish(uri, "tomar", param)
		c.String(http.StatusOK, "tomar %s", param)
	})

	r.GET("/caminhar/", func(c *gin.Context) {
		publish(uri, "tomar", "agora")
		c.String(http.StatusOK, "caminhar agora")
	})

	r.GET("/sentimento/:sent", func(c *gin.Context) {
		param := c.Param("sent")
		publish(uri, "tomar", param)
		c.String(http.StatusOK, "sentimento %s", param)
	})

	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Bem vindo ao emulador de sensor")
	})

	r.Run()

}

func publish(uri *url.URL, topic, data string) {
	client := connect("pub", uri)
	client.Publish(topic, 0, true, data)

	c := make(chan string)

	go func(t string) {
		listen(uri, t, c)
	}(topic)
}

func listen(uri *url.URL, topic string, c chan string) {
	client := connect("sub", uri)

	client.Subscribe(topic, 0, func(client mqtt.Client, msg mqtt.Message) {
		fmt.Printf("* [%s] %s\n", msg.Topic(), string(msg.Payload()))

		c <- "teste"
	})

}

func connect(clientID string, uri *url.URL) mqtt.Client {
	opts := createClientOptions(clientID, uri)
	client := mqtt.NewClient(opts)
	token := client.Connect()
	for !token.WaitTimeout(3 * time.Second) {
	}
	if err := token.Error(); err != nil {
		log.Fatal(err)
	}

	return client
}

func createClientOptions(clientID string, uri *url.URL) *mqtt.ClientOptions {
	opts := mqtt.NewClientOptions()
	opts.AddBroker(fmt.Sprintf("tcp://%s", os.Getenv("AWS_PATH")))
	opts.SetUsername(uri.User.Username())
	password, _ := uri.User.Password()
	opts.SetPassword(password)
	opts.SetClientID(clientID)
	return opts
}
