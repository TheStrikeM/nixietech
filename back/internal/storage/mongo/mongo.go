package mongo

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	consts "nixietech/internal"
	"strings"
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

func ObjectId(item interface{}) primitive.ObjectID {
	id, err := primitive.ObjectIDFromHex(fmt.Sprintf("%v", item))
	if err != nil {
		log.Fatal(err)
	}
	return id
}

func InsertedIdToString(item interface{}) string {
	return strings.Split(fmt.Sprintf("%v", item), "\"")[1]
}

func (s *Storage) GetCollection(name string) *mongo.Collection {
	return s.Client.Database(consts.DatabaseName).Collection(name)
}
