package handlers

import (
	"bot/internal/bot/keyboard"
	"bot/internal/bot/messages"
	"bot/internal/lib/logger/sl"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log/slog"
	"strconv"
)

func HandleTextMessage(log *slog.Logger, bot *tgbotapi.BotAPI, message *tgbotapi.Message) {
	//TODO обработка сообщений

	if message.SuccessfulPayment != nil {
		//TODO обработка платежа, добавление платежа в бд, выдача покупателю ключа
		//обработка успешного платежа
		//edit := tgbotapi.NewDeleteMessage(message.Chat.ID, message.MessageID)
		if message.SuccessfulPayment.InvoicePayload == "paid_VPN_NL_120_days" {

			//purchaedProduct := message.SuccessfulPayment.InvoicePayload
			//
			//base64string, key, location, encryption, password, address, port, err := database.PaidKeyProcess(message.Chat.ID, purchaedProduct)
			//if err != nil {
			//	fmt.Println(err)
			//}
			//fmt.Println(location)
			//fmt.Println(key)
			//
			//keyRefrashed := database.KeyReplaceOutlineToLocation(key, location)
			//
			////TODO перехват purchaedProduct, объединение с пользователем, генерация ключа, возврат ключа пользователю
			//
			//url := fmt.Sprintf("https://s3.amazonaws.com/outline-vpn/invite.html#ss%s%s%s%s%s%s%s%s", "%3A%2F%2F", base64string, "%40", address, "%3A", port, "%23", location)
			//
			//str := fmt.Sprintf("Спасибо за заказ, для настройки VPN перейдите по <a href=\"%s\">ССЫЛКЕ</a>\n\nДанные для ручной настройки клиента\naddress: %s\nport: %s\nencryption: %s\npassword: %s\n\nили скопируйте конфигурационный ключ в свой клиент:",
			//	url, address, port, encryption, password)
			//fmt.Println("-------")
			//fmt.Println(url)
			//fmt.Println("-------")
			//msg := tgbotapi.NewMessage(message.Chat.ID, fmt.Sprint(str))
			//msg.ParseMode = "HTML"
			//msg.DisableWebPagePreview = true
			////TODO регулярное выражение по изменению ключа, отброс "outline=1" и добавление локации сервера
			//msg2 := tgbotapi.NewMessage(message.Chat.ID, fmt.Sprint(keyRefrashed))
			//msg3 := tgbotapi.NewMessage(message.Chat.ID, "Вернуться в главное меню")
			//msg3.ReplyMarkup = &keyboard.ReturnToMain
			//
			//DuckVPN-TelegramBot.Send(msg)
			//DuckVPN-TelegramBot.Send(msg2)
			//DuckVPN-TelegramBot.Send(msg3)
			//TODO добавить клавиатуру для перехода в гравное меню
		} else if message.SuccessfulPayment.InvoicePayload == "paid_VPN_NL_30_days" {

			//purchaedProduct := message.SuccessfulPayment.InvoicePayload

			//base64string, key, location, encryption, password, address, port, err := database.PaidKeyProcess(message.Chat.ID, purchaedProduct)
			//if err != nil {
			//	fmt.Println(err)
			//}
			//fmt.Println(location)
			//fmt.Println(key)
			//
			//keyRefrashed := database.KeyReplaceOutlineToLocation(key, location)
			//
			////TODO перехват purchaedProduct, объединение с пользователем, генерация ключа, возврат ключа пользователю
			//fmt.Printf("\n\n\n\n\n\n\n\n\n\n\n_--------%s-------\n\n\n\n\n\n\n\n\n", purchaedProduct)
			//
			//url := fmt.Sprintf("https://s3.amazonaws.com/outline-vpn/invite.html#ss%s%s%s%s%s%s%s%s", "%3A%2F%2F", base64string, "%40", address, "%3A", port, "%2F%3F", location)
			//
			//str := fmt.Sprintf("Спасибо за заказ, для настройки VPN перейдите по <a href=\"%s\">ССЫЛКЕ</a>\n\n%s\n\nДанные для ручной настройки клиента\naddress: %s\nport: %s\nencryption: %s\npassword: %s\n\nили скопируйте конфигурационный ключ в свой клиент:",
			//	url, url, address, port, encryption, password)
			//fmt.Println("-------")
			//fmt.Println(url)
			//fmt.Println("-------")
			//msg := tgbotapi.NewMessage(message.Chat.ID, fmt.Sprint(str))
			//msg.ParseMode = "HTML"
			//msg.DisableWebPagePreview = true
			////TODO регулярное выражение по изменению ключа, отброс "outline=1" и добавление локации сервера
			//msg2 := tgbotapi.NewMessage(message.Chat.ID, fmt.Sprint(keyRefrashed))
			//msg3 := tgbotapi.NewMessage(message.Chat.ID, "Вернуться в главное меню")
			//msg3.ReplyMarkup = &keyboard.ReturnToMain
			//
			//DuckVPN-TelegramBot.Send(msg)
			//DuckVPN-TelegramBot.Send(msg2)
			//DuckVPN-TelegramBot.Send(msg3)
			////TODO добавить клавиатуру для перехода в гравное меню
		} else if message.SuccessfulPayment.InvoicePayload == "subscription_30_days" {
			//
			//
			//
			//!!!!!!!!!
			//

			//purchaedProduct := message.SuccessfulPayment.InvoicePayload

			//base64string, key, location, encryption, password, address, port, err := database.PaidKeyProcess(message.Chat.ID, purchaedProduct)
			//if err != nil {
			//	fmt.Println(err)
			//}
			//fmt.Println(location)
			//fmt.Println(key)
			//
			//keyRefrashed := database.KeyReplaceOutlineToLocation(key, location)

			//TODO перехват purchaedProduct, объединение с пользователем, генерация ключа, возврат ключа пользователю
			//fmt.Printf("\n\n\n\n\n\n\n\n\n\n\n_--------%s-------\n\n\n\n\n\n\n\n\n", purchaedProduct)
			//
			//url := fmt.Sprintf("https://s3.amazonaws.com/outline-vpn/invite.html#ss%s%s%s%s%s%s%s%s", "%3A%2F%2F", base64string, "%40", address, "%3A", port, "%2F%3F", location)
			//
			//str := fmt.Sprintf("Спасибо за заказ, для настройки VPN перейдите по <a href=\"%s\">ССЫЛКЕ</a>\n\n%s\n\nДанные для ручной настройки клиента\naddress: %s\nport: %s\nencryption: %s\npassword: %s\n\nили скопируйте конфигурационный ключ в свой клиент:",
			//	url, url, address, port, encryption, password)
			//fmt.Println("-------")
			//fmt.Println(url)
			//fmt.Println("-------")
			//msg := tgbotapi.NewMessage(message.Chat.ID, fmt.Sprint(str))
			//msg.ParseMode = "HTML"
			//msg.DisableWebPagePreview = true
			////TODO регулярное выражение по изменению ключа, отброс "outline=1" и добавление локации сервера
			//msg2 := tgbotapi.NewMessage(message.Chat.ID, fmt.Sprint(keyRefrashed))
			//msg3 := tgbotapi.NewMessage(message.Chat.ID, "Вернуться в главное меню")
			//msg3.ReplyMarkup = &keyboard.ReturnToMain
			//
			//DuckVPN-TelegramBot.Send(msg)
			//DuckVPN-TelegramBot.Send(msg2)
			//DuckVPN-TelegramBot.Send(msg3)
			////TODO добавить клавиатуру для перехода в гравное меню
		} else {
			fmt.Println("ERROR")
		}
	} else if state, exists := userStates[message.Chat.ID]; exists {
		processNonCommandMessage(log, bot, message, state)
	} else if message.Text != "" {
		if message.Text == messages.Msgs["get_free_VPN_key_NL"] {
			//base64string, key, location, encryption, password, address, port, err := database.FreeKeyProcess(message.Chat.ID, message.Text)
			//if err != nil {
			//	fmt.Println(err)
			//}
			//fmt.Println(location)
			//fmt.Println(key)
			//
			//keyRefrashed := database.KeyReplaceOutlineToLocation(key, location)

			//url := fmt.Sprintf("https://s3.amazonaws.com/outline-vpn/invite.html#ss%s%s%s%s%s%s%s%s", "%3A%2F%2F", base64string, "%40", address, "%3A", port, "%2F%3F", location)
			//
			//str := fmt.Sprintf("Спасибо за заказ, для настройки VPN перейдите по <a href=\"%s\">ССЫЛКЕ</a>\n\n%s\n\nДанные для ручной настройки клиента\naddress: %s\nport: %s\nencryption: %s\npassword: %s\n\nили скопируйте конфигурационный ключ в свой клиент:",
			//	url, url, address, port, encryption, password)

			//msg := tgbotapi.NewMessage(message.Chat.ID, fmt.Sprint(str))
			//msg.ParseMode = "HTML"
			//msg.DisableWebPagePreview = true
			//TODO регулярное выражение по изменению ключа, отброс "outline=1" и добавление локации сервера
			//msg2 := tgbotapi.NewMessage(message.Chat.ID, fmt.Sprint(keyRefrashed))
			msg := tgbotapi.NewMessage(message.Chat.ID, "Вернуться в главное меню")
			msg.ReplyMarkup = &keyboard.ReturnToMain

			//DuckVPN-TelegramBot.Send(msg)
			//DuckVPN-TelegramBot.Send(msg2)
			_, err := bot.Send(msg)
			if err != nil {
				// TODO: добавить user id
				// TODO: изменить при именеии ключа в словаре с сообщениями
				log.Info("Text get_free_VPN_key_NL error", sl.Err(err))
			}
		}

	}
}

