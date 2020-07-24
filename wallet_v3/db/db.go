package db

import (
	"database/sql"
	"log"
	_ "github.com/go-sql-driver/mysql"
)

type Account struct {
	Id   string
	Balance int
}
var db *sql.DB

func init() {
	dbconn, err := sql.Open(
		"mysql",
		"jason:8888@tcp(127.0.0.1:3306)/wallet",
	)
	if err != nil {
		log.Fatal(err)
	}
	//defer dbconn.Close()
	err = dbconn.Ping()
	if err != nil{
		log.Fatalln(err)
	}
	db = dbconn
}

func QueryBalance(id string) (balance int, valid bool){
	var acct Account
	acct.Id = id
	err := db.QueryRow("SELECT balance FROM account WHERE id = ?", acct.Id).Scan(&acct.Balance)
	/*if err == sql.ErrNoRows{
		//查無使用者
	}*/
	if err != nil{
		return
	}
	return acct.Balance, true
}

func Transact(id string, amount int)(valid bool){
	_, err:= db.Exec("UPDATE account SET balance = balance + ? WHERE id = ?",amount, id)
	if err != nil{
		return
	}
	return true
}

func TransactAndQueryBalance(id string, amount int) (balance int, valid bool){
	valid = Transact(id, amount)
	if valid {
		balance, valid = QueryBalance(id)
	}
	return
}

func InsertId(id string)(valid bool){
	_, err:= db.Exec("INSERT INTO account (id, balance) VALUES (?, 0)", id)
	if err != nil{
		return
	}
	return true
}