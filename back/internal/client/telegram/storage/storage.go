package storage

import (
	"time"
)

type TGStorageItem struct {
	Username string
	Type     string
	Message  string
	Date     time.Time
}

type TGStorage []TGStorageItem

type TGStorageManager struct {
	storage TGStorage
}

func New() *TGStorageManager {
	return &TGStorageManager{storage: make(TGStorage, 0, 10)}
}

func (manager *TGStorageManager) AddItem(
	username string,
	typeMessage string,
	message string,
) bool {
	for _, item := range manager.storage {
		if item.Username == username {
			if item.Type != typeMessage && item.Message != message {
				item.Message = message
				item.Type = typeMessage
				item.Date = time.Now()
				return true
			}
			return false
		}
	}

	manager.storage = append(manager.storage, TGStorageItem{
		Username: username,
		Type:     typeMessage,
		Message:  message,
		Date:     time.Now(),
	})
	return true
}

func (manager *TGStorageManager) RemoveItem(username string) bool {
	currentIndex := -1
	for index, item := range manager.storage {
		if item.Username == username {
			currentIndex = index
		}
	}
	if currentIndex == -1 {
		return false
	}

	manager.storage = append(manager.storage[:currentIndex], manager.storage[currentIndex+1:]...)
	return true
}

func (manager *TGStorageManager) Item(username string) *TGStorageItem {
	for _, item := range manager.storage {
		if item.Username == username {
			return &item
		}
	}
	return nil
}
