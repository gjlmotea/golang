package main

import (
	"log"
	"time"
)

func main() {
	log.SetFlags(5)

	ch := make(chan int)

	go autoSend(ch)

	for i := 0; i < 5; i++{
		time.Sleep(1000)
		select {
		case <-ch:
			log.Println("channel 01 Get")
		case <-ch:
			log.Println("channel 02 Get")
		default:
			log.Println("This is Default：代表接收不到，可能還在等待send")
		}
	}
}

func autoSend(ch chan int)  {
	i := 0
	for{
		i++
		ch <- i
		log.Println("send", i)
	}
}