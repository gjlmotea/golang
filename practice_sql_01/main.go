package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

func main() {

	db, err := sql.Open(
		"mysql",
		"jason:8888@tcp(127.0.0.1:3306)/wallet",
	)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	QueryAccount(db)

}

type User struct {
	id   string
	balance int
}

func QueryAccount(db *sql.DB){
	results, err := db.Query("SELECT * FROM account")
	if err != nil{
		panic((err.Error()))
	}

	for results.Next() {
		var user User
		err = results.Scan(&user.id, &user.balance)
		log.Println(user.id, user.balance)
	}
}