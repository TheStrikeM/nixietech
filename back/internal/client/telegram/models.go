package telegram

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	consts "nixietech/internal"
	"nixietech/utils/config"
)

type PermissionGroup struct {
	Prefix string
	Users  []string
	Clock  ClockPermissions
	Order  OrderPermissions
}

type ClockPermissions struct {
	ShowClockMenu bool
	CreateClock   bool
	DeleteClock   bool
	UpdateClock   bool
}

type OrderPermissions struct {
	ShowOrderMenu bool
	DeleteOrder   bool
}

func GetAllGroups() []*PermissionGroup {
	allGroups := config.GetConfig(consts.ConfigName).BotMessages.PermissionGroups
	typedAllGroups := make([]*PermissionGroup, 0, len(allGroups)+4)
	for _, group := range config.GetConfig(consts.ConfigName).BotMessages.PermissionGroups {
		parsedGroup := ParseGroup(group)
		typedAllGroups = append(typedAllGroups, parsedGroup)

	}
	return typedAllGroups
}

func UserGroup(update tgbotapi.Update) *PermissionGroup {
	for _, group := range GetAllGroups() {
		for _, user := range group.Users {
			if update.Message != nil {
				if user == update.Message.From.UserName {
					return group
				}
			} else {
				if user == update.CallbackQuery.From.UserName {
					return group
				}
			}

		}
	}
	return nil
}

func UserClockPerms(update tgbotapi.Update) *ClockPermissions {
	return &UserGroup(update).Clock
}

func UserOrderPerms(update tgbotapi.Update) *OrderPermissions {
	return &UserGroup(update).Order
}
