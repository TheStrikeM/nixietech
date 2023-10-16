package mongo

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	consts "nixietech/internal"
	"nixietech/internal/storage"
	_ "nixietech/internal/storage/mongo"
	mongoStorage "nixietech/internal/storage/mongo"
)

type Clock struct {
	collection *mongo.Collection
	storage    *mongoStorage.Storage
}

func New(storage *mongoStorage.Storage) *Clock {
	return &Clock{
		storage:    storage,
		collection: storage.GetCollection(consts.CollectionClockName),
	}
}

func (manager *Clock) AddClock(clock storage.ClockWithoutId) (*storage.Clock[primitive.ObjectID], error) {

	clockId, err := manager.collection.InsertOne(context.TODO(), clock)
	if err != nil {
		return nil, err
	}

	return &storage.Clock[primitive.ObjectID]{
		Id:        mongoStorage.ObjectId(clockId),
		Name:      clock.Name,
		Avatar:    clock.Avatar,
		Photos:    clock.Photos,
		Price:     clock.Price,
		Existence: clock.Existence,
		Type:      clock.Type,
	}, nil
}

func (manager *Clock) RemoveClock(id primitive.ObjectID) (*primitive.ObjectID, error) {
	if _, err := manager.collection.DeleteOne(context.TODO(), bson.D{{"_id", id}}); err != nil {
		return nil, err
	}
	return &id, nil
}

func (manager *Clock) UpdateClock(
	id primitive.ObjectID,
	updatedClock storage.ClockWithoutId,
) (*storage.Clock[primitive.ObjectID], error) {
	clockId, err := manager.collection.UpdateOne(context.TODO(), bson.D{{"_id", id}}, updatedClock)
	if err != nil {
		return nil, err
	}
	return &storage.Clock[primitive.ObjectID]{
		Id:        mongoStorage.ObjectId(clockId),
		Name:      updatedClock.Name,
		Avatar:    updatedClock.Avatar,
		Photos:    updatedClock.Photos,
		Price:     updatedClock.Price,
		Existence: updatedClock.Existence,
		Type:      updatedClock.Type,
	}, nil
}

func (manager *Clock) ClockById(id primitive.ObjectID) (*storage.Clock[primitive.ObjectID], error) {
	var foundClock storage.Clock[primitive.ObjectID]
	if err := manager.collection.FindOne(context.TODO(), bson.D{{"_id", id}}).Decode(&foundClock); err != nil {
		return nil, err
	}
	return &foundClock, nil
}

func (manager *Clock) AllClocks() ([]storage.Clock[primitive.ObjectID], error) {
	cursor, err := manager.collection.Find(context.TODO(), bson.D{{}})
	if err != nil {
		return nil, err
	}

	var allClocks []storage.Clock[primitive.ObjectID]
	if err = cursor.All(context.TODO(), &allClocks); err != nil {
		return nil, err
	}
	return allClocks, nil
}
