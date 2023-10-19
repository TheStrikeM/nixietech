package telegram

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
	"log/slog"
	"nixietech/internal/client/telegram/callbackquery"
	"nixietech/internal/client/telegram/conditions"
	"nixietech/internal/client/telegram/permissions"
	"nixietech/internal/config"
	"nixietech/internal/fetcher"
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

	slog.Info(fmt.Sprintf("[TG] Authorized on account %s", bot.Self.UserName))
	return &Api{
		bot:     bot,
		config:  config,
		fetcher: fetcher,
	}
}

func showMenu(update *tgbotapi.Update, api *Api) {
	initialHelpMessage := api.config.BotMessages.InitialHelpMessage
	permissionsPrefixFull := api.config.BotMessages.PermissionPrefixFull
	parsedMessage := permissions.ParseHashTags(initialHelpMessage, []permissions.HashTags{
		permissions.NewTag("$NAME$", fmt.Sprintf("%s %s", update.Message.From.FirstName, update.Message.From.LastName)),
		permissions.NewTag("$ROLE$", permissions.UserGroup(*update).Prefix),
		permissions.NewTag("$permission-prefix-full$", permissions.ParseHashTags(
			permissionsPrefixFull,
			[]permissions.HashTags{
				permissions.NewTag("$ROLE$", permissions.UserGroup(*update).Prefix),
			},
		)),
	})

	welcomeMessage := tgbotapi.NewMessage(update.Message.Chat.ID, parsedMessage)
	welcomeMessage.ReplyMarkup = tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("[⚜️] Часы", ShowClockMenu),
			tgbotapi.NewInlineKeyboardButtonData("[👨‍💻] Заказы", ShowOrderMenu),
		),
	)
	if _, err := api.bot.Send(welcomeMessage); err != nil {
		panic(err)
	}
}

func (api *Api) StartUpdatesChecker() {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	updates := api.bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil {
			continue
		}

		if conditions.IsPermissionGroupNil(update) {
			stopMessage := message.MessageWithPrefix(
				api.config.BotMessages.GlobalPrefix,
				"Ты не принадлежишь ни к одной группе прав, поэтому доступ запрещен❌",
				&update,
			)
			if _, err := api.bot.Send(stopMessage); err != nil {
				panic(err)
			}
			continue
		}

		if conditions.IsStartMessage(update) {
			showMenu(&update, api)
		}

		if conditions.IsCallbackQuery(update) {
			callbackquery.CallbackQueryRouter(update.CallbackQuery.Data)
		}

	}
}
