package main

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
	"os"

	"tgBot/controller"
)

func main() {
	//Config tg bot
	botToken := os.Getenv("TG_BOT_TOKEN")
	bot, err := tgbotapi.NewBotAPI(botToken)
	if botToken == "" {
		log.Fatal("TELEGRAM_BOT_TOKEN не установлен в переменных окружения")
	}
	if err != nil {
		log.Panic(err)
	}

	// Создаем объект для получения обновлений от Telegram
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		// Проверяем, есть ли сообщение
		if update.Message != nil {
			// Логируем входящее сообщение
			log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

			controller.HandleMessage(bot, update)
		}
	}

}

// Получать текущую цену ETH - DONE
// Выводить цену опользаку по запросу - DONE
// По адресу получать цену кошелька и выводить пользаку - DONE
