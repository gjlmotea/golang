package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	c := make(chan int, 1)
	go getSomething(c)

	select {
		case i := <- c:
			fmt.Println("收到C", i)
		case <- time.After(time.Microsecond * 1000):
			fmt.Println("時間超時")
	}
}

func getSomething(c chan int)  {
	rand.Seed(time.Now().UnixNano())
	randNum := rand.Int() % 1000
	fmt.Println("Get something takes", randNum, "ms")
	time.Sleep(time.Microsecond *time.Duration(randNum))
	c <- 1000
}