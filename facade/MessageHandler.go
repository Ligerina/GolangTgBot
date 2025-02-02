package facade

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

func HandleCommand(bot *tgbotapi.BotAPI, update tgbotapi.Update) {

	switch update.Message.Text {
	case "/start":
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Hello, World!")
		bot.Send(msg)
	default:
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Unknown")
		bot.Send(msg)
	}

}
