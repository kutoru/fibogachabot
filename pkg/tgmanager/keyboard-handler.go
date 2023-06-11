package tgmanager

import (
	tg "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

var mainMenuKeyboard = tg.NewInlineKeyboardMarkup(
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

func keyboardHandler(update tg.Update) {
	msg := tg.NewMessage(update.CallbackQuery.From.ID, "")

	switch update.CallbackQuery.Data {
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

	msg.ReplyMarkup = mainMenuKeyboard
	openMenu(msg, "main")
}
