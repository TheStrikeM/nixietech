package telegram

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func IsPermissionGroupNil(update tgbotapi.Update) bool {
	return UserGroup(update) == nil && update.Message != nil
}

func IsStartMessage(update tgbotapi.Update) bool {
	return update.Message.Command() == "start" && update.Message != nil
}

func IsCallbackQuery(update tgbotapi.Update) bool {
	return update.CallbackQuery != nil && update.CallbackQuery.Data != ""
}
