package mongo

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"nixietech/internal/storage"
)

type IClock interface {
	AddClock(clock storage.ClockWithoutId) (*storage.Clock[primitive.ObjectID], error)
	RemoveClock(id primitive.ObjectID) (*primitive.ObjectID, error)
	UpdateClock(id primitive.ObjectID, updatedClock storage.ClockWithoutId) (*storage.Clock[primitive.ObjectID], error)
	ClockById(id primitive.ObjectID) (*storage.Clock[primitive.ObjectID], error)
	AllClocks() ([]storage.Clock[primitive.ObjectID], error)
}
