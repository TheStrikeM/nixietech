package messagerouter

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	consts "nixietech/internal"
	"nixietech/internal/client/telegram/routers"
	"nixietech/internal/client/telegram/storage"
	"nixietech/internal/config"
	"nixietech/internal/fetcher"
	cfg "nixietech/utils/config"
)

type MessageRouter struct {
	tgStorage *storage.TGStorageManager
	update    tgbotapi.Update
	fetcher   fetcher.Fetcher
	config    *config.Config
}

func New(tgStorage *storage.TGStorageManager,
	update tgbotapi.Update,
	fetcher fetcher.Fetcher,
) *MessageRouter {
	return &MessageRouter{
		tgStorage: tgStorage,
		update:    update,
		fetcher:   fetcher,
		config:    cfg.GetConfig(consts.ConfigName),
	}
}

func (router *MessageRouter) Route() tgbotapi.MessageConfig {
	if router.update.Message.Text == "/start" {

	}

	var message tgbotapi.MessageConfig

	switch router.tgStorage.Item(router.update.Message.From.UserName).Type {
	case routers.TGStorageCreateClock:

	}
}
