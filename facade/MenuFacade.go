package facade

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"tgBot/service"
)

func createStartMenu(update tgbotapi.Update) tgbotapi.MessageConfig {
	userID := getChatId(update)

	menu := tgbotapi.NewMessage(userID, "Выберите действие:")
	menu.ReplyMarkup = tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Показать адреса", "myAddresses"),
		),
	)
	return menu
}

func sendDefaultMenu(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	response := createStartMenu(update)
	sentMsg, _ := bot.Send(response)
	chatId := getChatId(update)

	service.SaveMessage(chatId, sentMsg.MessageID)
}

func getChatId(update tgbotapi.Update) int64 {
	if update.Message != nil {
		return update.Message.Chat.ID
	}
	return update.CallbackQuery.Message.Chat.ID
}

func HandleButtonCommand(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	text := ""
	sendDefaultMenu(bot, update)

	switch update.CallbackQuery.Data {
	case "addAddress":
		text = service.HandleAddAddress(update)
		break
	case "myAddresses":
		text = service.HandleMyAddresses(update)
		break
	}
	response := tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID, text)
	sentMsg, _ := bot.Send(response)
	service.SaveMessage(update.CallbackQuery.Message.Chat.ID, sentMsg.MessageID)

}

func DeleteAllMessages(bot *tgbotapi.BotAPI, chatID int64) {
	service.DeleteAllKnownMessagesInCurrentChat(bot, chatID)
}
