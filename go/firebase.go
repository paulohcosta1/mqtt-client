package main

import (
	"context"
	"log"
	"time"
	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go"
	"google.golang.org/api/option"
)

var client *firestore.Client

func init() {

	opt := option.WithCredentialsFile("serviceAccountKey.json")
	ctx := context.Background()

	app, err := firebase.NewApp(ctx, nil, opt)

	if err != nil {
		log.Fatalf("firebase.NewApp: %v", err)
	}
	client, err = app.Firestore(ctx)

	if err != nil {
		log.Fatalf("app.Firestore: %v", err)
	}

}

func addTopic(topic, payload string) {
	ctx := context.Background()

	_, _, err := client.Collection("topics").Add(ctx, map[string]interface{}{
		"payload": payload,
		"description": topic,
		"insertedAt": time.Now(),
	})

	if err != nil {
		log.Fatalf("Failed adding topic: %v", err)
	}

}
