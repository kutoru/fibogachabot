package main

import (
	"database/sql"
	"fmt"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
	tg "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/kutoru/fibogachabot/pkg/global"
	updatehandler "github.com/kutoru/fibogachabot/pkg/update_handler"
)

type User struct {
	ID          int
	Name        string
	DateCreated string
}

func ConnectToDB() *sql.DB {
	fmt.Println("Connecting to db")

	dbInfo := fmt.Sprintf("root:%s@tcp(db:3306)/fibobase", os.Getenv("DB_PASS"))
	conn, err := sql.Open("mysql", dbInfo)
	global.CE(err)

	// go can start before the database sometimes, this avoids any issues related to that
	for conn.Ping() != nil {
		fmt.Println("Attempting connection to db")
		time.Sleep(2 * time.Second)
	}

	fmt.Println("Connected")
	return conn
}

func dbtest() {
	conn := ConnectToDB()

	_, err := conn.Exec(`DROP TABLE IF EXISTS users;`)
	global.CE(err)
	fmt.Println("Dropped users table")

	_, err = conn.Exec(`
	create table users (
		id int auto_increment,
		name varchar(255) not null,
		date_created datetime not null,
		primary key (id)
	);
	`)
	global.CE(err)
	fmt.Println("Created users table")

	_, err = conn.Query(`
		INSERT INTO users(name, date_created) VALUES
		('vostexx', now()),
		('mgosu', now());
	`)
	global.CE(err)
	fmt.Println("Inserted into users")

	results, err := conn.Query("SELECT * FROM users;")
	global.CE(err)
	fmt.Println("Selected from users")

	for results.Next() {
		var user User
		err = results.Scan(
			&user.ID,
			&user.Name,
			&user.DateCreated,
		)
		global.CE(err)

		fmt.Println(user)
	}
}

var numericKeyboard = tg.NewInlineKeyboardMarkup(
	tg.NewInlineKeyboardRow(
		tg.NewInlineKeyboardButtonData("1", "1"),
		tg.NewInlineKeyboardButtonData("2", "2"),
		tg.NewInlineKeyboardButtonData("3", "3"),
	),
	tg.NewInlineKeyboardRow(
		tg.NewInlineKeyboardButtonData("4", "4"),
		tg.NewInlineKeyboardButtonData("5", "5"),
		tg.NewInlineKeyboardButtonData("6", "6"),
	),
)

func botTest() {
	global.BotInit()
	updateConfig := tg.NewUpdate(0)
	updateConfig.Timeout = 30
	updates := global.Bot.GetUpdatesChan(updateConfig)

	for update := range updates {
		if update.Message != nil {
			msg := tg.NewMessage(update.Message.Chat.ID, update.Message.Text)
			msg.ReplyMarkup = numericKeyboard

			_, err := global.Bot.Send(msg)
			global.CE(err)
		} else if update.CallbackQuery != nil {
			updatehandler.KeyboardHandler(update)
		}
	}
}

func main() {
	fmt.Println("Start")

	botTest()

	fmt.Println("End")
}
