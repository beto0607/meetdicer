package data

import (
	"context"
	"log"
	"meetdicer/configs"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ConnectToDB() *mongo.Client {
	mongoUri := configs.EnvMongoURI()

	client := createClient(mongoUri)

	initDB(client)

	return client
}

func createClient(mongoUri string) *mongo.Client {
	serverApiOptions := options.ServerAPI(options.ServerAPIVersion1)

	opts := options.Client().ApplyURI(mongoUri).SetServerAPIOptions(serverApiOptions)

	client, err := mongo.Connect(context.TODO(), opts)
	if err != nil {
		log.Fatal("Could not connect to the DB:", err)
	}
	ctx, cancelFunc := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancelFunc()
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Connected to MongoDB")
	return client
}

func initDB(client *mongo.Client) {
	db := client.Database("meetdicer")
	command := bson.D{{Key: "create", Value: "events"}}
	var result bson.M
	if err := db.RunCommand(context.TODO(), command).Decode(&result); err != nil {
		log.Fatal(err)
	}
	log.Println("'events' collection created")
}
