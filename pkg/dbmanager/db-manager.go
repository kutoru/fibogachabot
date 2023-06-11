package dbmanager

import (
	"database/sql"
	"fmt"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/kutoru/fibogachabot/pkg/glb"
	"github.com/kutoru/fibogachabot/pkg/models"
)

func ConnectToDB() {
	fmt.Println("Connecting to db")

	dbInfo := fmt.Sprintf("root:%s@tcp(localhost:3306)/fibobase?multiStatements=true", os.Getenv("DB_PASS"))

	var err error
	glb.DB, err = sql.Open("mysql", dbInfo)
	glb.CE(err)

	// go can start before the database sometimes, this avoids any issues related to that
	for glb.DB.Ping() != nil {
		fmt.Println("Attempting connection to db")
		time.Sleep(3 * time.Second)
	}

	fmt.Println("Connected")
}

func dbtest() {
	_, err := glb.DB.Query(`
		INSERT INTO users(name, date_created) VALUES
		('vostexx', now()),
		('mgosu', now());
	`)
	glb.CE(err)
	fmt.Println("Inserted into users")

	results, err := glb.DB.Query("SELECT * FROM users;")
	glb.CE(err)
	fmt.Println("Selected from users")

	for results.Next() {
		var user models.User
		err = results.Scan(
			&user.ID,
			&user.Name,
			&user.DateCreated,
		)
		glb.CE(err)

		fmt.Println(user)
	}
}

// Executes create_db.sql on the connected database
func InitializeDB() {
	cont, err := os.ReadFile("./create_db.sql")
	glb.CE(err)
	_, err = glb.DB.Exec(string(cont))
	glb.CE(err)
}
