package main

import (
	"fmt"
	consts "nixietech/internal"
	"nixietech/internal/storage/mongo"
	"nixietech/utils/config"
	"nixietech/utils/logger"
)

func main() {
	log := logger.SetupLogger("local", false)
	cfg := config.GetConfig()

	log.Info(fmt.Sprintf("Go service success started. Mongo URI - %s", cfg.MongoURI))

	storage, disconnect := mongo.New(cfg.MongoURI)
	defer disconnect()

	storage.GetCollection(consts.CollectionClockName)
}
