package tgmanager

import (
	"fmt"
	"strings"

	tg "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/kutoru/fibogachabot/pkg/glb"
	"github.com/kutoru/fibogachabot/pkg/models"
)

func messageHandler(update tg.Update) {
	if strings.HasPrefix(update.Message.Text, "/") {
		commandHandler(update)
	} else {
		fmt.Println("Not a command")
	}
}

func commandHandler(update tg.Update) {
	userId := update.Message.Chat.ID
	args := strings.Fields(update.Message.Text[1:])
	if len(args) == 0 {
		return
	}

	command := args[0]

	if command == "menu" {
		msg := tg.NewMessage(userId, "Dolbaeb?")
		msg.ReplyMarkup = mainMenuKeyboard
		openMenu(msg, models.MainMenu, nil)

	} else if command == "start" {
		msg := tg.NewMessage(userId, "Start not implemented")
		_, err := glb.Bot.Send(msg)
		glb.CE(err)

	} else if command == "close" { // test command, delete it later
		if glb.OpenedMenus.UserHasMenu(userId) {
			currMenu := glb.OpenedMenus.GetMenu(userId)
			if currMenu.MenuType == models.MainMenu {
				deleteMessage(userId, currMenu.MessageID)
			}
		}

	} else {
		msg := tg.NewMessage(userId, "Unknown command")
		_, err := glb.Bot.Send(msg)
		glb.CE(err)
	}
}
