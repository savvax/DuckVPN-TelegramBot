package handlers

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func HandlePreCheckoutQuery(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	pca := tgbotapi.PreCheckoutConfig{
		OK:                 true,
		PreCheckoutQueryID: update.PreCheckoutQuery.ID,
	}
	_, err := bot.AnswerPreCheckoutQuery(pca)
	if err != nil {
		fmt.Println(err)
	}
	//TODO handle Error
}
