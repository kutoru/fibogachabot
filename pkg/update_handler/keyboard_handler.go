package updatehandler

import (
	tg "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/kutoru/fibogachabot/pkg/global"
)

func KeyboardHandler(update tg.Update) {
	if update.CallbackQuery.Data == "1" {
		callback := tg.NewDeleteMessage(
			update.CallbackQuery.Message.Chat.ID,
			update.CallbackQuery.Message.MessageID,
		)
		_, err := global.Bot.Request(callback)
		global.CE(err)
	}
}
