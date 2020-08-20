package main

import (
	"log"
	"math/rand"
	"sync"
	"time"
)

func main() {
	log.SetFlags(5)

	outC := make(chan int)
	//errC := make(chan error)
	//finC := make(chan struct{})
	wg := sync.WaitGroup{}
	wg.Add(100)
	for i := 0; i < 100; i++{
		go job(outC, &wg)
	}

	for i := 0; i < 100; i++{
		select {
		case <- outC:
			log.Println("...Get outC 1")
		case <- outC:
			log.Println("...Get outC 2")
		}
	}
}

func job(outC chan int, wg *sync.WaitGroup)  {
	rand.Seed(time.Now().UnixNano())
	randNum := rand.Int31n(1000)
	time.Sleep(time.Duration(randNum) * time.Millisecond)
	log.Println("finished", randNum)
	outC <- int(randNum)
}