package tgmanager

import (
	"fmt"
	"strings"

	tg "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/kutoru/fibogachabot/pkg/glb"
)

func MessageHandler(update tg.Update) {
	if strings.HasPrefix(update.Message.Text, "/") {
		commandHandler(update)
	} else {
		fmt.Println("Not a command")
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
		msg := tg.NewMessage(update.Message.Chat.ID, "start not implemented")
		_, err := glb.Bot.Send(msg)
		glb.CE(err)
	} else {
		msg := tg.NewMessage(update.Message.Chat.ID, "unknown command")
		_, err := glb.Bot.Send(msg)
		glb.CE(err)
	}
}
