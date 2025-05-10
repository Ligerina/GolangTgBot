package service

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
)

var (
	userMessages = make(map[int64][]int) // Сохраняем сообщения для каждого чата
)

func SaveMessage(chatID int64, messageID int) {
	userMessages[chatID] = append(userMessages[chatID], messageID)
}

func DeleteAllKnownMessagesInCurrentChat(bot *tgbotapi.BotAPI, chatID int64) {
	messageIDs, exists := userMessages[chatID]
	if !exists {
		return
	}
	for _, messageID := range messageIDs {
		deleteConfig := tgbotapi.NewDeleteMessage(chatID, messageID)
		_, err := bot.Request(deleteConfig)
		if err != nil {
			log.Printf("Ошибка при удалении сообщения %d: %v", messageID, err)
		}
	}
	userMessages[chatID] = []int{}
}
