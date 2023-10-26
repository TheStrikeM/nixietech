package routers

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"nixietech/internal/config"
	"nixietech/internal/fetcher"
	"nixietech/internal/fetcher/permissions"
)

func ShowMenu(chatId int64, from *tgbotapi.User, config *config.Config) tgbotapi.MessageConfig {

	parsedText := fetcher.ParseHashTags(config.BotMessages.InitialHelpMessage, []fetcher.HashTags{
		permissions.ParseFullPerms(from.UserName),
		fetcher.ParseFullName(from.FirstName, from.LastName),
	})
	message := tgbotapi.NewMessage(chatId, parsedText)
	message.ReplyMarkup = tgbotapi.NewInlineKeyboardMarkup(tgbotapi.NewInlineKeyboardRow(tgbotapi.NewInlineKeyboardButtonData("[⚜️] Часы", CallbackShowClockMenu)))
	return message
}
