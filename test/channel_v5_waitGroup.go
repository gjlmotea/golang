package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main()  {
	wg := sync.WaitGroup{}
	for i := 0; i < 10; i++{
		wg.Add(1)
		go dosomething(i , &wg)
	}
	wg.Wait()
	fmt.Println("end")
}

func dosomething(i int, wg *sync.WaitGroup){
	rand.Seed(time.Now().UnixNano())
	randNum := rand.Int() % 2000
	time.Sleep(time.Duration(randNum) * time.Millisecond)
	fmt.Println("task", i, "takes:", randNum, "ms")
	wg.Done()
}