package service

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"strings"
)

// userAddresses — глобальная map: userID ➔ список адресов
var userAddresses = make(map[int64][]string)

// HandleAddAddress обрабатывает команду /addaddress
func HandleAddAddress(update tgbotapi.Update) string {
	userID := update.Message.Chat.ID
	args := strings.TrimSpace(update.Message.CommandArguments())

	if args == "" {
		msg := "❗ Ты не указал адрес! Пример: `/addaddress 0x123...`"
		return msg
	}

	if !isValidAddress(args) {
		msg := "❗ Неверный формат адреса! Адрес должен начинаться с `0x` и содержать 42 символа."
		return msg
	}

	userAddresses[userID] = append(userAddresses[userID], args)

	msg := fmt.Sprintf("✅ Адрес %s успешно добавлен!", args)
	return msg
}

// handleMyAddresses отвечает на команду /myaddresses
func HandleMyAddresses(update tgbotapi.Update) string {
	userID := update.Message.Chat.ID

	addresses := userAddresses[userID]
	if len(addresses) == 0 {
		msg := "У тебя пока нет сохранённых адресов."
		return msg
	}

	text := "Твои адреса:\n"
	for i, addr := range addresses {
		text += fmt.Sprintf("%d. %s\n", i+1, addr)
	}
	return text
}

// isValidAddress — проверка валидности Ethereum-адреса
func isValidAddress(addr string) bool {
	return strings.HasPrefix(addr, "0x") && len(addr) == 42
}
