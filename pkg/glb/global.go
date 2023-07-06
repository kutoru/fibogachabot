package glb

import (
	"database/sql"
	"fmt"

	tg "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/joho/godotenv"
	"github.com/kutoru/fibogachabot/pkg/models"
)

var Bot *tg.BotAPI
var DB *sql.DB
var OpenedMenus *models.Menus
var OpenedMessages []*models.OpenedMessage

// Checks error
func CE(err error) {
	if err != nil {
		fmt.Println(err)
	}
}

// Loads environment variables from .env
func LoadEnv() {
	err := godotenv.Load()
	CE(err)
}

// Initializes OpenedMenus list
func InitializeMenuList() {
	OpenedMenus = &models.Menus{}
}
