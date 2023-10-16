package main

import (
	"fmt"
	"nixietech/utils/config"
	"nixietech/utils/logger"
)

func main() {
	log := logger.SetupLogger("local", false)
	cfg := config.GetConfig()

	log.Info(fmt.Sprintf("Go service success started. Mongo URI - %s", cfg.MongoURI))
}
