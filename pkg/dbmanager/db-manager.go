package dbmanager

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/kutoru/fibogachabot/pkg/glb"
	"github.com/kutoru/fibogachabot/pkg/models"
)

func ConnectToDB() {
	dbInfo := fmt.Sprintf(
		"root:%s@tcp(localhost:3306)/%s?multiStatements=true&loc=Europe%%2FLondon",
		os.Getenv("DB_PASS"), os.Getenv("DB_NAME"),
	)

	var err error
	glb.DB, err = sql.Open("mysql", dbInfo)
	glb.CE(err)

	// go can start before the database sometimes, this avoids any issues related to that
	for glb.DB.Ping() != nil {
		fmt.Println("Attempting connection to db")
		time.Sleep(3 * time.Second)
	}

	fmt.Println("Connected to db")
}

func ResetDB() {
	script, err := os.ReadFile("./scripts/reset_db.sql")
	glb.CE(err)

	_, err = glb.DB.Exec(string(script))
	glb.CE(err)

	fmt.Println("Reset the DB")
}

// Executes create_db.sql on the connected database and loads all the static data into it
// TODO: make this function run only if the database is not initialized already
func InitializeDB() {
	script, err := os.ReadFile("./scripts/create_db.sql")
	glb.CE(err)

	_, err = glb.DB.Exec(string(script))
	glb.CE(err)

	fmt.Println("Initialized the DB")

	LoadCharactersIntoDB()
	LoadGiftsIntoDB()
}

func LoadCharactersIntoDB() {
	jsonDir := "./assets/char_jsons"
	items, err := os.ReadDir(jsonDir)
	glb.CE(err)

	for _, item := range items {
		json_data, err := os.ReadFile(jsonDir + "/" + item.Name())
		glb.CE(err)

		var char models.Character
		err = json.Unmarshal(json_data, &char)
		glb.CE(err)

		if len(char.Nickname) == 0 {
			char.Nickname = "No nickname"
		}

		if len(char.Description) == 0 {
			char.Description = "No description"
		}

		_, err = glb.DB.Query(`
			insert into characters
			values (?, ?, ?, ?, ?);
		`, char.ID, char.Name, char.Nickname, char.Description, char.Rarity)
		glb.CE(err)
	}
}

func LoadGiftsIntoDB() {
	var gifts []models.Gift

	json_data, err := os.ReadFile("./assets/gifts.json")
	glb.CE(err)

	err = json.Unmarshal(json_data, &gifts)
	glb.CE(err)

	for _, gift := range gifts {
		_, err = glb.DB.Query(`
			insert into gifts
			values (?, ?, ?, ?, ?);
		`, gift.ID, gift.Name, gift.Type, gift.Rarity, gift.Price)
		glb.CE(err)
	}
}
