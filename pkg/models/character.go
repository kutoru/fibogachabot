package models

import (
	"database/sql"
	"encoding/json"
)

type Character struct {
	ID          int
	Name        string
	Nickname    string
	Description string
	Rarity      int
}

type AcqCharacter struct {
	UserID          int64
	CharacterID     int
	FriendshipEXP   int
	FriendshipLVL   int
	Enigma          int
	CompletedQuests []string
	ReceivedGifts   []string
	DateAcquired    string
	CharacterInfo   *Character
}

// Scans into the struct from the DB result.
//
// This function does not call result.Next(), so you will have to do it beforehand.
//
// Returns potential result error.
func (char *Character) ScanFromResult(result *sql.Rows) error {
	return result.Scan(
		&char.ID,
		&char.Name,
		&char.Nickname,
		&char.Description,
		&char.Rarity,
	)
}

func (acqChar *AcqCharacter) ScanFromResult(result *sql.Rows) error {
	var bytesCompletedQuests []uint8
	var bytesReceivedGifts []uint8

	err := result.Scan(
		&acqChar.UserID,
		&acqChar.CharacterID,
		&acqChar.FriendshipEXP,
		&acqChar.FriendshipLVL,
		&acqChar.Enigma,
		&bytesCompletedQuests,
		&bytesReceivedGifts,
		&acqChar.DateAcquired,
	)

	if err != nil {
		return err
	}

	err = json.Unmarshal(bytesCompletedQuests, &acqChar.CompletedQuests)
	if err != nil {
		return err
	}

	err = json.Unmarshal(bytesReceivedGifts, &acqChar.ReceivedGifts)

	return err
}
