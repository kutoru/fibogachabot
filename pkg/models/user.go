package models

import "database/sql"

type User struct {
	ID            int
	Name          string
	DateCreated   string
	Illusions     int
	XCards        int
	Notifications bool
	RedeemedCodes []string
}

func (user *User) ScanFromResult(result *sql.Rows) error {
	return result.Scan(
		&user.ID,
		&user.Name,
		&user.DateCreated,
		&user.Illusions,
		&user.XCards,
		&user.Notifications,
		&user.RedeemedCodes,
	)
}
