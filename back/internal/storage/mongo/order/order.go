package order

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	consts "nixietech/internal"
	"nixietech/internal/storage"
	mongoStorage "nixietech/internal/storage/mongo"
)

type Order struct {
	collection *mongo.Collection
	storage    *mongoStorage.Storage
}

func New(storage *mongoStorage.Storage) *Order {
	return &Order{
		storage:    storage,
		collection: storage.GetCollection(consts.CollectionOrderName),
	}
}

func (manager *Order) AddOrder(order storage.OrderWithoutId[primitive.ObjectID]) (*storage.Order[primitive.ObjectID], error) {
	orderId, err := manager.collection.InsertOne(context.TODO(), order)
	if err != nil {
		return nil, err
	}
	return &storage.Order[primitive.ObjectID]{
		Id:      mongoStorage.ObjectId(mongoStorage.InsertedIdToString(orderId.InsertedID)),
		Wishes:  order.Wishes,
		Contact: order.Contact,
		ClockId: order.ClockId,
		Base:    order.Base,
	}, nil
}

func (manager *Order) RemoveOrder(id primitive.ObjectID) (*primitive.ObjectID, error) {
	if _, err := manager.collection.DeleteOne(context.TODO(), bson.D{{"_id", id}}); err != nil {
		return nil, err
	}
	return &id, nil
}

func (manager *Order) UpdateOrder(
	id primitive.ObjectID,
	updatedOrder storage.OrderWithoutId[primitive.ObjectID],
) (*storage.Order[primitive.ObjectID], error) {
	orderId, err := manager.collection.UpdateOne(context.TODO(), bson.D{{"_id", id}}, updatedOrder)
	if err != nil {
		return nil, err
	}
	return &storage.Order[primitive.ObjectID]{
		Id:      mongoStorage.ObjectId(mongoStorage.InsertedIdToString(orderId.UpsertedID)),
		Wishes:  updatedOrder.Wishes,
		Contact: updatedOrder.Contact,
		ClockId: updatedOrder.ClockId,
		Base:    updatedOrder.Base,
	}, nil
}

func (manager *Order) OrderById(id primitive.ObjectID) (*storage.Order[primitive.ObjectID], error) {
	var foundOrder storage.Order[primitive.ObjectID]
	if err := manager.collection.FindOne(context.TODO(), bson.D{{"_id", id}}).Decode(&foundOrder); err != nil {
		return nil, err
	}
	return &foundOrder, nil
}

func (manager *Order) AllOrders() ([]storage.Order[primitive.ObjectID], error) {
	cursor, err := manager.collection.Find(context.TODO(), bson.D{{}})
	if err != nil {
		return nil, err
	}

	var allOrders []storage.Order[primitive.ObjectID]
	if err = cursor.All(context.TODO(), &allOrders); err != nil {
		return nil, err
	}
	return allOrders, nil
}
