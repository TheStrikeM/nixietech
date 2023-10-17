package settings

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"nixietech/internal/storage"
)

type ITypeOneSettings interface {
	AddSettings(order storage.TypeOneSettingsWithoutId) (*storage.TypeOneSettings[primitive.ObjectID], error)
	RemoveSettings(id primitive.ObjectID) (*primitive.ObjectID, error)
	UpdateSettings(id primitive.ObjectID, updatedSettings storage.TypeOneSettingsWithoutId) (*storage.TypeOneSettings[primitive.ObjectID], error)
	OrderById(id primitive.ObjectID) (*storage.TypeOneSettings[primitive.ObjectID], error)
	AllSettings() ([]storage.TypeOneSettings[primitive.ObjectID], error)
}

type ITypeTwoSettings interface {
	AddSettings(order storage.TypeTwoSettingsWithoutId) (*storage.TypeTwoSettings[primitive.ObjectID], error)
	RemoveSettings(id primitive.ObjectID) (*primitive.ObjectID, error)
	UpdateSettings(id primitive.ObjectID, updatedSettings storage.TypeTwoSettingsWithoutId) (*storage.TypeTwoSettings[primitive.ObjectID], error)
	OrderById(id primitive.ObjectID) (*storage.TypeTwoSettings[primitive.ObjectID], error)
	AllSettings() ([]storage.TypeTwoSettings[primitive.ObjectID], error)
}
