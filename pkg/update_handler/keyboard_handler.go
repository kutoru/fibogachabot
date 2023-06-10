package updatehandler

import (
	tg "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/kutoru/fibogachabot/pkg/global"
)

var MainMenuKeyboard = tg.NewInlineKeyboardMarkup(
	tg.NewInlineKeyboardRow(
		tg.NewInlineKeyboardButtonData("💫 Play", "1"),
		tg.NewInlineKeyboardButtonData("✨ Dream", "2"),
		tg.NewInlineKeyboardButtonData("💰 Shop", "3"),
	),
	tg.NewInlineKeyboardRow(
		tg.NewInlineKeyboardButtonData("🏆 Statistics", "4"),
		tg.NewInlineKeyboardButtonData("📚 Archive", "5"),
		tg.NewInlineKeyboardButtonData("⚙️ Settings", "6"),
	),
)

func DeletePreviousMessage(update tg.Update) {
	callback := tg.NewDeleteMessage(
		update.CallbackQuery.Message.Chat.ID,
		update.CallbackQuery.Message.MessageID,
	)
	_, err := global.Bot.Request(callback)
	global.CE(err)
}

func KeyboardHandler(update tg.Update) {
	DeletePreviousMessage(update)
	data := update.CallbackQuery.Data
	msg := tg.NewMessage(update.CallbackQuery.From.ID, "")
	switch data {
	case "1":
		msg.Text = "Clicked Play"
	case "2":
		msg.Text = "Clicked Dream"
	case "3":
		msg.Text = "Clicked Shop"
	case "4":
		msg.Text = "Clicked Statistics"
	case "5":
		msg.Text = "Clicked Archive"
	case "6":
		msg.Text = "Clicked Settings"
	}
	msg.ReplyMarkup = MainMenuKeyboard
	_, err := global.Bot.Send(msg)
	global.CE(err)
}
