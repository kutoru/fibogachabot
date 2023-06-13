package models

import (
	"database/sql"
	"errors"
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
// Returns potential result error.
func (char *Character) ScanFromResult(result *sql.Rows) error {
	if !result.Next() {
		return errors.New("could not scan Character result because result.Next() is false")
	}

	return result.Scan(
		&char.ID,
		&char.Name,
		&char.Nickname,
		&char.Description,
		&char.Rarity,
		&char.CardPath,
	)
}
