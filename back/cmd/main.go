package main

import (
	"fmt"
	"log/slog"
	strikeMongo "nixietech/internal/storage/mongo"
	clock "nixietech/internal/storage/mongo/clock"
	"nixietech/utils/config"
	"nixietech/utils/logger"
)

func main() {
	log := logger.SetupLogger("local", false)
	cfg := config.GetConfig()

	log.Info(fmt.Sprintf("Go service success started. Mongo URI - %s", cfg.MongoURI))

	storage, disconnect := strikeMongo.New(cfg.MongoURI)
	defer disconnect()

	clockManager := clock.New(storage)
	//clockItem, err := clockManager.AddClock(storage2.ClockWithoutId{
	//	Name:      "Gena",
	//	Avatar:    "net",
	//	Photos:    []string{"photo1", "photo2"},
	//	Price:     10 * 2,
	//	Existence: false,
	//	Type:      consts.One,
	//})
	//if err != nil {
	//	panic(err)
	//}

	allClocks, err := clockManager.AllClocks()
	if err != nil {
		panic(err)
	}

	for _, item := range allClocks {
		slog.Info(fmt.Sprintf("Id is %s", item.Id))
	}

	lastClock, err := clockManager.ClockById(strikeMongo.ObjectId("652ea8a56c2f06dcf4c55de3"))
	slog.Info(fmt.Sprintf("Name of this clock is %s", lastClock.Name))
}
