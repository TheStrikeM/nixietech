package telegram

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
	"log/slog"
	"nixietech/internal/client/telegram/storage"
	"nixietech/internal/config"
	"nixietech/internal/fetcher"
	"nixietech/utils/e"
	"nixietech/utils/telegram/message"
)

type Api struct {
	bot     *tgbotapi.BotAPI
	config  *config.Config
	fetcher fetcher.Fetcher
}

func New(config *config.Config, fetcher fetcher.Fetcher) *Api {
	bot, err := tgbotapi.NewBotAPI(config.TelegramApiToken)
	if err != nil {
		log.Panic(err)
	}

	slog.Info(fmt.Sprintf("[TG] Success authorized on account %s", bot.Self.UserName))
	return &Api{
		bot:     bot,
		config:  config,
		fetcher: fetcher,
	}
}

func (api *Api) Bot() *tgbotapi.BotAPI {
	return api.bot
}

func (api *Api) Config() *config.Config {
	return api.config
}

func (api *Api) StartUpdatesChecker(fetcherApi fetcher.Fetcher) (err error) {
	defer func() { err = e.WrapIfErr("start updates checker error", err) }()
	fetcherManager = fetcherApi
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	for update := range api.bot.GetUpdatesChan(u) {
		if IsPermissionGroupNil(update) {
			stopMessage := message.MessageWithPrefix(
				api.config.BotMessages.GlobalPrefix,
				api.config.BotMessages.PermissionDenied,
				update.Message.Chat.ID,
			)
			if _, err := api.bot.Send(stopMessage); err != nil {
				return err
			}
			continue
		}

		tgStorage := storage.New()

		if isMessageQuery(update) {

		}
		if IsCallbackQuery(update) {

		}

	}
}
