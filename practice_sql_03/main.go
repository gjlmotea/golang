package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

var db *sql.DB

func main() {
	connectDb()

	//createDb()

	//createTable()

	//alterTable()

	//insert()

	query()
}

func connectDb() {
	Dbconn, err := sql.Open("mysql", "root:8888@tcp(127.0.0.1:3306)/?parseTime=true")
	if err != nil {
		log.Fatal(err)
	}
	//defer Dbconn.Close()

	db = Dbconn
	db.SetMaxOpenConns(1)
	db.SetMaxIdleConns(1)
	err = db.Ping()
	if err != nil {
		log.Fatalln(err)
	}

}

func createDb() {
	_, err := db.Exec("CREATE DATABASE `tutor`;")
	if err != nil {
		log.Fatalln(err)
	}
}

func createTable() {
	_, err := db.Exec("CREATE TABLE IF NOT EXISTS tutor.hello(world varchar(50))")
	if err != nil {
		log.Fatalln(err)
	}
}

func alterTable() {
	//_, err := db.Exec("ALTER TABLE tutor.hello ADD COLUMN id int(10)")
	_, err := db.Exec("ALTER TABLE tutor.hello ADD COLUMN id int(10) PRIMARY KEY AUTO_INCREMENT")
	if err != nil {
		log.Fatalln(err)
	}
}

func insert() {
	rs, err := db.Exec("INSERT INTO tutor.hello(world) VALUES ('wqqqq')")
	if err != nil {
		log.Fatalln(err)
	}
	rowCount, err := rs.RowsAffected()
	rowId, err := rs.LastInsertId()
	if err != nil {
		log.Fatalln(err)
	}
	log.Printf("inserted %d rows where id = %d", rowCount, rowId)
}

func query() {
	rows, err := db.Query("SELECT world FROM tutor.hello")
	if err != nil {
		log.Fatalln(err)
	}

	for rows.Next() {
		var s string
		err = rows.Scan(&s)
		if err != nil {
			log.Fatalln(err)
		}
		log.Printf("found row containing %q", s)
	}
	rows.Close()
}
