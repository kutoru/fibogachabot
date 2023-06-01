package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	Name string `json:"name"`
}

func main() {
	fmt.Println("amogus")
	db, err := sql.Open("mysql", "root:amogus@tcp(localhost:3306)/Fibobase")

	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	insert, err := db.Query("INSERT INTO Users VALUES ('AMONGUS')")

	results, err := db.Query("SELECT name FROM Users")
	if err != nil {
		panic(err.Error())
	}

	defer insert.Close()

	for results.Next() {
		var user User
		err = results.Scan(&user.Name)
		if err != nil {
			panic(err.Error())
		}

		fmt.Println(user.Name)
	}

	fmt.Println("Sus")
}
