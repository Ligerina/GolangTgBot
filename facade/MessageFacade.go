package facade

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"tgBot/service"
)

func HandleMessage(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	service.SaveMessage(update.Message.Chat.ID, update.Message.MessageID)
	if update.Message.IsCommand() {
		sendDefaultMenu(bot, update)
		response := handleCommand(update)
		sentMsg, _ := bot.Send(response)
		service.SaveMessage(update.Message.Chat.ID, sentMsg.MessageID)
	}
}
