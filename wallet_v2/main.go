package main

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"net/http"
	"strconv"
)

var db *sql.DB

type Account struct {
	Id   string
	Balance int
}

type Res struct{
	Amount  int `form:"amount" json:"amount"`
	Status  string `form:"status" json:"status"`
	Message string `form:"message" json:"message"`
}

func dbConnect() {
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

func main() {
	dbConnect()

	router := gin.Default()
	router.POST("/deposit", deposit)
	router.POST("/withdraw", withdraw)
	router.GET("/balance", getBalance)

	router.Run(":80")
}

func queryBalance(id string) (balance int, valid bool){
	var acct Account
	acct.Id = id
	err := db.QueryRow("SELECT balance FROM account WHERE id = ?", acct.Id).Scan(&acct.Balance)
	if err != nil{
		return
	}
	return acct.Balance, true
}

func transact(id string, amount int)(valid bool){
	_, err:= db.Exec("UPDATE account SET balance = balance + ? WHERE id = ?",amount, id)
	if err != nil{
		return
	}
	return true
}

func transactAndQueryBalance(id string, amount int) (balance int, valid bool){
	valid = transact(id, amount)
	if valid {
		balance, valid = queryBalance(id)
	}
	return
}

func getBalance(context *gin.Context) {
	var res = Res{
		Amount: 0,
		Status: "failed",
		Message: "",
	}

	user := context.Query("user")
	amount, valid := queryBalance(user)
	if valid{
		res.Amount = amount
		res.Status = "ok"
	} else {
		res.Message = "操作失敗，查無此使用者！"
	}

	context.JSON(http.StatusOK, res)
}

func deposit(context *gin.Context) {
	var res = Res{
		Amount: 0,
		Status: "failed",
		Message: "",
	}

	user := context.PostForm("user")
	amount, err := strconv.Atoi(context.PostForm("amount"))
	if err != nil {
		res.Message = "操作失敗，輸入有誤！"
	} else {
		if amount <= 0 {
			res.Message = "操作失敗，金額需大於0元！"
		} else {
			//執行存款交易
			balance, valid := transactAndQueryBalance(user, amount)
			if valid {
				res.Amount = balance
				res.Status = "ok"
			} else {
				res.Message = "操作失敗，查無此使用者！"
			}
		}
	}
	context.JSON(http.StatusOK, res)
}

func withdraw(context *gin.Context) {
	var res = Res{
		Amount: 0,
		Status: "failed",
		Message: "",
	}

	user := context.PostForm("user")
	amount, err := strconv.Atoi(context.PostForm("amount"))
	if err != nil {
		res.Message = "操作失敗，輸入有誤！"
	} else {
		if amount <= 0 {
			res.Message = "操作失敗，金額需大於0元！"
		} else {
			//判斷餘額，執行提款交易
			balance, valid := queryBalance(user)
			if valid {
				if balance-amount < 0{
					res.Message = "操作失敗，餘額不足！"
				} else {
					balance, _ = transactAndQueryBalance(user, -amount)
					res.Status = "ok"
					res.Amount = balance
				}
			} else {
				res.Message = "操作失敗，查無此使用者！"
			}
		}
	}
	context.JSON(http.StatusOK, res)
}