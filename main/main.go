package main

import (
	"database/sql"
	"fmt"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	ID          int
	Name        string
	DateCreated string
}

// check error
func ce(err error) {
	if err != nil {
		fmt.Println(err)
	}
}

func ConnectToDB() *sql.DB {
	fmt.Println("Connecting to db")

	dbInfo := fmt.Sprintf("root:%s@tcp(db:3306)/fibobase", os.Getenv("DB_PASS"))
	conn, err := sql.Open("mysql", dbInfo)
	ce(err)

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
	ce(err)
	fmt.Println("Dropped users table")

	_, err = conn.Exec(`
	create table users (
		id int auto_increment,
		name varchar(255) not null,
		date_created datetime not null,
		primary key (id)
	);
	`)
	ce(err)
	fmt.Println("Created users table")

	_, err = conn.Query(`
		INSERT INTO users(name, date_created) VALUES
		('vostexx', now()),
		('mgosu', now());
	`)
	ce(err)
	fmt.Println("Inserted into users")

	results, err := conn.Query("SELECT * FROM users;")
	ce(err)
	fmt.Println("Selected from users")

	for results.Next() {
		var user User
		err = results.Scan(
			&user.ID,
			&user.Name,
			&user.DateCreated,
		)
		ce(err)

		fmt.Println(user)
	}
}

func main() {
	fmt.Println("Start")

	dbtest()

	fmt.Println("End")
}
