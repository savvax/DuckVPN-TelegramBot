package start

import (
	"bot/internal/bot/handlers"
	"bot/internal/lib/logger/sl"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log/slog"
)

func BotStart(token, providerToken string, log *slog.Logger) {
	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Error("Failed to initialize bot:", err)
	}

	bot.Debug = true

	//TODO: поменять имя бота в BotFather
	log.Info("Authorized on account", slog.String("username", bot.Self.UserName))

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, _ := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message != nil { // ignore any non-Message updates
			// TODO: antispam
			//antispam

			command := update.Message.Command()
			if command != "" {
				// command processing
				log.Info("received empty command", slog.String("command", command))
				handlers.HandleCommand(bot, update.Message, log) //
			} else {
				// message processing
				log.Info("received text message", slog.String("text", update.Message.Text))
				handlers.HandleTextMessage(bot, update.Message)
			}
		} else if update.CallbackQuery != nil {
			//antispam
			//antiSpamCheck(DuckVPN-TelegramBot, update.CallbackQuery.From.ID)

			// query processing
			log.Info("received query", slog.String("query", update.CallbackQuery.Data))
			handlers.HandleQuery(bot, providerToken, update.CallbackQuery)
		} else if update.PreCheckoutQuery != nil {
			//antispam
			//antiSpamCheck(DuckVPN-TelegramBot, update.PreCheckoutQuery.From.ID)

			// PreCheckoutQuery processing
			log.Info("received PreCheckoutQuery", slog.String("PreCheckoutQuery", update.CallbackQuery.Data))
			handlers.HandlePreCheckoutQuery(bot, update)
		} else if update.Message != nil && update.Message.SuccessfulPayment == nil {
			//обработка неудачного платежа

			// PreCheckoutQuery processing
			log.Info("potential PreCheckoutQuery error", slog.String("PreCheckoutQueryError", update.CallbackQuery.Data))
			// TODO: добавить сообщение в пул сообщений
			// определить, выносить ли на уровень ниже отправку сообщения
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Свяжитесь с поддержкой бота")

			_, err := bot.Send(msg)
			if err != nil {
				// TODO: добавить user id
				log.Info("potential PreCheckoutQuery (processing) error (message send error)", sl.Err(err))
			}
		}
	}
}
