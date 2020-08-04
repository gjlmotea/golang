package main

import (
	"log"
	"time"
)

func main()  {
	//c := new(chan int)
	//var c chan int

	c := make(chan int)
	go test(c)
	c <- 100
	c <- 10
	c <- 1
	time.Sleep(1000)
	return
}

func test(c chan int){
	i := <- c
	log.Println(i)
	j := <- c
	log.Println(j)
	k := <- c
	log.Println(k)
}