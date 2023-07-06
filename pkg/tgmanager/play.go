package tgmanager

import (
	"fmt"

	tg "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/kutoru/fibogachabot/pkg/dbmanager"
)

var playMenuKeyboard = tg.NewInlineKeyboardMarkup(
	tg.NewInlineKeyboardRow(
		tg.NewInlineKeyboardButtonData("🍆 Quests", "play_quests"),
		tg.NewInlineKeyboardButtonData("🦽 Dailies", "play_dailies"),
	),
	tg.NewInlineKeyboardRow(
		tg.NewInlineKeyboardButtonData("🔙 Go Back", "main_menu"),
	),
)

func getPlayMenuText(userId int64) string {
	text := "Here, you can select to either play the quests that you have unlocked or your unfinished dailies.\n"

	acqQuests := dbmanager.GetAllAcqQuests(userId, false)
	avaliableQuests := 0
	for _, quest := range acqQuests {
		if !quest.Completed {
			avaliableQuests++
		}
	}

	text += fmt.Sprintf("\nAvailable quests: %d", avaliableQuests)

	dailies := dbmanager.GetAllDailies(userId)
	unfinishedDailies := 0
	for _, daily := range dailies {
		if !daily.Completed {
			unfinishedDailies++
		}
	}

	text += fmt.Sprintf("\nUnfinished dailies: %d", unfinishedDailies)

	return text
}
