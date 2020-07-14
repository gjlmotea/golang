package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

var balance int = 0

func main() {
	router := gin.Default()
	router.POST("/deposit/:input", deposit)
	router.POST("/withdraw/:input", withdraw)
	router.GET("/balance/", getBalance)

	router.Run(":80")
}

func getBalance(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"amount": balance,
		"status": "ok",
	})
}

func deposit(context *gin.Context) {
	var amount int = 0
	var status string = "failed"
	var msg string = ""

	input := context.Param("input")
	amount, err := strconv.Atoi(input)
	if err != nil {
		amount = 0
		msg = "操作失敗，輸入有誤！"
	} else {
		if amount <= 0 {
			amount = 0
			msg = "操作失敗，存款金額需大於0元！"
		} else {
			balance += amount
			status = "ok"
		}
	}
	context.JSON(http.StatusOK, gin.H{
		"amount":  amount,
		"status":  status,
		"message": msg,
	})
}

func withdraw(context *gin.Context) {
	var amount int = 0
	var status string = "failed"
	var msg string = ""

	input := context.Param("input")
	amount, err := strconv.Atoi(input)
	if err != nil {
		amount = 0
		msg = "操作失敗，輸入有誤！"
	} else {
		if amount <= 0 {
			amount = 0
			msg = "操作失敗，提款金額需大於0元！"
		} else {
			if balance-amount < 0 {
				amount = 0
				msg = "操作失敗，餘額不足！"
			} else {
				balance -= amount
				status = "ok"
			}
		}
	}
	context.JSON(http.StatusOK, gin.H{
		"amount":  amount,
		"status":  status,
		"message": msg,
	})
}
