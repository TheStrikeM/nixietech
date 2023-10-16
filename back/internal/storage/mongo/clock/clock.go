package clock

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"nixietech/internal/storage"
)

type Clock struct{}

func (c Clock) AddClock(clock storage.ClockWithoutId) (storage.Clock[primitive.ObjectID], error) {}
