package controller

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"tgBot/facade"
)

func HandleMessage(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	if update.Message == nil {
		return
	}

	facade.HandleMessage(bot, update)
}
