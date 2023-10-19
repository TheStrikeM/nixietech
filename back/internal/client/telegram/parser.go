package telegram

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"strings"
)

type HashTags struct {
	message string
	item    string
}

func NewTag(message string, item string) HashTags {
	return HashTags{
		message: message,
		item:    item,
	}
}

func ParseGroup(message string) *PermissionGroup {
	splitedMsg := strings.Split(message, ";")
	return &PermissionGroup{
		Prefix: splitedMsg[0],
		Users:  strings.Split(splitedMsg[len(splitedMsg)-1], ","),
		Clock: ClockPermissions{
			ShowClockMenu: ParseTrueFalse(splitedMsg[1]),
			CreateClock:   ParseTrueFalse(splitedMsg[2]),
			DeleteClock:   ParseTrueFalse(splitedMsg[3]),
			UpdateClock:   ParseTrueFalse(splitedMsg[4]),
		},
		Order: OrderPermissions{
			ShowOrderMenu: ParseTrueFalse(splitedMsg[5]),
			DeleteOrder:   ParseTrueFalse(splitedMsg[6]),
		},
	}
}

func ParseHashTags(message string, items []HashTags) string {
	for _, item := range items {
		message = strings.Replace(message, item.message, item.item, -1)
	}
	return message
}

func ParseTrueFalse(message string) bool {
	if message == "✅" {
		return true
	}
	return false
}

func ParseFullPerms(update *tgbotapi.Update, api *Api) HashTags {
	return NewTag("$permission-prefix-full$", ParseHashTags(
		api.config.BotMessages.PermissionPrefixFull,
		[]HashTags{
			NewTag("$ROLE$", UserGroup(*update).Prefix),
		},
	))
}

func ParseMinimumPerms(update *tgbotapi.Update, api *Api) HashTags {
	return NewTag("$permission-prefix-minimum$", ParseHashTags(
		api.config.BotMessages.PermissionPrefixMinimum,
		[]HashTags{
			ParseGroupPrefix(update),
			ParseFullName(update),
		},
	))
}

func ParseGroupPrefix(update *tgbotapi.Update) HashTags {
	return NewTag("$ROLE$", UserGroup(*update).Prefix)
}

func ParseFullName(update *tgbotapi.Update) HashTags {
	if update.Message != nil {
		return NewTag("$NAME$", fmt.Sprintf("%s %s", update.Message.From.FirstName, update.Message.From.LastName))
	}
	return NewTag("$NAME$", fmt.Sprintf("%s %s", update.CallbackQuery.From.FirstName, update.CallbackQuery.From.LastName))
}

func ParsePrefix(update *tgbotapi.Update, message string) HashTags {
	return NewTag("$PREFIX$", message)
}
