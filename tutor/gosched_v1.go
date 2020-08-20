package main

import (
	"fmt"
	"time"
	"runtime"
)

func say(s string) {
	for i := 0; i < 2; i++ {
		runtime.Gosched()
		fmt.Println(s)
	}
}

func main() {
	runtime.GOMAXPROCS(1)
	say("hello")
	go say("world")
	time.Sleep(100000)
}
/*
result:
hello
hello
world
world
 */