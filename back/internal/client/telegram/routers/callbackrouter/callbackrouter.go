package callbackrouter

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	consts "nixietech/internal"
	"nixietech/internal/client/telegram/routers"
	"nixietech/internal/client/telegram/storage"
	"nixietech/internal/config"
	"nixietech/internal/fetcher"
	"nixietech/internal/fetcher/permissions"
	cfg "nixietech/utils/config"
	tgMessage "nixietech/utils/telegram/message"
)

type CallbackRouter struct {
	tgStorage *storage.TGStorageManager
	update    tgbotapi.Update
	fetcher   fetcher.Fetcher
	config    *config.Config
}

func New(tgStorage *storage.TGStorageManager,
	update tgbotapi.Update,
	fetcher fetcher.Fetcher,
) *CallbackRouter {
	return &CallbackRouter{
		tgStorage: tgStorage,
		update:    update,
		fetcher:   fetcher,
		config:    cfg.GetConfig(consts.ConfigName),
	}
}

func (router *CallbackRouter) Route() tgbotapi.MessageConfig {
	var message tgbotapi.MessageConfig

	switch router.update.CallbackQuery.Data {
	case routers.CallbackShowClockMenu:
		if permissions.UserGroup(router.update.Message.From.UserName).Clock.ShowClockMenu {
			message = ShowClockMenu(&router.update, router.config)
		} else {
			message = ShowDontPerms(&router.update, router.config, router.config.BotMessages.ClockPrefix)
		}
	case routers.CallbackCreateClock:
		if permissions.UserGroup(router.update.Message.From.UserName).Clock.CreateClock {
			message = ShowCreateClock(&router.update, router.config, router.tgStorage)
		} else {
			message = ShowDontPerms(&router.update, router.config, router.config.BotMessages.ClockPrefix)
		}
	case routers.CallbackRemoveItemStorage:
		message = ShowRemoveItemStorage(&router.update, router.config, router.tgStorage)
	}

	return message
}

func ShowClockMenu(update *tgbotapi.Update, config *config.Config) tgbotapi.MessageConfig {
	parsedText := fetcher.ParseHashTags(config.BotMessages.ClockMenuMessage, []fetcher.HashTags{
		permissions.ParseMinimumPerms(update.CallbackQuery.From),
		fetcher.ParsePrefix(config.BotMessages.ClockPrefix),
	})

	message := tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID, parsedText)
	message.ReplyMarkup = tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("[❎] Добавить новый", routers.CallbackCreateClock),
			tgbotapi.NewInlineKeyboardButtonData("[📑] Статистика (dev)", routers.CallbackInDev),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("[📖] В главное меню", routers.CallbackShowMenu)),
	)

	return message
}

func ShowCreateClock(update *tgbotapi.Update, config *config.Config, tgStorage *storage.TGStorageManager) tgbotapi.MessageConfig {
	message := tgMessage.MessageWithPrefix(config.BotMessages.ClockPrefix, config.BotMessages.ClockCreateStart, update.CallbackQuery.Message.Chat.ID)
	message.ReplyMarkup = tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("[❌] Отменить действие", routers.CallbackRemoveItemStorage),
		),
	)

	tgStorage.AddItem(update.CallbackQuery.From.UserName, routers.TGStorageCreateClock, "")
	return message
}

func ShowRemoveItemStorage(update *tgbotapi.Update, config *config.Config, tgStorage *storage.TGStorageManager) tgbotapi.MessageConfig {
	tgStorage.RemoveItem(update.CallbackQuery.From.UserName)
	return tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID, config.BotMessages.RemoveItemStorage)
}

func ShowDontPerms(update *tgbotapi.Update, config *config.Config, prefix string) tgbotapi.MessageConfig {
	return tgMessage.MessageWithPrefix(prefix, config.BotMessages.PermissionDenied, update.CallbackQuery.Message.Chat.ID)
}
