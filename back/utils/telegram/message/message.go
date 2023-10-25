package message

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func MessageWithPrefix(prefix string, message string, chatId int64) tgbotapi.MessageConfig {
	return tgbotapi.NewMessage(chatId, fmt.Sprintf("%s %s", prefix, message))
}
