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
	userId := update.CallbackQuery.From.ID
	msg := tg.NewMessage(userId, "")

	switch update.CallbackQuery.Data {

	// main menu cases
	case "main_menu":
		msg.Text = mainMenuText
		msg.ReplyMarkup = mainMenuKeyboard
		openMenu(msg, models.MainMenu, nil)
	case "profile":
		msg.Text = getProfileText(userId)
		msg.ReplyMarkup = profileKeyboard
		openMenu(msg, models.ProfileMenu, nil)

		// play cases

		// dream cases

		// shop cases

	// profile cases
	case "profile_characters":
		openListMenu(msg, models.ProfileCharactersMenu, true, false)
	case "profile_characters|0":
		openListMenu(msg, models.ProfileCharactersMenu, false, false)
	case "profile_characters|1":
		openListMenu(msg, models.ProfileCharactersMenu, false, true)
	case "profile_inventory":
		openListMenu(msg, models.ProfileInventoryMenu, true, false)
	case "profile_inventory|0":
		openListMenu(msg, models.ProfileInventoryMenu, false, false)
	case "profile_inventory|1":
		openListMenu(msg, models.ProfileInventoryMenu, false, true)

		// archive cases

		// settings cases

	// not implemented
	default:
		msg.Text = notImplementedText
		msg.ReplyMarkup = notImplementedKeyboard
		openMenu(msg, models.NotImplementedMenu, nil)
	}
}
