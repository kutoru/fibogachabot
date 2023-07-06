package models

type UserAction int

const (
	account_creation UserAction = iota
)

type OpenedMessage struct {
	ChatID       int64
	Action       UserAction
	BotMessageID int
}
