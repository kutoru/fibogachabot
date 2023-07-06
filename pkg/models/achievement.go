package models

import "database/sql"

type Achievement struct {
	ID          int
	Title       string
	Description string
}

type AcqAchievement struct {
	UserID          int64
	AchievementID   int
	DateAcquired    string
	AchievementInfo *Achievement
}

func (achiev *Achievement) ScanFromResult(result *sql.Rows) error {
	return result.Scan(
		&achiev.ID,
		&achiev.Title,
		&achiev.Description,
	)
}

func (acqAchiev *AcqAchievement) ScanFromResult(result *sql.Rows) error {
	return result.Scan(
		&acqAchiev.UserID,
		&acqAchiev.AchievementID,
		&acqAchiev.DateAcquired,
	)
}
