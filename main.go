package main

import (
	"fmt"
	"log"
	"net/url"
	"time"
	"net/http"
	"github.com/gin-gonic/gin"
	mqtt "github.com/eclipse/paho.mqtt.golang"
)



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
	uri.Path = "ec2-18-228-38-211.sa-east-1.compute.amazonaws.com:1883"
	opts.AddBroker(fmt.Sprintf("tcp://%s", uri.Path))
	opts.SetUsername(uri.User.Username())
	password, _ := uri.User.Password()
	opts.SetPassword(password)
	opts.SetClientID(clientID)
	return opts
}

func listen(uri *url.URL, topic string, c chan string) {
	client := connect("sub", uri)


	client.Subscribe(topic, 0, func(client mqtt.Client, msg mqtt.Message) {
		fmt.Printf("* [%s] %s\n", msg.Topic(), string(msg.Payload()))

		c <- "teste"
	})	
	
}

func main() {

	uri, err := url.Parse("mqtt://admin:hivemq@http://ec2-18-228-38-211.sa-east-1.compute.amazonaws.com:1883")

	if err != nil {
		log.Fatal(err)
	}

	api(uri)	

}

func publish(uri *url.URL,topic,data string){
	client := connect("pub", uri)
	client.Publish(topic, 0, true, data)

	c := make(chan string)


	go func (t string){
		listen(uri, t, c)
	}(topic)



}


func api(uri *url.URL){

	r := gin.Default()
	
	r.GET("/temperatura/:temp", func(c *gin.Context) {
		param := c.Param("temp")
		publish(uri, "temperatura", param )
		c.String(http.StatusOK, "temperatura %s", param)
	})

	r.GET("/tomar/:algo", func(c *gin.Context) {
		param := c.Param("algo")
		publish(uri, "tomar", param )
		c.String(http.StatusOK, "tomar %s", param)
	})

	r.GET("/caminhar/", func(c *gin.Context) {
		publish(uri, "tomar", "agora" )
		c.String(http.StatusOK, "caminhar agora")
	})

	r.GET("/sentimento/:sent", func(c *gin.Context) {
		param := c.Param("sent")
		publish(uri, "tomar", param )
		c.String(http.StatusOK, "sentimento %s", param)
	})

	r.GET("/", func(c *gin.Context) {		
		c.String(http.StatusOK, "Bem vindo ao emulador de sensor")
	})


	r.Run() 

}
