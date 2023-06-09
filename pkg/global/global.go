package global

import (
	"fmt"
	"os"

	tg "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

var Bot *tg.BotAPI

func BotInit() {
	var err error
	Bot, err = tg.NewBotAPI(os.Getenv("token"))
	CE(err)
	Bot.Debug = true
}

// check error
func CE(err error) {
	if err != nil {
		fmt.Println(err)
	}
}
