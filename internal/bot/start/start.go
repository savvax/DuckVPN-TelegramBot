package start

import (
	"bot/internal/bot/handlers"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
)

func StartBot(token, providerToken string) {
	//TODO transfer the token to the configuration file
	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	//TODO: поменять имя бота в BotFather
	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, _ := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message != nil { // ignore any non-Message updates
			// TODO: antispam
			//antispam

			command := update.Message.Command()
			if command != "" {
				log.Printf("Received command: %s", command)
				handlers.HandleCommand(bot, update.Message)
			} else {
				log.Printf("Received text message: %s", update.Message.Text)
				// Обрабатываем обычное сообщение
				handlers.HandleTextMessage(bot, update.Message)
			}
		} else if update.CallbackQuery != nil {

			//antispam
			//antiSpamCheck(DuckVPN-TelegramBot, update.CallbackQuery.From.ID)

			handlers.HandleQuery(bot, providerToken, update.CallbackQuery)
		} else if update.PreCheckoutQuery != nil {
			//antispam
			//antiSpamCheck(DuckVPN-TelegramBot, update.PreCheckoutQuery.From.ID)

			handlers.HandlePreCheckoutQuery(bot, update)
		} else if update.Message != nil && update.Message.SuccessfulPayment == nil {
			//обработка неудачного платежа
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Payment unsuccessful. Please try again.")
			bot.Send(msg)
		}
	}
}
