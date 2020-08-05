package main

import (
	"errors"
	"log"
	"math/rand"
	"sync"
	"time"
)

func main() {
	log.SetFlags(5)

	outC := make(chan int)
	errC := make(chan error)
	finC := make(chan struct{})

	goNum := 10
	wg := sync.WaitGroup{}
	wg.Add(goNum)
	for i := 0; i < goNum; i++{
		go job(outC, errC, &wg)
	}

	go waitAll(finC,&wg)

	LOOP:
	for i := 0; i < goNum; i++{
		select {
		case c := <- outC:
			log.Println("...Get outC", c)
		case err := <- errC:
			log.Println("==============ERROR in", err)
			break LOOP	//跳出到LOOP不再執行
		case f := <- finC:
			log.Println("f", f)
			break LOOP
		}
	}
}

func waitAll(finC chan struct{}, wg *sync.WaitGroup)  {
	wg.Wait()
	log.Println("關閉channel")
	close(finC)
}

func job(outC chan int, errC chan error, wg *sync.WaitGroup)  {
	defer wg.Done()
	rand.Seed(time.Now().UnixNano())
	randNum := rand.Int31n(1000)
	time.Sleep(time.Duration(randNum) * time.Millisecond)
	log.Println("finished", randNum)

	if randNum >= 500 && randNum <= 600{
		errC <- errors.New("出現數字 500~600")
		return
	}
	outC <- int(randNum)
}