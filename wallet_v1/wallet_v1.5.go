package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

var balance int = 0

type Res struct{
	Amount  int `form:"amount" json:"amount"`
	Status  string `form:"status" json:"status"`
	Message string `form:"message" json:"message"`
}

func main() {
	var router = gin.Default()
	router.POST("/deposit/:input", deposit)
	router.POST("/withdraw/:input", withdraw)
	router.GET("/balance/", getBalance)

	router.Run(":80")
}

func getBalance(context *gin.Context) {
	var res = Res{
		Amount: balance,
		Status: "ok",
		Message: "",
	}
	context.JSON(http.StatusOK, res)
}

func deposit(context *gin.Context) {
	var res = Res{
		Amount: 0,
		Status: "failed",
		Message: "",
	}

	input := context.Param("input")
	amount, err := strconv.Atoi(input)
	if err != nil {
		res.Message = "操作失敗，輸入有誤！"
	} else {
		if amount <= 0 {
			res.Message = "操作失敗，存款金額需大於0元！"
		} else {
			balance += amount
			res.Status = "ok"
		}
	}
	res.Amount = balance
	context.JSON(http.StatusOK, res)
}

func withdraw(context *gin.Context) {
	var res = Res{
		Amount: 0,
		Status: "failed",
		Message: "",
	}

	input := context.Param("input")
	amount, err := strconv.Atoi(input)
	if err != nil {
		res.Message = "操作失敗，輸入有誤！"
	} else {
		if amount <= 0 {
			res.Message = "操作失敗，提款金額需大於0元！"
		} else {
			if balance-amount < 0 {
				res.Message = "操作失敗，餘額不足！"
			} else {
				balance -= amount
				res.Status = "ok"
			}
		}
	}
	res.Amount = balance
	context.JSON(http.StatusOK, res)
}
