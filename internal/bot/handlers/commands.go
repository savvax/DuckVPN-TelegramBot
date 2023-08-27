package handlers

import (
	"bot/internal/bot/keyboard"
	"bot/internal/bot/messages"
	"bot/internal/client/client"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
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

func HandleCommand(bot *tgbotapi.BotAPI, message *tgbotapi.Message) {
	switch message.Command() {
	case "start":
		c := client.NewClient("localhost", 4001)
		code, err := c.CreateUser(message.Chat.ID)
		if err != nil {
			panic(err)
		}

		if code == 201 {
			msg := tgbotapi.NewMessage(message.Chat.ID, messages.Msgs["main"])
			msg.ReplyMarkup = &keyboard.MainKeyboard
			bot.Send(msg)
		} else {
			//TODO записать ошибку, ничего не возвращать пользователю(?)
		}
		//InitUser(int64(message.From.ID))

	case "help":
		//TODO добавить текст описывеающий команду /help
		msg := tgbotapi.NewMessage(message.Chat.ID, "HELP command, такой команды пока нет")
		bot.Send(msg)
	case "terms":
		//TODO проверка на нахождение пользоватлея в таблице admin
		msg := tgbotapi.NewMessage(message.Chat.ID, messages.Msgs["terms"])
		msg.ReplyMarkup = &keyboard.ReturnToMain
		msg.ParseMode = "HTML"
		bot.Send(msg)
	case "add":
		//TODO проверка на нахождение пользоватлея в таблице admin
		userStates[message.Chat.ID] = &UserState{Step: 1}
		msg := tgbotapi.NewMessage(message.Chat.ID, "Enter API URL:")
		bot.Send(msg)
	default:
		msg := tgbotapi.NewMessage(message.Chat.ID, "unknown command")
		bot.Send(msg)
	}
}
