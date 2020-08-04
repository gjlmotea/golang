package main

import (
	"fmt"
	"time"
)

func main()  {
	c := make(chan int)	//num: buffer size
	go printer(c)
	c <- 10
	c <- 100
	//c <- 1000		//it will block
	time.Sleep(1000000)
}

func printer(c chan int){
	fmt.Println("hi")
	i := <- c
	fmt.Println(i)
}