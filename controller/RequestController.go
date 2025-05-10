package controller

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"tgBot/facade"
)

func HandleMessage(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	facade.DeleteAllMessages(bot, update.Message.Chat.ID)
	facade.HandleMessage(bot, update)
}

func HandleButtonCommand(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	facade.DeleteAllMessages(bot, update.CallbackQuery.Message.Chat.ID)
	facade.HandleButtonCommand(bot, update)
}
