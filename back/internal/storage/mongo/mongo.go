package mongo

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	consts "nixietech/internal"
)

type Storage struct {
	Client *mongo.Client
}

func New(mongoURI string) (*Storage, func()) {
	// Здесь будет логика работы с таймаутами
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

func (s *Storage) GetCollection(name string) *mongo.Collection {
	return s.Client.Database(consts.DatabaseName).Collection(name)
}
