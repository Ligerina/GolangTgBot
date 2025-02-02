package main

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
	"os"
)

func main() {
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

			// Обрабатываем команду /start
			if update.Message.Text == "/start" {
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Hello, World!")
				bot.Send(msg)
			}
		}
	}

}
