package models

import (
	"database/sql"
	"encoding/json"
)

type QuestRequirements struct {
	AccountLVL    int
	QuestIDs      []int
	FriendshipLVL int
}

type QuestRewards struct {
	Illusions int
	XCards    int
}

type Quest struct {
	ID           int
	CharID       int
	Title        string
	Description  string
	Type         string
	Requirements QuestRequirements
	Rewards      QuestRewards
}

type AcqQuest struct {
	UserID    int64
	QuestID   int
	Completed bool
	QuestInfo *Quest
}

func (quest *Quest) ScanFromResult(result *sql.Rows) error {
	var bytesRequirements []uint8
	var bytesRewards []uint8

	err := result.Scan(
		&quest.ID,
		&quest.CharID,
		&quest.Title,
		&quest.Description,
		&quest.Type,
		&bytesRequirements,
		&bytesRewards,
	)
	if err != nil {
		return err
	}

	err = json.Unmarshal(bytesRequirements, &quest.Requirements)
	if err != nil {
		return err
	}

	return json.Unmarshal(bytesRewards, &quest.Rewards)
}

func (acqQuest *AcqQuest) ScanFromResult(result *sql.Rows) error {
	return result.Scan(
		&acqQuest.UserID,
		&acqQuest.QuestID,
		&acqQuest.Completed,
	)
}
