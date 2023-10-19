package telegram

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

func CreateKeyboard(buttonText, buttonCode string) []tgbotapi.InlineKeyboardButton {
	return tgbotapi.NewInlineKeyboardRow(tgbotapi.NewInlineKeyboardButtonData(buttonText, buttonCode))
}
