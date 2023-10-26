package permissions

import (
	consts "nixietech/internal"
	"nixietech/utils/config"
)

func GetAllGroups() []*PermissionGroup {
	allGroups := config.GetConfig(consts.ConfigName).BotMessages.PermissionGroups
	typedAllGroups := make([]*PermissionGroup, 0, len(allGroups)+4)
	for _, group := range config.GetConfig(consts.ConfigName).BotMessages.PermissionGroups {
		parsedGroup := ParseGroup(group)
		typedAllGroups = append(typedAllGroups, parsedGroup)

	}
	return typedAllGroups
}

func UserGroup(username string) *PermissionGroup {
	for _, group := range GetAllGroups() {
		for _, user := range group.Users {
			if user == username {
				return group
			}
		}
	}
	return nil
}

func UserClockPerms(username string) *ClockPermissions {
	return &UserGroup(username).Clock
}

func UserOrderPerms(username string) *OrderPermissions {
	return &UserGroup(username).Order
}
