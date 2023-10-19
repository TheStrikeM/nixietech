package message

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func MessageWithPrefix(prefix string, message string, update *tgbotapi.Update) tgbotapi.MessageConfig {
	return tgbotapi.NewMessage(update.Message.Chat.ID, fmt.Sprintf("%s %s", prefix, message))
}
