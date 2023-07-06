package models

import "database/sql"

type Daily struct {
	UserID     int64
	DailyIndex int
	Type       string
	Completed  bool
}

func (daily *Daily) ScanFromResult(result *sql.Rows) error {
	return result.Scan(
		&daily.UserID,
		&daily.DailyIndex,
		&daily.Type,
		&daily.Completed,
	)
}
