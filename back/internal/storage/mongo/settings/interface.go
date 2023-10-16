package settings

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"nixietech/internal/storage"
)

type ISettings[
	SettingsWithId storage.TypeOneSettings[primitive.ObjectID] | storage.TypeTwoSettings[primitive.ObjectID],
	SettingsWithoutId storage.TypeOneSettingsWithoutId | storage.TypeTwoSettingsWithoutId,
] interface {
	AddSettings(settings SettingsWithoutId) (SettingsWithId, error)
	RemoveSettings(id primitive.ObjectID) (SettingsWithId, error)
	UpdateSettings(id primitive.ObjectID, updatedSettings SettingsWithoutId) (SettingsWithId, error)
	SettingsById(id primitive.ObjectID) (SettingsWithId, error)
}
