package permissions

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	consts "nixietech/internal"
	"nixietech/internal/fetcher"
	"nixietech/utils/config"
	"strings"
)

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

func ParseTrueFalse(message string) bool {
	if message == "✅" {
		return true
	}
	return false
}

func ParseFullPerms(username string) fetcher.HashTags {
	return fetcher.NewTag("$permission-prefix-full$", fetcher.ParseHashTags(
		config.GetConfig(consts.ConfigName).BotMessages.PermissionPrefixFull,
		[]fetcher.HashTags{
			fetcher.NewTag("$ROLE$", UserGroup(username).Prefix),
		},
	))
}

func ParseMinimumPerms(from *tgbotapi.User) fetcher.HashTags {
	return fetcher.NewTag("$permission-prefix-minimum$", fetcher.ParseHashTags(
		config.GetConfig(consts.ConfigName).BotMessages.PermissionPrefixMinimum,
		[]fetcher.HashTags{
			ParseGroupPrefix(from.UserName),
			fetcher.ParseFullName(from.FirstName, from.LastName),
		},
	))
}

func ParseGroupPrefix(username string) fetcher.HashTags {
	return fetcher.NewTag("$ROLE$", UserGroup(username).Prefix)
}
