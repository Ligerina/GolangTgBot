package facade

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func HandleMessage(bot *tgbotapi.BotAPI, update tgbotapi.Update) {

	if update.Message.IsCommand() {
		response := handleCommand(update)
		bot.Send(response)
	} else {
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Напиши команду")
		bot.Send(msg)
	}

}
