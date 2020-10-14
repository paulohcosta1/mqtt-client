package main

import (
	"context"
	"log"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/db"
)

var client *db.Client

func init() {
	ctx := context.Background()
	conf := &firebase.Config{
		DatabaseURL: "https://mqtt-client-af8e1.firebaseio.com/",
	}
	app, err := firebase.NewApp(ctx, conf)
	if err != nil {
		log.Fatalf("firebase.NewApp: %v", err)
	}
	client, err = app.Database(ctx)
	if err != nil {
		log.Fatalf("app.Firestore: %v", err)
	}
}

func setTopic(topic, msg string) {
	client.NewRef(topic).Child(msg)

}
