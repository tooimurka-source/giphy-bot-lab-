package main

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

	"lab3giphybot/internal/bot"
	"lab3giphybot/pkg/config"
	"lab3giphybot/pkg/logger"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		logger.Error(err)
		return
	}

	b, err := tgbotapi.NewBotAPI(cfg.BotToken)
	if err != nil {
		logger.Error(err)
		return
	}

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := b.GetUpdatesChan(u)

	logger.Info("Bot started")

	for update := range updates {
		if update.Message != nil {
			bot.HandleMessage(b, update.Message, cfg.GiphyKey)
		}
	}
}
