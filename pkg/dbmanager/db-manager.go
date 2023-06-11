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

func InitializeDB() *sql.DB {
	fmt.Println("Connecting to db")

	dbInfo := fmt.Sprintf("root:%s@tcp(db:3306)/fibobase", os.Getenv("DB_PASS"))

	var err error
	glb.DB, err = sql.Open("mysql", dbInfo)
	glb.CE(err)

	// go can start before the database sometimes, this avoids any issues related to that
	for glb.DB.Ping() != nil {
		fmt.Println("Attempting connection to db")
		time.Sleep(2 * time.Second)
	}

	fmt.Println("Connected")
	return glb.DB
}

func dbtest() {
	_, err := glb.DB.Exec(`DROP TABLE IF EXISTS users;`)
	glb.CE(err)
	fmt.Println("Dropped users table")

	_, err = glb.DB.Exec(`
	create table users (
		id int auto_increment,
		name varchar(255) not null,
		date_created datetime not null,
		primary key (id)
	);
	`)
	glb.CE(err)
	fmt.Println("Created users table")

	_, err = glb.DB.Query(`
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
