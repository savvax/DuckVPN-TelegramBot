package handlers

import (
	"bot/internal/bot/keyboard"
	"bot/internal/bot/messages"
	"bot/internal/lib/logger/sl"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log/slog"
)

type UserState struct {
	ApiURL     string
	MaxKeys    int
	RegionCode string
	ServerType string
	Blocked    bool
	Step       int
}

var userStates = make(map[int64]*UserState)

func HandleCommand(bot *tgbotapi.BotAPI, message *tgbotapi.Message, log *slog.Logger) {
	switch message.Command() {
	case "start":
		msg := tgbotapi.NewMessage(message.Chat.ID, messages.Msgs["main"])
		msg.ReplyMarkup = &keyboard.MainKeyboard
		_, err := bot.Send(msg)
		if err != nil {
			// TODO: добавить user id
			log.Info("Command /start error", sl.Err(err))
		}

		// Проверка находжения пользователя в бд, в случае отсутствия -> Инициализация нового пользователя
		// если пользователь есть в бд - возврат в главное меню.

		//// TODO: перенос клиетна в main
		//c := client.NewClient("localhost", 4001)
		//code, err := c.CreateUser(message.Chat.ID)
		//if err != nil {
		//	panic(err)
		//}
		//
		//if code == 201 {
		//	msg := tgbotapi.NewMessage(message.Chat.ID, messages.Msgs["main"])
		//	msg.ReplyMarkup = &keyboard.MainKeyboard
		//	bot.Send(msg)
		//} else {
		//	//TODO записать ошибку, ничего не возвращать пользователю(?)
		//}
		////InitUser(int64(message.From.ID))

	case "help":
		//TODO добавить текст описывеающий команду /help
		msg := tgbotapi.NewMessage(message.Chat.ID, "HELP command, такой команды пока нет")
		_, err := bot.Send(msg)
		if err != nil {
			// TODO: добавить user id
			log.Info("Command /help error", sl.Err(err))
		}
	case "terms":
		//TODO проверка на нахождение пользоватлея в таблице admin
		msg := tgbotapi.NewMessage(message.Chat.ID, messages.Msgs["terms"])
		msg.ReplyMarkup = &keyboard.ReturnToMain
		msg.ParseMode = "HTML"
		_, err := bot.Send(msg)
		if err != nil {
			// TODO: добавить user id
			log.Info("Command /terms error", sl.Err(err))
		}
	case "add":
		//TODO проверка на нахождение пользоватлея в таблице admin
		userStates[message.Chat.ID] = &UserState{Step: 1}
		msg := tgbotapi.NewMessage(message.Chat.ID, "Enter API URL:")
		_, err := bot.Send(msg)
		if err != nil {
			// TODO: добавить user id
			log.Info("Command /add error", sl.Err(err))
		}
	default:
		msg := tgbotapi.NewMessage(message.Chat.ID, "unknown command")
		_, err := bot.Send(msg)
		if err != nil {
			// TODO: добавить user id
			log.Info("Default command error", sl.Err(err))
		}
	}
}
