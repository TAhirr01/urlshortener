package config

import (
	"context"
	"fmt"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// NOTE: Bura util packagende olan Retry funcsiyasini implement elemek lazimdir.
// Retry heleki yaxsi deyil belkede oxunaqliqi pozur lazimsiz cetinlik yaradir
func ConnectDB(uri string) (*mongo.Client, error) {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Ping to verify connection
	err = client.Ping(ctx, nil)
	if err != nil {
		return nil, err
	}

	fmt.Println("Connected to MongoDB!")
	return client, nil
}

// GetCollection helper
func GetCollection(client *mongo.Client, collectionName string) *mongo.Collection {
	dbName := os.Getenv("DB_NAME")
	if dbName == "" {
		dbName = "url_shortener"
	}
	return client.Database(dbName).Collection(collectionName)
}
