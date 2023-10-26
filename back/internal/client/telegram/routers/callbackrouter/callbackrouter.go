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
	"strconv"
	"strings"
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

func (router *CallbackRouter) Route() []tgbotapi.MessageConfig {
	message := make([]tgbotapi.MessageConfig, 0, 10)

	//Advanced router
	if strings.Contains(router.update.CallbackQuery.Data, routers.CallbackDeleteClockById) {
		router.tgStorage.AddItem(router.update.CallbackQuery.From.UserName, routers.TGStorageDeleteClockById, strings.Split(router.update.CallbackQuery.Data, ";")[1])
		deleteMessage := tgbotapi.NewMessage(router.update.CallbackQuery.Message.Chat.ID, router.config.BotMessages.ConfirmDeleteClock)
		deleteMessage.ReplyMarkup = tgbotapi.NewInlineKeyboardMarkup(tgbotapi.NewInlineKeyboardRow(tgbotapi.NewInlineKeyboardButtonData("[❌] Отмена", routers.CallbackRemoveItemStorage)))
		message = append(message, deleteMessage)
		return message
	}

	if strings.Contains(router.update.CallbackQuery.Data, routers.CallbackUpdateClockById) {
		router.tgStorage.AddItem(router.update.CallbackQuery.From.UserName, routers.TGStorageUpdateClockById, strings.Split(router.update.CallbackQuery.Data, ";")[1])
		updateMessage := tgbotapi.NewMessage(router.update.CallbackQuery.Message.Chat.ID, "Братик, ты уверен, что хочешь начать изменение?")
		updateMessage.ReplyMarkup = tgbotapi.NewInlineKeyboardMarkup(
			tgbotapi.NewInlineKeyboardRow(
				tgbotapi.NewInlineKeyboardButtonData("[❌] Нет", routers.CallbackRemoveItemStorage),
				tgbotapi.NewInlineKeyboardButtonData("[✅] Да", routers.TGStorageUpdateClockById),
			),
		)
		message = append(message, updateMessage)
		return message
	}

	// Default router
	switch router.update.CallbackQuery.Data {
	case routers.CallbackShowMenu:
		message = append(message, routers.ShowMenu(router.update.CallbackQuery.Message.Chat.ID, router.update.CallbackQuery.From, router.config))
	case routers.CallbackShowClockMenu:
		if permissions.UserGroup(router.update.CallbackQuery.From.UserName).Clock.ShowClockMenu {
			message = append(message, ShowClockMenu(&router.update, router.config))
		} else {
			message = append(message, ShowDontPerms(&router.update, router.config, router.config.BotMessages.ClockPrefix))
		}
	case routers.CallbackCreateClock:
		if permissions.UserGroup(router.update.CallbackQuery.From.UserName).Clock.CreateClock {
			message = append(message, ShowCreateClock(&router.update, router.config, router.tgStorage))
		} else {
			message = append(message, ShowDontPerms(&router.update, router.config, router.config.BotMessages.ClockPrefix))
		}
	case routers.CallbackRemoveItemStorage:
		message = append(message, ShowRemoveItemStorage(&router.update, router.config, router.tgStorage))
	case routers.CallbackShowAllClocks:
		for _, itemMessage := range ShowAllClocks(&router.update, router.config, router.fetcher) {
			message = append(message, itemMessage)
		}
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
			tgbotapi.NewInlineKeyboardButtonData("[❎] Добавить", routers.CallbackCreateClock),
			tgbotapi.NewInlineKeyboardButtonData("[📑] Показать все", routers.CallbackShowAllClocks),
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

func ShowAllClocks(update *tgbotapi.Update, config *config.Config, fetcherManager fetcher.Fetcher) []tgbotapi.MessageConfig {
	parsedText := fetcher.ParseHashTags(config.BotMessages.ClockShowAll, []fetcher.HashTags{
		permissions.ParseMinimumPerms(update.CallbackQuery.From),
		fetcher.ParsePrefix(config.BotMessages.ClockPrefix),
	})

	messages := make([]tgbotapi.MessageConfig, 0, 5)
	messages = append(messages, tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID, parsedText))

	clocks, err := fetcherManager.GetAll()
	if err != nil {
		panic(err)
	}

	for _, clock := range clocks {
		clockText := fetcher.ParseHashTags(config.BotMessages.ClockItemMessage, []fetcher.HashTags{
			fetcher.NewTag("$CLOCK_NAME$", clock.Name),
			fetcher.NewTag("$CLOCK_PRICE$", strconv.Itoa(clock.Price)),
			fetcher.NewTag("$CLOCK_EXISTENCE$", strconv.FormatBool(clock.Existence)),
			fetcher.NewTag("$CLOCK_TYPE$", strconv.Itoa(consts.ClockTypeToInt(clock.Type))),
		})
		clockMessage := tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID, clockText)
		clockMessage.ReplyMarkup = tgbotapi.NewInlineKeyboardMarkup(
			tgbotapi.NewInlineKeyboardRow(
				tgbotapi.NewInlineKeyboardButtonData("[❌] Удалить", strings.Join([]string{routers.CallbackDeleteClockById, clock.Id.String()}, ";")),
				tgbotapi.NewInlineKeyboardButtonData("[✏️] Изменить", strings.Join([]string{routers.CallbackUpdateClockById, clock.Id.String()}, ";")),
			),
		)
		messages = append(messages, clockMessage)
	}

	backToClockMenu := tgMessage.MessageWithPrefix(config.BotMessages.ClockPrefix, "Вернуть в меню часов", update.CallbackQuery.Message.Chat.ID)
	backToClockMenu.ReplyMarkup = tgbotapi.NewInlineKeyboardMarkup(tgbotapi.NewInlineKeyboardRow(tgbotapi.NewInlineKeyboardButtonData("[⚜️] Вернуться", routers.CallbackShowClockMenu)))
	messages = append(messages, backToClockMenu)

	return messages
}

func ShowRemoveItemStorage(update *tgbotapi.Update, config *config.Config, tgStorage *storage.TGStorageManager) tgbotapi.MessageConfig {
	tgStorage.RemoveItem(update.CallbackQuery.From.UserName)
	return tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID, config.BotMessages.RemoveItemStorage)
}

func ShowDontPerms(update *tgbotapi.Update, config *config.Config, prefix string) tgbotapi.MessageConfig {
	return tgMessage.MessageWithPrefix(prefix, config.BotMessages.PermissionDenied, update.CallbackQuery.Message.Chat.ID)
}
