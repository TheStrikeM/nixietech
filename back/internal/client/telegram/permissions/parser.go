package permissions

import "strings"

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
		clock: ClockPermissions{
			ShowClockMenu: ParseTrueFalse(splitedMsg[1]),
			CreateClock:   ParseTrueFalse(splitedMsg[2]),
			DeleteClock:   ParseTrueFalse(splitedMsg[3]),
			UpdateClock:   ParseTrueFalse(splitedMsg[4]),
		},
		order: OrderPermissions{
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
