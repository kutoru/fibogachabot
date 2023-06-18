package models

import (
	"database/sql"
	"encoding/json"
)

type User struct {
	ID               int64
	Name             string
	DateCreated      string
	Illusions        int
	TotalIllusions   int
	XCards           int
	TotalXCards      int
	GiftsBought      int
	GiftsGifted      int
	DailiesCompleted int
	Notifications    bool
	RedeemedCodes    []string
}

func (user *User) ScanFromResult(result *sql.Rows) error {
	var redeemedCodesBytes []uint8

	err := result.Scan(
		&user.ID,
		&user.Name,
		&user.DateCreated,
		&user.Illusions,
		&user.TotalIllusions,
		&user.XCards,
		&user.TotalXCards,
		&user.GiftsBought,
		&user.GiftsGifted,
		&user.DailiesCompleted,
		&user.Notifications,
		&redeemedCodesBytes,
	)

	if err != nil {
		return err
	}

	return json.Unmarshal(redeemedCodesBytes, &user.RedeemedCodes)
}
