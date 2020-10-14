package main

import (
	"github.com/gin-gonic/gin"
	"net/http"	
)

// api simulates an amount of sensors to feed our broker
func api() {

	r := gin.Default()

	topics := map[string]string{
		"temp" : "temperatura", 
		"to" : "tomar", 
		"ca" : "temcaminharperatura", 
		"sent" : "sentimento", 
	}
	
	r.GET("/"+ topics["temp"] + "/:temp", func(c *gin.Context) {
		param := c.Param("temp")
		
		publish( topics["temp"], param)
		c.String(http.StatusOK, "temperatura %s", param)
	})

	r.GET("/"+ topics["to"] + "/:algo", func(c *gin.Context) {
		param := c.Param("algo")
		publish( topics["to"], param)
		c.String(http.StatusOK, "tomar %s", param)
	})

	r.GET("/"+ topics["ca"] + "/", func(c *gin.Context) {
		publish( topics["ca"], "agora")
		c.String(http.StatusOK, "caminhar agora")
	})

	r.GET("/"+ topics["sent"] + "/:sent", func(c *gin.Context) {
		param := c.Param("sent")
		publish(topics["sent"], param)
		c.String(http.StatusOK, "sentimento %s", param)
	})

	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Bem vindo ao emulador de sensor")
	})

	r.Run()

}

