package telegram

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
	"log/slog"
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

func (api *Api) Bot() *tgbotapi.BotAPI {
	return api.bot
}

func (api *Api) Config() *config.Config {
	return api.config
}

func (api *Api) StartUpdatesChecker() {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	for update := range api.bot.GetUpdatesChan(u) {
		if IsCallbackQuery(update) {
			QueryRouter(&update, api)
		}

		if update.Message == nil {
			continue
		}

		if IsPermissionGroupNil(update) {
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

		if IsStartMessage(update) {
			ShowGlMenu(&update, api)
		}

	}
}

func QueryRouter(update *tgbotapi.Update, api *Api) {
	switch update.CallbackQuery.Data {
	case ShowClockMenu:
		ShowClockMenuFunc(update, api)
	case ShowOrderMenu:
		slog.Info("Пока здесь ничего нет")
	case ShowMenu:
		ShowGlMenu(update, api)
	}
}

func ShowGlMenu(update *tgbotapi.Update, api *Api) {
	initialHelpMessage := api.config.BotMessages.InitialHelpMessage
	parsedMessage := ParseHashTags(initialHelpMessage, []HashTags{
		ParseFullName(update),
		ParseGroupPrefix(update),
		ParseFullPerms(update, api),
	})

	var welcomeMessage tgbotapi.MessageConfig
	if update.Message != nil {
		welcomeMessage = tgbotapi.NewMessage(update.Message.Chat.ID, parsedMessage)
	} else {
		welcomeMessage = tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID, parsedMessage)
	}

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

func ShowClockMenuFunc(update *tgbotapi.Update, api *Api) {
	parsedMessage := ParseHashTags(api.config.BotMessages.ClockMenuMessage, []HashTags{
		ParseMinimumPerms(update, api),
		ParsePrefix(update, api.config.BotMessages.ClockPrefix),
	})
	finalMessage := tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID, parsedMessage)
	finalMessage.ReplyMarkup = tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("[❎] Добавить новый", InDev),
			tgbotapi.NewInlineKeyboardButtonData("[📑] Статистика (dev)", CreateNewClock),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("[📖] В главное меню", ShowMenu)),
	)
	if _, err := api.bot.Send(finalMessage); err != nil {
		panic(err)
	}
}
