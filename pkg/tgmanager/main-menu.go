package tgmanager

import tg "github.com/go-telegram-bot-api/telegram-bot-api/v5"

var mainMenuKeyboard = tg.NewInlineKeyboardMarkup(
	tg.NewInlineKeyboardRow(
		tg.NewInlineKeyboardButtonData("💫 Play", "1"),
		tg.NewInlineKeyboardButtonData("✨ Dream", "2"),
		tg.NewInlineKeyboardButtonData("💰 Shop", "3"),
	),
	tg.NewInlineKeyboardRow(
		tg.NewInlineKeyboardButtonData("👤 Profile", "profile"),
		tg.NewInlineKeyboardButtonData("📚 Archive", "5"),
		tg.NewInlineKeyboardButtonData("⚙️ Settings", "6"),
	),
)

var mainMenuText = "Menu"
