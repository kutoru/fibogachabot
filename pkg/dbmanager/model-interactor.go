package dbmanager

import (
	"fmt"

	"github.com/kutoru/fibogachabot/pkg/glb"
	"github.com/kutoru/fibogachabot/pkg/models"
)

// User

func LoadUser(userId int64) models.User {
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

// TODO: change this so it works similarly to GetAllAcqCharacters
func GetTotalAcqAchievements(userId int64) int {
	result, err := glb.DB.Query(`
		select * from acquired_achievements where user_id = ?;
	`, userId)
	glb.CE(err)

	total := 0
	for result.Next() {
		total += 1
	}

	return total
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

func LoadCharacter() models.Character {
	return models.Character{}
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
