package tgmanager

import (
	tg "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/kutoru/fibogachabot/pkg/glb"
	"github.com/kutoru/fibogachabot/pkg/models"
)

func openMenu(msg tg.MessageConfig, menuType models.MenuType, menuArgs ...string) {
	// fmt.Println("Before:")
	// glb.OpenedMenus.PrintAll()

	userId := msg.ChatID

	openedMenu := glb.OpenedMenus.GetMenu(userId)
	if openedMenu != nil {
		deleteMessage(openedMenu.UserID, openedMenu.MessageID)
	}

	sentMessage, err := glb.Bot.Send(msg)
	glb.CE(err)

	glb.OpenedMenus.Append(&models.Menu{
		UserID:    userId,
		MenuType:  menuType,
		MessageID: sentMessage.MessageID,
		Args:      menuArgs,
	})

	// fmt.Println("After:")
	// glb.OpenedMenus.PrintAll()
}
