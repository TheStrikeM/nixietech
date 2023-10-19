package fetcher

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"nixietech/internal/config"
	"nixietech/internal/storage"
	mongoStorage "nixietech/internal/storage/mongo"
	mongoClock "nixietech/internal/storage/mongo/clock"
	mongoOrder "nixietech/internal/storage/mongo/order"
	mongoSettings "nixietech/internal/storage/mongo/settings"
)

type Fetcher struct {
	storage                  *mongoStorage.Storage
	config                   *config.Config
	clockDbManager           *mongoClock.Clock
	orderDbManager           *mongoOrder.Order
	typeOneSettingsDbManager *mongoSettings.TypeOneSettings
	typeTwoSettingsDbManager *mongoSettings.TypeTwoSettings
}

func New(storage *mongoStorage.Storage, config *config.Config) *Fetcher {
	return &Fetcher{
		storage:                  storage,
		config:                   config,
		clockDbManager:           mongoClock.New(storage),
		orderDbManager:           mongoOrder.New(storage),
		typeOneSettingsDbManager: mongoSettings.NewOne(storage),
		typeTwoSettingsDbManager: (*mongoSettings.TypeTwoSettings)(mongoSettings.NewTwo(storage)),
	}
}

func (fetch *Fetcher) AddNewClock(clock storage.ClockWithoutId) (*storage.Clock[primitive.ObjectID], error) {
	clockItem, err := fetch.clockDbManager.AddClock(clock)
	if err != nil {
		return nil, err
	}
	return clockItem, nil
}

func (fetch *Fetcher) RemoveClock(id primitive.ObjectID) (*primitive.ObjectID, error) {
	clockId, err := fetch.clockDbManager.RemoveClock(id)
	if err != nil {
		return nil, err
	}
	return clockId, nil
}

func (fetch *Fetcher) GetAll() ([]storage.Clock[primitive.ObjectID], error) {
	allClocks, err := fetch.clockDbManager.AllClocks()
	if err != nil {
		return nil, err
	}
	return allClocks, nil
}

func (fetch *Fetcher) ClockById(id primitive.ObjectID) (*storage.Clock[primitive.ObjectID], error) {
	clock, err := fetch.clockDbManager.ClockById(id)
	if err != nil {
		return nil, err
	}
	return clock, nil
}
