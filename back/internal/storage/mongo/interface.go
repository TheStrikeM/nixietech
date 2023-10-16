package mongo

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"nixietech/internal/storage"
)

type Setting = storage.TypeOneSettings[primitive.ObjectID]

type StorageInterface interface {
	New(string) (*Storage, func())

	// AddClock Clock CRUD
	AddClock(clock storage.ClockWithoutId) storage.Clock[primitive.ObjectID]
	RemoveClock(id primitive.ObjectID) storage.Clock[primitive.ObjectID]
	UpdateClock(id primitive.ObjectID, updatedClock storage.ClockWithoutId) storage.Clock[primitive.ObjectID]
	ClockById(id primitive.ObjectID) storage.Clock[primitive.ObjectID]

	// AddOrder Order CRUD
	AddOrder(order storage.OrderWithoutId[primitive.ObjectID]) storage.Order[primitive.ObjectID]
	RemoveOrder(id primitive.ObjectID) storage.Order[primitive.ObjectID]
	UpdateOrder(id primitive.ObjectID, updatedOrder storage.OrderWithoutId[primitive.ObjectID])
	OrderById(id primitive.ObjectID) storage.Order[primitive.ObjectID]

	AddSettings(settings storage.TypeOneSettingsWithoutId) storage.TypeOneSettings[primitive.ObjectID]
}
