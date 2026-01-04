package bot

import (
	"strings"

	"lab3giphybot/internal/api/giphy"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func HandleMessage(bot *tgbotapi.BotAPI, msg *tgbotapi.Message, giphyKey string) {
	if !msg.IsCommand() {
		return
	}

	switch msg.Command() {
	case "start":
		bot.Send(tgbotapi.NewMessage(msg.Chat.ID,
			"Привіт!👋 Напиши команду:\n/gif <слово> і тобі знайде гіфку!"))

	case "gif":
		query := strings.TrimSpace(msg.CommandArguments())
		if query == "" {
			bot.Send(tgbotapi.NewMessage(msg.Chat.ID,
				"❗ Вкажи ключове слово, наприклад:\n/gif cat"))
			return
		}

		url, err := giphy.GetGif(giphyKey, query)
		if err != nil {
			bot.Send(tgbotapi.NewMessage(msg.Chat.ID,
				"Нічого не знайдено. ❌"))
			return
		}

		bot.Send(tgbotapi.NewAnimation(msg.Chat.ID,
			tgbotapi.FileURL(url)))
	}
}
