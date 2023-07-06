package tgmanager

import (
	"fmt"
	"strings"

	tg "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/kutoru/fibogachabot/pkg/glb"
)

type USER_MESSAGE struct {
	user int64
	msg  string
}

func messageHandler(update tg.Update) {
	if strings.HasPrefix(update.Message.Text, "/") {
		commandHandler(update)
	} else {
		fmt.Println("Not a command")
		actionHandler(update)
	}
}

func commandHandler(update tg.Update) {
	args := strings.Fields(update.Message.Text[1:])
	if len(args) == 0 {
		return
	}

	command := args[0]

	if command == "menu" {
		msg := tg.NewMessage(update.Message.Chat.ID, "Dolbaeb?")
		msg.ReplyMarkup = mainMenuKeyboard
		openMenu(msg, "main")
	} else if command == "start" {
		createAccount(update)
	} else {
		msg := tg.NewMessage(update.Message.Chat.ID, "unknown command")
		_, err := glb.Bot.Send(msg)
		glb.CE(err)
	}
}

func actionHandler(update tg.Update) {
	message := update.Message.Text

	for i := 0; i < len(glb.OpenedMessages); i++ {
		if update.Message.Chat.ID == glb.OpenedMessages[i].ChatID {
			_, err := glb.DB.Query(`
				insert into users (id, name, date_created, redeemed_codes)
				values(?,?,now(),?);`,
				update.Message.Chat.ID, message, "[]")
			glb.CE(err)

			msg := tg.NewMessage(update.Message.Chat.ID, "Не иди нахуй")
			_, err = glb.Bot.Send(msg)
			glb.CE(err)

			return
		}
	}
	msg := tg.NewMessage(update.Message.Chat.ID, "Ты ебанат?")
	_, err := glb.Bot.Send(msg)
	glb.CE(err)

}
