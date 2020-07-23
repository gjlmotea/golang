package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/gjlmotea/golang/wallet_v3/db"
	"net/http"
	"strconv"
)

type Res struct{
	Amount  int `form:"amount" json:"amount"`
	Status  string `form:"status" json:"status"`
	Message string `form:"message" json:"message"`
}

func GetBalance(context *gin.Context) {
	var res = Res{
		Amount: 0,
		Status: "failed",
		Message: "",
	}

	user := context.Query("user")
	amount, valid := db.QueryBalance(user)
	if valid{
		res.Amount = amount
		res.Status = "ok"
	} else {
		res.Message = "使用者帳號不存在！"
	}

	context.JSON(http.StatusOK, res)
}

func Register(context *gin.Context){
	var res = Res{
		Amount: 0,
		Status: "failed",
		Message: "",
	}
	user := context.PostForm("user")
	if user == ""{
		res.Message = "請輸入註冊帳號！"
	} else {
		valid := db.InsertId(user)
		if valid {
			res.Message = "註冊成功。"
			res.Status = "ok"
		} else {
			res.Message = "此帳號已有人使用！"
		}
	}
	context.JSON(http.StatusOK, res)
}

func Deposit(context *gin.Context) {
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
			balance, valid := db.TransactAndQueryBalance(user, amount)
			if valid {
				res.Amount = balance
				res.Status = "ok"
			} else {
				res.Message = "使用者帳號不存在！"
			}
		}
	}
	context.JSON(http.StatusOK, res)
}

func Withdraw(context *gin.Context) {
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
			balance, valid := db.QueryBalance(user)
			if valid {
				if balance-amount < 0{
					res.Message = "操作失敗，餘額不足！"
				} else {
					balance, _ = db.TransactAndQueryBalance(user, -amount)
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
