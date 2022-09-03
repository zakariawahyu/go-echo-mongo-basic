package config

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

func ConnectDB() *mongo.Client {
	env := EnvMongo()
	connPattern := "mongodb://%v:%v@%v:%v"
	if env.Username == "" {
		connPattern = "mongodb://%s%s%v:%v"
	}

	clientUrl := fmt.Sprintf(connPattern,
		env.Username,
		env.Password,
		env.Host,
		env.Port,
	)

	client, err := mongo.NewClient(options.Client().ApplyURI(clientUrl))
	if err != nil {
		log.Fatal(err)
	}

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	if err := client.Connect(ctx); err != nil {
		log.Fatal(err)
	}

	//Ping Database
	if err := client.Ping(ctx, nil); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connect to MongoDB")
	return client
}

// Client Instance
var DB *mongo.Client = ConnectDB()

// Get Database Collection
func GetCollection(client *mongo.Client, collectionName string) *mongo.Collection {
	env := EnvMongo()
	return client.Database(env.DBName).Collection(collectionName)
}
