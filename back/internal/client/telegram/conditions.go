package telegram

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"nixietech/internal/fetcher/permissions"
)

func IsPermissionGroupNil(update tgbotapi.Update) bool {
	return update.Message != nil && permissions.UserGroup(update.Message.From.UserName) == nil
}

func IsCallbackQuery(update tgbotapi.Update) bool {
	return update.CallbackQuery != nil && update.CallbackQuery.Data != ""
}

func isMessageQuery(update tgbotapi.Update) bool {
	return update.Message != nil && update.Message.Text != ""
}
