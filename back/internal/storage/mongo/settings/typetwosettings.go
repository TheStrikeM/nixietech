package settings

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	consts "nixietech/internal"
	"nixietech/internal/storage"
	mongoStorage "nixietech/internal/storage/mongo"
)

type TypeTwoSettings struct {
	collection *mongo.Collection
	storage    *mongoStorage.Storage
}

func NewTwo(storage *mongoStorage.Storage) *TypeOneSettings {
	return &TypeOneSettings{
		storage:    storage,
		collection: storage.GetCollection(consts.CollectionOrderName),
	}
}

func (manager *TypeTwoSettings) AddSettings(settings storage.TypeTwoSettingsWithoutId) (*storage.TypeTwoSettings[primitive.ObjectID], error) {
	orderId, err := manager.collection.InsertOne(context.TODO(), settings)
	if err != nil {
		return nil, err
	}
	return &storage.TypeTwoSettings[primitive.ObjectID]{
		Id:   mongoStorage.ObjectId(mongoStorage.InsertedIdToString(orderId.InsertedID)),
		Test: settings.Test,
	}, nil
}

func (manager *TypeTwoSettings) RemoveSettings(id primitive.ObjectID) (*primitive.ObjectID, error) {
	if _, err := manager.collection.DeleteOne(context.TODO(), bson.D{{"_id", id}}); err != nil {
		return nil, err
	}
	return &id, nil
}

func (manager *TypeTwoSettings) UpdateSettings(
	id primitive.ObjectID,
	updatedSettings storage.TypeTwoSettingsWithoutId,
) (*storage.TypeTwoSettings[primitive.ObjectID], error) {
	settingsId, err := manager.collection.UpdateOne(context.TODO(), bson.D{{"_id", id}}, updatedSettings)
	if err != nil {
		return nil, err
	}
	return &storage.TypeTwoSettings[primitive.ObjectID]{
		Id:   mongoStorage.ObjectId(mongoStorage.InsertedIdToString(settingsId.UpsertedID)),
		Test: updatedSettings.Test,
	}, nil
}

func (manager *TypeTwoSettings) SettingsById(id primitive.ObjectID) (*storage.TypeTwoSettings[primitive.ObjectID], error) {
	var foundSettings storage.TypeTwoSettings[primitive.ObjectID]
	if err := manager.collection.FindOne(context.TODO(), bson.D{{"_id", id}}).Decode(&foundSettings); err != nil {
		return nil, err
	}
	return &foundSettings, nil
}

func (manager *TypeTwoSettings) AllSettings() ([]storage.TypeTwoSettings[primitive.ObjectID], error) {
	cursor, err := manager.collection.Find(context.TODO(), bson.D{{}})
	if err != nil {
		return nil, err
	}

	var allSettings []storage.TypeTwoSettings[primitive.ObjectID]
	if err = cursor.All(context.TODO(), &allSettings); err != nil {
		return nil, err
	}
	return allSettings, nil
}
