package tgmanager

import tg "github.com/go-telegram-bot-api/telegram-bot-api/v5"

var mainMenuKeyboard = tg.NewInlineKeyboardMarkup(
	tg.NewInlineKeyboardRow(
		tg.NewInlineKeyboardButtonData("💫 Play", "play_menu"),
		tg.NewInlineKeyboardButtonData("✨ Dream", "dream_menu"),
		tg.NewInlineKeyboardButtonData("💰 Shop", "shop_menu"),
	),
	tg.NewInlineKeyboardRow(
		tg.NewInlineKeyboardButtonData("👤 Profile", "profile"),
		tg.NewInlineKeyboardButtonData("📚 Archive", "archive"),
		tg.NewInlineKeyboardButtonData("⚙️ Settings", "settings"),
	),
)

var mainMenuText = "Menu"
