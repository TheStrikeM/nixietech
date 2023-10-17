package telegram

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
	"log/slog"
	"nixietech/internal/config"
	"nixietech/internal/fetcher"
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

	slog.Info(fmt.Sprintf("[TG] Authorized on account %s", bot.Self.UserName))
	return &Api{
		bot:     bot,
		config:  config,
		fetcher: fetcher,
	}
}

func (api *Api) IsAdmin(username string) bool {
	for _, admin := range api.config.Admins {
		if admin == username {
			return true
		}
	}
	return false
}

func (api *Api) StartUpdatesChecker() {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	updates := api.bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil {
			continue
		}

		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")
		if !api.IsAdmin(update.Message.From.UserName) {
			msg.Text = "Братик, ты не админ, не имеешь права на этого бота"
			if _, err := api.bot.Send(msg); err != nil {
				log.Fatal(err)
			}
			continue
		}

		switch update.Message.Command() {
		case "start":
			msg.ReplyToMessageID = update.Message.MessageID
			msg.Text = api.config.BotMessages.InitialHelp
		case "help":
			msg.ReplyToMessageID = update.Message.MessageID
			msg.Text = api.config.BotMessages.InitialHelp
		}

		if _, err := api.bot.Send(msg); err != nil {
			log.Panic(err)
		}
	}
}
