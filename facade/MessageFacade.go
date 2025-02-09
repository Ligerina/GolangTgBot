package facade

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"strconv"
	"strings"
	"tgBot/service/Api"
)

func HandleCommand(bot *tgbotapi.BotAPI, update tgbotapi.Update) {

	words := strings.Split(update.Message.Text, " ")
	command := words[0]

	switch command {
	case "/start":
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Hello, World!")
		bot.Send(msg)
	case "/getETHPrice":
		response, _ := Api.GetETHPrice()
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "ETH Price = "+strconv.FormatFloat(response, 'f', 2, 64))
		bot.Send(msg)
	case "/getAdressPrice":
		var resonse, _ = Api.GetWalletBalance(words[1])
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Your account price in ETH = "+strconv.FormatFloat(resonse, 'f', 6, 64)+" ETH")
		bot.Send(msg)
	default:
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Unknown")
		bot.Send(msg)
	}

}
