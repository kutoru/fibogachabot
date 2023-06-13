package tgmanager

import tg "github.com/go-telegram-bot-api/telegram-bot-api/v5"

var profileKeyboard = tg.NewInlineKeyboardMarkup(
	tg.NewInlineKeyboardRow(
		tg.NewInlineKeyboardButtonData("👥 Characters", "profile_characters"),
		tg.NewInlineKeyboardButtonData("🎏 Quests", "profile_quests"),
		tg.NewInlineKeyboardButtonData("🏆 Achievements", "profile_achievements"),
	),
	tg.NewInlineKeyboardRow(
		tg.NewInlineKeyboardButtonData("🔙 Go Back", "main_menu"),
	),
)

var profileText = "This is a profile menu"