func processNonCommandMessage(log *slog.Logger, bot *tgbotapi.BotAPI, message *tgbotapi.Message, state *UserState) {
	state, exists := userStates[message.Chat.ID]
	if !exists {
		return
	}

	switch state.Step {
	case 1:
		state.ApiURL = message.Text
		msg := tgbotapi.NewMessage(message.Chat.ID, "Enter max keys:")
		_, err := bot.Send(msg)
		if err != nil {
			// TODO: добавить user id
			log.Info("Text processNonCommandMessage state 1 error", sl.Err(err))
		}
		state.Step++
	case 2:
		// TODO: Add proper integer parsing and validation
		// TODO wrap error
		state.MaxKeys, _ = strconv.Atoi(message.Text)
		msg := tgbotapi.NewMessage(message.Chat.ID, "Enter region code:")
		_, err := bot.Send(msg)
		if err != nil {
			// TODO: добавить user id
			log.Info("Text processNonCommandMessage state 2 error", sl.Err(err))
		}
		state.Step++
	case 3:
		state.RegionCode = message.Text
		msg := tgbotapi.NewMessage(message.Chat.ID, "Enter server type:")
		_, err := bot.Send(msg)
		if err != nil {
			// TODO: добавить user id
			log.Info("Text processNonCommandMessage state 3 error", sl.Err(err))
		}
		state.Step++
	case 4:
		state.ServerType = message.Text
		msg := tgbotapi.NewMessage(message.Chat.ID, "Enter blocked status (true/false):")
		_, err := bot.Send(msg)
		if err != nil {
			// TODO: добавить user id
			log.Info("Text processNonCommandMessage state 4 error", sl.Err(err))
		}
		state.Step++
	case 5:
		// TODO: Add proper boolean parsing and validation
		// TODO wrap error
		state.Blocked, _ = strconv.ParseBool(message.Text)
		state.Step++

		//TODO удалить print
		fmt.Println(state)

		//TODO проверка типа данных
		//TODO изменить назваине таблицы
		//_, err := database.DB.Exec("INSERT INTO server_info (api_url, max_keys, region, server_type, blocked) VALUES ($1, $2, $3, $4, $5)",
		//	state.ApiURL, state.MaxKeys, state.RegionCode, state.ServerType, state.Blocked)

		//if err != nil {
		//	fmt.Println(err)
		//	log.Panic(err)
		//}

		delete(userStates, message.Chat.ID)

		//TODO добавить клавиатуру

		//TODO добавить проверку URL
		//приведеине региона к заклавным буквам
		//приведение типа сервера к прописным и проверка (free, paid, mix)
		//вывод всех показателей для утверждения и только после этого добавление в базу данных
	}
}
