package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gjlmotea/golang/wallet_v3/handler"
)

func main() {
	router := gin.Default()
	router.GET("/balance", handler.GetBalance)
	router.POST("/register", handler.Register)
	router.POST("/deposit", handler.Deposit)
	router.POST("/withdraw", handler.Withdraw)

	router.Run(":80")
}
