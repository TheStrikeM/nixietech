package mongo

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

type Storage struct {
	Client *mongo.Client
}

func New(mongoURI string) (*Storage, func()) {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(mongoURI))
	if err != nil {
		log.Fatal(err)
	}
	return &Storage{
			Client: client,
		}, func() {
			if err := client.Disconnect(context.TODO()); err != nil {
				log.Fatal(err)
			}
		}
}
