package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"reflect"
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

	results, err := db.Query("SELECT * FROM account")
	if err != nil{
		panic((err.Error()))
	}
	fmt.Println()
	fmt.Println(results)
	fmt.Println(reflect.TypeOf(results))

	for results.Next() {
		var name string
		var bal int
		err = results.Scan(&name, &bal)
		log.Println(name, bal)
	}
}
