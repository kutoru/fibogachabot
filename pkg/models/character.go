package models

import (
	"database/sql"
)

type Character struct {
	ID          int
	Name        string
	Nickname    string
	Description string
	Rarity      int
	CardPath    string
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
		&char.CardPath,
	)
}
