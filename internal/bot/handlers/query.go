package handlers

import (
	"bot/internal/bot/keyboard"
	"bot/internal/bot/messages"
	"bot/internal/lib/logger/sl"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log/slog"
)

func HandleQuery(bot *tgbotapi.BotAPI, providerToken string, query *tgbotapi.CallbackQuery, log *slog.Logger) {
	switch query.Data {
	case "main":
		msg := tgbotapi.NewEditMessageText(query.Message.Chat.ID, query.Message.MessageID, "Главное меню (Main)")
		msg.ReplyMarkup = &keyboard.MainKeyboard
		_, err := bot.Send(msg)
		if err != nil {
			// TODO: добавить user id
			log.Info("Query main error", sl.Err(err))
		}
	case "VPN":
		msg := tgbotapi.NewEditMessageText(query.Message.Chat.ID, query.Message.MessageID, messages.Msgs["VPN"]) //TODO ищменить текст
		//edit.ParseMode = "MarkdownV2"
		msg.ReplyMarkup = &keyboard.VpnKeyboard
		_, err := bot.Send(msg)
		if err != nil {
			// TODO: добавить user id
			log.Info("Query VPN error", sl.Err(err))
		}
	case "free_VPN":
		msg := tgbotapi.NewEditMessageText(query.Message.Chat.ID, query.Message.MessageID, messages.Msgs["free_VPN"])
		msg.ParseMode = "HTML"
		msg.ReplyMarkup = &keyboard.FreeVPNKeyboard
		_, err := bot.Send(msg)
		if err != nil {
			// TODO: добавить user id
			log.Info("Query free_VPN error", sl.Err(err))
		}
	case "get_free_VPN_key_NL":
		//TODO Проверка в бд и выдача ключа
		//message := tgbotapi.NewMessage(query.Message.Chat.ID, "core.Msgs[\"get_free_VPN_key_NL\"]")
		//fmt.Println(messages.Msgs["get_free_VPN_key_NL"])
		//base64string, key, location, encryption, password, address, port, err := database.FreeKeyProcess(message.ChatID, messages.Msgs["get_free_VPN_key_NL"])
		//if err != nil {
		//	fmt.Println(err)
		//}
		//fmt.Println(location)
		//fmt.Println(key)
		//
		//keyRefrashed := database.KeyReplaceOutlineToLocation(key, location)
		//
		//url := fmt.Sprintf("https://s3.amazonaws.com/outline-vpn/invite.html#ss%s%s%s%s%s%s%s%s", "%3A%2F%2F", base64string, "%40", address, "%3A", port, "%2F%3F", location)
		//
		//str := fmt.Sprintf("Спасибо за заказ, для настройки VPN перейдите по <a href=\"%s\">ССЫЛКЕ</a>\n\n%s\n\nДанные для ручной настройки клиента\naddress: %s\nport: %s\nencryption: %s\npassword: %s\n\nили скопируйте конфигурационный ключ в свой клиент:",
		//	url, url, address, port, encryption, password)
		//
		//msg := tgbotapi.NewMessage(message.ChatID, fmt.Sprint(str))
		//msg.ParseMode = "HTML"
		//msg.DisableWebPagePreview = true
		////TODO регулярное выражение по изменению ключа, отброс "outline=1" и добавление локации сервера
		//msg2 := tgbotapi.NewMessage(message.ChatID, fmt.Sprint(keyRefrashed))
		//msg3 := tgbotapi.NewMessage(message.ChatID, "Вернуться в главное меню")
		//msg3.ReplyMarkup = &keyboard.ReturnToMain
		//
		//DuckVPN-TelegramBot.Send(msg)
		//DuckVPN-TelegramBot.Send(msg2)
		//DuckVPN-TelegramBot.Send(msg3)

	case "subscription_30":
		deleteMsg := tgbotapi.NewDeleteMessage(query.Message.Chat.ID, query.Message.MessageID)
		invoice := tgbotapi.NewInvoice(int64(query.From.ID), "Subscription", "Купить подписку на 30 дней", "subscription_30_days",
			providerToken, "StartParameterSubscription", "RUB", &[]tgbotapi.LabeledPrice{{"Label", 20000}})
		bot.Send(invoice)
		bot.Send(deleteMsg)
		msg := tgbotapi.NewMessage(query.Message.Chat.ID, "Чтобы вернуться назад, нажмете:")
		msg.ReplyMarkup = &keyboard.ReturnPaidKeyVPNKeyboard
		_, err := bot.Send(msg)
		if err != nil {
			// TODO: добавить user id
			log.Info("Query subscription_30 error", sl.Err(err))
		}

	case "paid_VPN":
		msg := tgbotapi.NewEditMessageText(query.Message.Chat.ID, query.Message.MessageID, messages.Msgs["paid_VPN"])
		msg.ReplyMarkup = &keyboard.PaidVPNKeyboard
		_, err := bot.Send(msg)
		if err != nil {
			// TODO: добавить user id
			log.Info("Query paid_VPN error", sl.Err(err))
		}
	case "get_paid_VPN_key_NL":
		//TODO добавить проверку наличия возможнрости выдать платный ключ
		msg := tgbotapi.NewEditMessageText(query.Message.Chat.ID, query.Message.MessageID, messages.Msgs["get_paid_VPN_key_NL"])
		msg.ReplyMarkup = &keyboard.PaidVPNNLKeyboard
		_, err := bot.Send(msg)
		if err != nil {
			// TODO: добавить user id
			log.Info("Query get_paid_VPN_key_NL error", sl.Err(err))
		}

	case "get_paid_VPN_key_NL_30":
		deleteMsg := tgbotapi.NewDeleteMessage(query.Message.Chat.ID, query.Message.MessageID)
		invoice := tgbotapi.NewInvoice(int64(query.From.ID), "VPN (NL)", "VPN на 30 дней c сервером в Нидерландах", "paid_VPN_NL_30_days",
			providerToken, "StartParameter", "RUB", &[]tgbotapi.LabeledPrice{{"Label", 20000}})
		bot.Send(invoice)
		bot.Send(deleteMsg)
		msg := tgbotapi.NewMessage(query.Message.Chat.ID, "Чтобы вернуться назад, нажмете:")
		msg.ReplyMarkup = &keyboard.ReturnPaidKeyVPNKeyboard
		_, err := bot.Send(msg)
		if err != nil {
			// TODO: добавить user id
			log.Info("Query get_paid_VPN_key_NL_30 error", sl.Err(err))
		}

	case "get_paid_VPN_key_NL_120":
		//TODO - резервация ключа
		deleteMsg := tgbotapi.NewDeleteMessage(query.Message.Chat.ID, query.Message.MessageID)
		invoice := tgbotapi.NewInvoice(int64(query.From.ID), "VPN (NL) 120", "VPN на 180 дней c сервером в Нидерландах", "paid_VPN_NL_120_days",
			providerToken, "StartParameter2", "RUB", &[]tgbotapi.LabeledPrice{{"Label2", 80000}})
		bot.Send(invoice)
		bot.Send(deleteMsg)
		msg := tgbotapi.NewMessage(query.Message.Chat.ID, "Чтобы вернуться назад, нажмете:")
		msg.ReplyMarkup = &keyboard.ReturnPaidKeyVPNKeyboard
		_, err := bot.Send(msg)
		if err != nil {
			// TODO: добавить user id
			log.Info("Query get_paid_VPN_key_NL_120 error", sl.Err(err))
		}

	//case "get_paid_VPN_key_RU":
	//	//TODO добавить проверку наличия возможнрости выдать платный ключ
	//	deleteMsg := tgbotapi.NewDeleteMessage(query.Message.Chat.ID, query.Message.MessageID)
	//	invoice := tgbotapi.NewInvoice(int64(query.From.ID), "VPN", "VPN на 6 месяцев Россия (RU)", "paid_VPN_RU_1_month",
	//		os.Getenv("PROVIDER_TOKEN"), "StartParameter", "RUB", &[]tgbotapi.LabeledPrice{{"Label", 100000}})
	//	DuckVPN-TelegramBot.Send(invoice)
	//	DuckVPN-TelegramBot.Send(deleteMsg)
	//	msg := tgbotapi.NewMessage(query.Message.Chat.ID, "Чтобы вернуться назад, нажмете:")
	//	msg.ReplyMarkup = &ReturnPaidKeyVPNKeyboard
	//	DuckVPN-TelegramBot.Send(msg)
	case "apps":
		msg := tgbotapi.NewEditMessageText(query.Message.Chat.ID, query.Message.MessageID, "скачать приложение")
		msg.ReplyMarkup = &keyboard.AppsKeyboard
		_, err := bot.Send(msg)
		if err != nil {
			// TODO: добавить user id
			log.Info("Query apps error", sl.Err(err))
		}
	case "my_keys":
		msg := tgbotapi.NewEditMessageText(query.Message.Chat.ID, query.Message.MessageID, "РАЗДЕЛ \"МОИ КЛЮЧИ\" ")
		msg.ReplyMarkup = &keyboard.MyKeysKeyboard
		_, err := bot.Send(msg)
		if err != nil {
			// TODO: добавить user id
			log.Info("Query my_keys error", sl.Err(err))
		}
	case "my_paid_keys":
		msg := tgbotapi.NewEditMessageText(query.Message.Chat.ID, query.Message.MessageID, "Ваши платные ключ: ")
		msg.ReplyMarkup = &keyboard.MyPaidKeysKeyboard
		_, err := bot.Send(msg)
		if err != nil {
			// TODO: добавить user id
			log.Info("Query my_paid_keys error", sl.Err(err))
		}
	case "my_free_key":
		//TODO запросы к базе данных
	}
}
