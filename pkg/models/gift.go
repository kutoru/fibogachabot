package models

import "database/sql"

type GiftType int

const (
	Food GiftType = iota
	Tech
	Music
	Literature
	Art
	Toys
)

type Gift struct {
	ID     int
	Name   string
	Type   string
	Rarity int
}

type AcqGift struct {
	UserID int64
	GiftID int
	Amount int
}

func (acqGift *AcqGift) ScanFromResult(result *sql.Rows) error {
	return result.Scan(
		&acqGift.UserID,
		&acqGift.GiftID,
		&acqGift.Amount,
	)
}
