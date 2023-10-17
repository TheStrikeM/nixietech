package order

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"nixietech/internal/storage"
)

type IOrder interface {
	AddOrder(order storage.OrderWithoutId[primitive.ObjectID]) (*storage.Order[primitive.ObjectID], error)
	RemoveOrder(id primitive.ObjectID) (*primitive.ObjectID, error)
	UpdateOrder(id primitive.ObjectID, updatedOrder storage.OrderWithoutId[primitive.ObjectID]) (*storage.Order[primitive.ObjectID], error)
	OrderById(id primitive.ObjectID) (*storage.Order[primitive.ObjectID], error)
	AllOrders() ([]storage.Order[primitive.ObjectID], error)
}
