package main

import (
	"fmt"
	consts "nixietech/internal"
	"nixietech/internal/client/telegram"
	"nixietech/internal/fetcher"
	strikeMongo "nixietech/internal/storage/mongo"
	"nixietech/utils/config"
	"nixietech/utils/logger"
)

func main() {
	log := logger.SetupLogger("local", false)
	cfg := config.GetConfig(consts.ConfigName)

	log.Info(fmt.Sprintf("Go service success started. Mongo URI - %s", cfg.MongoURI))

	storage, disconnect := strikeMongo.New(cfg.MongoURI)
	defer disconnect()

	fetcher := fetcher.New(storage, cfg)
	botAPI := telegram.New(cfg, *fetcher)

	botAPI.StartUpdatesChecker(*fetcher)
}
