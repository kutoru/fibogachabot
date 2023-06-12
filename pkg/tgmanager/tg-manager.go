package tgmanager

import (
	"fmt"
	"os"

	tg "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/kutoru/fibogachabot/pkg/glb"
)

func InitializeBot() {
	var err error
	glb.Bot, err = tg.NewBotAPI(os.Getenv("BOT_TOKEN"))
	glb.CE(err)
	glb.Bot.Debug = true
	fmt.Printf("Initialized @%s\n", glb.Bot.Self.UserName)
}

func StartPolling() {
	updateConfig := tg.NewUpdate(0)
	updateConfig.Timeout = 30
	updates := glb.Bot.GetUpdatesChan(updateConfig)

	for update := range updates {
		if update.Message != nil {
			messageHandler(update)
		} else if update.CallbackQuery != nil {
			keyboardHandler(update)
		}
	}
}

func deleteMessage(chatId int64, messageId int) {
	callback := tg.NewDeleteMessage(chatId, messageId)
	_, err := glb.Bot.Request(callback)
	glb.CE(err)

	glb.OpenedMenus.Remove(chatId)
}
