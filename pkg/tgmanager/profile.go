package tgmanager

import (
	"fmt"
	"strings"

	tg "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/kutoru/fibogachabot/pkg/dbmanager"
)

var profileKeyboard = tg.NewInlineKeyboardMarkup(
	tg.NewInlineKeyboardRow(
		tg.NewInlineKeyboardButtonData("👥 Characters", "profile_characters"),
		tg.NewInlineKeyboardButtonData("🎒 Inventory", "profile_inventory"),
	),
	tg.NewInlineKeyboardRow(
		tg.NewInlineKeyboardButtonData("🏆 Achievements", "profile_achievements"),
		tg.NewInlineKeyboardButtonData("🎏 Quests", "profile_quests"),
	),
	tg.NewInlineKeyboardRow(
		tg.NewInlineKeyboardButtonData("🔙 Go Back", "main_menu"),
	),
)

func getProfileText(userId int64) string {
	user := dbmanager.GetUser(userId)

	name := user.Name
	dateCreated := strings.Fields(user.DateCreated)[0]

	totalCharacters := len(dbmanager.GetAllAcqCharacters(userId, false))
	totalAchievements := len(dbmanager.GetAllAcqAchievements(userId, false))
	dailiesCompleted := user.DailiesCompleted
	questsCompleted := 0
	giftsBought := user.GiftsBought
	giftsGifted := user.GiftsGifted
	totalIllusions := user.TotalIllusions
	totalXCards := user.TotalXCards

	currentIllusions := user.Illusions
	currentXCards := user.XCards
	currentGifts := 0

	for _, quest := range dbmanager.GetAllAcqQuests(userId, false) {
		if quest.Completed {
			questsCompleted++
		}
	}

	for _, gift := range dbmanager.GetAllAcqGifts(userId, false) {
		currentGifts += gift.Amount
	}

	profileText := fmt.Sprintf(
		`Name: %s
Date registered: %s

Total characters: %d
Total achievements: %d
Total dailies completed: %d
Total quests completed: %d
Total gifts bought: %d
Total gifts gifted: %d
Total Illusions: %d
Total X-Cards: %d

Current Illusions: %d
Current X-Cards: %d
Current Gifts: %d`,
		name, dateCreated,
		totalCharacters, totalAchievements, dailiesCompleted, questsCompleted, giftsBought, giftsGifted, totalIllusions, totalXCards,
		currentIllusions, currentXCards, currentGifts)

	return profileText
}
