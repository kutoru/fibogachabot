package tgmanager

import (
	tg "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/kutoru/fibogachabot/pkg/models"
)

var notImplementedKeyboard = tg.NewInlineKeyboardMarkup(
	tg.NewInlineKeyboardRow(
		tg.NewInlineKeyboardButtonData("🔙 Go Back", "main_menu"),
	),
)

var notImplementedText = "The menu is not implemented"

func keyboardHandler(update tg.Update) {
	msg := tg.NewMessage(update.CallbackQuery.From.ID, "")

	switch update.CallbackQuery.Data {

	// main menu cases
	case "main_menu":
		msg.Text = mainMenuText
		msg.ReplyMarkup = mainMenuKeyboard
		openMenu(msg, models.MainMenu)
	case "1":
		msg.Text = "Clicked Play"
		msg.ReplyMarkup = mainMenuKeyboard
		openMenu(msg, models.MainMenu)
	case "2":
		msg.Text = "Clicked Dream"
		msg.ReplyMarkup = mainMenuKeyboard
		openMenu(msg, models.MainMenu)
	case "3":
		msg.Text = "Clicked Shop"
		msg.ReplyMarkup = mainMenuKeyboard
		openMenu(msg, models.MainMenu)
	case "profile":
		msg.Text = profileText
		msg.ReplyMarkup = profileKeyboard
		openMenu(msg, models.ProfileMenu)
	case "5":
		msg.Text = "Clicked Archive"
		msg.ReplyMarkup = mainMenuKeyboard
		openMenu(msg, models.MainMenu)
	case "6":
		msg.Text = "Clicked Settings"
		msg.ReplyMarkup = mainMenuKeyboard
		openMenu(msg, models.MainMenu)

		// play cases

		// dream cases

		// shop cases

		// profile cases

		// archive cases

		// settings cases

	// not implemented
	default:
		msg.Text = notImplementedText
		msg.ReplyMarkup = notImplementedKeyboard
		openMenu(msg, models.NotImplementedMenu)
	}
}
