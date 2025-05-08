package facade

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"tgBot/service"
)

func handleCommand(update tgbotapi.Update) tgbotapi.MessageConfig {

	command := update.Message.Command()
	text := ""

	switch command {
	case "addAddress":
		text = service.HandleAddAddress(update)
		break
	case "myAddresses":
		text = service.HandleMyAddresses(update)
		break
	case "help":
		text = ""
		break
	default:
		text = "Неизвестная команда. Используй команду /help чтобы узнать список доступных команд."
	}

	return tgbotapi.NewMessage(update.Message.Chat.ID, text)

}
