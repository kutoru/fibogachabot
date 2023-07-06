package tgmanager

import (
	tg "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/kutoru/fibogachabot/pkg/glb"
	"github.com/kutoru/fibogachabot/pkg/models"
)

func createAccount(update tg.Update) {
	result, err := glb.DB.Query(`select id from users where id = ?;`, update.Message.Chat.ID)
	glb.CE(err)

	if !result.Next() {
		msg := tg.NewMessage(update.Message.Chat.ID, "Hello! Enter a username that is going to be used throughout the game. Or enter \"none\" to use "+update.FromChat().UserName)
		bot_msg, err := glb.Bot.Send(msg)
		glb.CE(err)

		OpenedMessage := &models.OpenedMessage{
			ChatID:       update.Message.Chat.ID,
			Action:       0,
			BotMessageID: bot_msg.MessageID,
		}
		glb.OpenedMessages = append(glb.OpenedMessages, OpenedMessage)

	} else {
		msg := tg.NewMessage(update.Message.Chat.ID, "Иди нахуй")
		_, err = glb.Bot.Send(msg)
		glb.CE(err)
	}
}
