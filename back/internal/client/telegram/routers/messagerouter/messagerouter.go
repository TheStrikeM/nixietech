package messagerouter

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	consts "nixietech/internal"
	"nixietech/internal/client/telegram/routers"
	"nixietech/internal/client/telegram/storage"
	"nixietech/internal/config"
	"nixietech/internal/fetcher"
	"nixietech/internal/fetcher/permissions"
	strg "nixietech/internal/storage"
	cfg "nixietech/utils/config"
	msg "nixietech/utils/telegram/message"
	"strconv"
	"strings"
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
	var message tgbotapi.MessageConfig

	if router.update.Message.Text == "/start" {
		message = routers.ShowMenu(router.update.Message.Chat.ID, router.update.Message.From, router.config)
		return message
	}

	user := router.tgStorage.Item(router.update.Message.From.UserName)
	if user == nil {
		message = UndefinedCommand(router.update, router.config)
		return message
	}

	switch user.Type {
	case routers.TGStorageCreateClock:
		message = CreateClock(router.update, router.config, router.tgStorage, router.fetcher)
	case routers.TGStorageDeleteClockById:
		message = DeleteClockById(router.update, router.config, router.tgStorage, router.fetcher)
	}

	return message
}

func DeleteClockById(update tgbotapi.Update, config *config.Config, tgStorage *storage.TGStorageManager, fetcher fetcher.Fetcher) tgbotapi.MessageConfig {
	if update.Message.Text != "Да" {
		return tgbotapi.NewMessage(update.Message.Chat.ID, "Либо жми отменить, либо пиши \"Да\"")
	}
	_, err := fetcher.RemoveClock(tgStorage.Item(update.Message.From.UserName).Message)
	if err != nil {
		panic(err)
	}

	tgStorage.RemoveItem(update.Message.From.UserName)

	successDeletedMessage := tgbotapi.NewMessage(update.Message.Chat.ID, config.BotMessages.SuccessDeletedClock)
	return successDeletedMessage
}

func CreateClock(update tgbotapi.Update, config *config.Config, tgStorage *storage.TGStorageManager, fetcher fetcher.Fetcher) tgbotapi.MessageConfig {
	parsedText := strings.Split(update.Message.Text, ";")

	tgStorage.RemoveItem(update.Message.From.UserName)

	price, err := strconv.Atoi(parsedText[1])
	if err != nil {
		panic(err)
	}
	var Type consts.ClockType
	if parsedText[3] == "1" {
		Type = consts.One
	} else {
		Type = consts.Two
	}
	_, err = fetcher.AddNewClock(strg.ClockWithoutId{
		Name:      parsedText[0],
		Avatar:    "Ничего нет",
		Photos:    []string{"1", "2"},
		Price:     price,
		Existence: permissions.ParseTrueFalse(parsedText[2]),
		Type:      Type,
	})
	if err != nil {
		panic(err)
	}

	return msg.MessageWithPrefix(config.BotMessages.ClockPrefix, fmt.Sprintf("Вы создали часы с названием %s и ценой %s", parsedText[0], parsedText[1]), update.Message.Chat.ID)
}

func UndefinedCommand(update tgbotapi.Update, config *config.Config) tgbotapi.MessageConfig {
	return tgbotapi.NewMessage(update.Message.Chat.ID, config.BotMessages.UndefinedMessage)
}
