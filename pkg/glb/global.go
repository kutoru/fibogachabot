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

// check error
func CE(err error) {
	if err != nil {
		fmt.Println(err)
	}
}

func LoadEnv() {
	err := godotenv.Load()
	CE(err)
}

func InitializeMenuList() {
	OpenedMenus = &models.Menus{}
}
