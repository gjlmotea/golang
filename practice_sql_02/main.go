package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

var db *sql.DB

func main() {
	dbconn, err := sql.Open(
		"mysql",
		"jason:8888@tcp(127.0.0.1:3306)/wallet",
		)
	if err != nil {
		log.Fatal(err)
	}
	defer dbconn.Close()
	dbconn.Ping()
	db = dbconn

	QueryUser("Jason")
	TransactBalance("Jason", -2)
	QueryUser("Jason")
}

type Account struct {
	id   string
	balance int
}

func QueryUser(id string) (balance int){
	var acct Account
	acct.id = id
	err := db.QueryRow("SELECT balance FROM account WHERE id = ?", acct.id).Scan(&acct.balance)
	if err != nil{
		panic((err.Error()))
	}
	log.Println(acct.balance)
	return acct.balance
}

func QueryAllUsers(){
	results, err := db.Query("SELECT * FROM account")
	if err != nil{
		panic((err.Error()))
	}
	for results.Next() {
		var acct Account
		err = results.Scan(&acct.id, &acct.balance)
		log.Println(acct.id, acct.balance)
		//TODO return
	}
}

func TransactBalance(id string, amount int){
	var acct Account
	acct.id = id
	result, err:= db.Exec("UPDATE account SET balance = balance + ? WHERE id = ?",amount, id)
	if err != nil{
		panic((err.Error()))
	}
	log.Println("update data success:",result)
}

