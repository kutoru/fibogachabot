package dbmanager

import (
	"fmt"

	"github.com/kutoru/fibogachabot/pkg/glb"
	"github.com/kutoru/fibogachabot/pkg/models"
)

// User

func GetUser(userId int64) models.User {
	result, err := glb.DB.Query(`
		select * from users where id = ?;
	`, userId)
	glb.CE(err)

	var user models.User
	if result.Next() {
		err = user.ScanFromResult(result)
		glb.CE(err)
	}

	return user
}

// Dailies

func GetAllDailies(userId int64) []models.Daily {
	result, err := glb.DB.Query(`
		select * from dailies where user_id = ?;
	`, userId)
	glb.CE(err)

	var allDailies []models.Daily

	for result.Next() {
		var daily models.Daily
		err := daily.ScanFromResult(result)
		if err != nil {
			continue
		}

		allDailies = append(allDailies, daily)
	}

	return allDailies
}

// AcqCharacters

func GetAllAcqCharacters(userId int64, loadCharInfo bool) []models.AcqCharacter {
	result, err := glb.DB.Query(`
		select * from acquired_chars where user_id = ?;
	`, userId)
	glb.CE(err)

	var allAcqChars []models.AcqCharacter

	for result.Next() {
		var acqChar models.AcqCharacter
		err := acqChar.ScanFromResult(result)
		if err != nil {
			continue
		}

		if loadCharInfo {
			charResult, err := glb.DB.Query(`
				select * from characters where id = ?;
			`, acqChar.CharacterID)
			glb.CE(err)

			char := &models.Character{}
			if charResult.Next() {
				char.ScanFromResult(charResult)
			}

			acqChar.CharacterInfo = char
		}

		allAcqChars = append(allAcqChars, acqChar)
	}

	return allAcqChars
}

// AcqAchievements

func GetAllAcqAchievements(userId int64, loadAchievementInfo bool) []models.AcqAchievement {
	result, err := glb.DB.Query(`
		select * from acquired_achievements where user_id = ?;
	`, userId)
	glb.CE(err)

	var allAcqAchievs []models.AcqAchievement

	for result.Next() {
		var acqAchiev models.AcqAchievement
		err := acqAchiev.ScanFromResult(result)
		if err != nil {
			glb.CE(err)
			continue
		}

		if loadAchievementInfo {
			giftResult, err := glb.DB.Query(`
				select * from achievements where id = ?;
			`, acqAchiev.AchievementID)
			glb.CE(err)

			achiev := &models.Achievement{}
			if giftResult.Next() {
				achiev.ScanFromResult(giftResult)
			}

			acqAchiev.AchievementInfo = achiev
		}

		allAcqAchievs = append(allAcqAchievs, acqAchiev)
	}

	return allAcqAchievs
}

// AcqGifts

func GetAllAcqGifts(userId int64, loadGiftInfo bool) []models.AcqGift {
	result, err := glb.DB.Query(`
		select * from acquired_gifts where user_id = ?;
	`, userId)
	glb.CE(err)

	var allAcqGifts []models.AcqGift

	for result.Next() {
		var acqGift models.AcqGift
		err := acqGift.ScanFromResult(result)
		if err != nil {
			glb.CE(err)
			continue
		}

		if loadGiftInfo {
			giftResult, err := glb.DB.Query(`
				select * from gifts where id = ?;
			`, acqGift.GiftID)
			glb.CE(err)

			gift := &models.Gift{}
			if giftResult.Next() {
				gift.ScanFromResult(giftResult)
			}

			acqGift.GiftInfo = gift
		}

		allAcqGifts = append(allAcqGifts, acqGift)
	}

	return allAcqGifts
}

// Character

func GetCharacter() models.Character {
	return models.Character{}
}

// AcqQuests
func GetAllAcqQuests(userId int64, loadGiftInfo bool) []models.AcqQuest {
	result, err := glb.DB.Query(`
		select * from acquired_quests where user_id = ?;
	`, userId)
	glb.CE(err)

	var allAcqQuests []models.AcqQuest

	for result.Next() {
		var acqQuest models.AcqQuest
		err := acqQuest.ScanFromResult(result)
		if err != nil {
			glb.CE(err)
			continue
		}

		if loadGiftInfo {
			questResult, err := glb.DB.Query(`
				select * from quests where id = ?;
			`, acqQuest.QuestID)
			glb.CE(err)

			quest := &models.Quest{}
			if questResult.Next() {
				quest.ScanFromResult(questResult)
			}

			acqQuest.QuestInfo = quest
		}

		allAcqQuests = append(allAcqQuests, acqQuest)
	}

	return allAcqQuests
}

func TestPrintAllCharacters() {
	result, err := glb.DB.Query(`
		select * from characters;
	`)
	glb.CE(err)

	for result.Next() {
		var char models.Character
		err = char.ScanFromResult(result)
		glb.CE(err)
		fmt.Printf("%+v\n\n", char)
	}
}

func TestPrintAllUsers() {
	result, err := glb.DB.Query(`
		select * from users;
	`)
	glb.CE(err)

	for result.Next() {
		var user models.User
		err = user.ScanFromResult(result)
		glb.CE(err)
		fmt.Printf("%+v\n", user)

		for _, code := range user.RedeemedCodes {
			fmt.Println(code)
		}
		fmt.Println()
	}
}
