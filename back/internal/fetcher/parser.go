package fetcher

import (
	"fmt"
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

func ParseHashTags(message string, items []HashTags) string {
	for _, item := range items {
		message = strings.Replace(message, item.message, item.item, -1)
	}
	return message
}

func ParseFullName(firstname, lastname string) HashTags {
	return NewTag("$NAME$", fmt.Sprintf("%s %s", firstname, lastname))
}

func ParsePrefix(prefix string) HashTags {
	return NewTag("$PREFIX$", prefix)
}
