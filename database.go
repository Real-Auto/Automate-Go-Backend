package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Get() string {
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb+srv://dev-man:Dev4life4freeonB@cluster1.viw64ds.mongodb.net/RealAutoDB?retryWrites=true&w=majority"))
	if err != nil {
		log.Fatal(err)
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(ctx)

	collection := client.Database("RealAutoDB").Collection("workers")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(collection.Name())
	return collection.Name()
}
