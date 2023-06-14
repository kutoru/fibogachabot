package dbmanager

import (
	"fmt"

	"github.com/kutoru/fibogachabot/pkg/glb"
	"github.com/kutoru/fibogachabot/pkg/models"
)

// Users

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

func GetTotalAcqCharacters(userId int64) int {
	result, err := glb.DB.Query(`
		select * from acquired_chars where user_id = ?;
	`, userId)
	glb.CE(err)

	total := 0
	for result.Next() {
		total += 1
	}

	return total
}

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

func GetTotalAcqGifts(userId int64) int {
	result, err := glb.DB.Query(`
		select * from acquired_gifts where user_id = ?;
	`, userId)
	glb.CE(err)

	total := 0
	for result.Next() {
		var acqGift models.AcqGift
		err = acqGift.ScanFromResult(result)

		if err == nil {
			total += acqGift.Amount
		}
	}

	return total
}

func TestReadAllUsers() {
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

// Characters

func LoadCharacter() models.Character {
	return models.Character{}
}

func TestReadAllCharacters() {
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
