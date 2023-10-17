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

type TypeOneSettings struct {
	collection *mongo.Collection
	storage    *mongoStorage.Storage
}

func NewOne(storage *mongoStorage.Storage) *TypeOneSettings {
	return &TypeOneSettings{
		storage:    storage,
		collection: storage.GetCollection(consts.CollectionOrderName),
	}
}

func (manager *TypeOneSettings) AddSettings(settings storage.TypeOneSettingsWithoutId) (*storage.TypeOneSettings[primitive.ObjectID], error) {
	orderId, err := manager.collection.InsertOne(context.TODO(), settings)
	if err != nil {
		return nil, err
	}
	return &storage.TypeOneSettings[primitive.ObjectID]{
		Id:                  mongoStorage.ObjectId(mongoStorage.InsertedIdToString(orderId.InsertedID)),
		LampsType:           settings.LampsType,
		DecorativeRingsType: settings.DecorativeRingsType,
	}, nil
}

func (manager *TypeOneSettings) RemoveSettings(id primitive.ObjectID) (*primitive.ObjectID, error) {
	if _, err := manager.collection.DeleteOne(context.TODO(), bson.D{{"_id", id}}); err != nil {
		return nil, err
	}
	return &id, nil
}

func (manager *TypeOneSettings) UpdateSettings(
	id primitive.ObjectID,
	updatedSettings storage.TypeOneSettingsWithoutId,
) (*storage.TypeOneSettings[primitive.ObjectID], error) {
	settingsId, err := manager.collection.UpdateOne(context.TODO(), bson.D{{"_id", id}}, updatedSettings)
	if err != nil {
		return nil, err
	}
	return &storage.TypeOneSettings[primitive.ObjectID]{
		Id:                  mongoStorage.ObjectId(mongoStorage.InsertedIdToString(settingsId.UpsertedID)),
		LampsType:           updatedSettings.LampsType,
		DecorativeRingsType: updatedSettings.DecorativeRingsType,
	}, nil
}

func (manager *TypeOneSettings) SettingsById(id primitive.ObjectID) (*storage.TypeOneSettings[primitive.ObjectID], error) {
	var foundSettings storage.TypeOneSettings[primitive.ObjectID]
	if err := manager.collection.FindOne(context.TODO(), bson.D{{"_id", id}}).Decode(&foundSettings); err != nil {
		return nil, err
	}
	return &foundSettings, nil
}

func (manager *TypeOneSettings) AllSettings() ([]storage.TypeOneSettings[primitive.ObjectID], error) {
	cursor, err := manager.collection.Find(context.TODO(), bson.D{{}})
	if err != nil {
		return nil, err
	}

	var allSettings []storage.TypeOneSettings[primitive.ObjectID]
	if err = cursor.All(context.TODO(), &allSettings); err != nil {
		return nil, err
	}
	return allSettings, nil
}
