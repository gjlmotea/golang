package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	// 使用一個邏輯處理器給調度器用
	runtime.GOMAXPROCS(1)
	//可以改2試試

	// wg + 2 表示要等待2個goroutine完成
	var wg sync.WaitGroup
	wg.Add(2)

	fmt.Println("start goroutines")

	// 創建goroutine
	go func() {
		defer wg.Done()
		for count := 0; count < 3; count++ {
			for i := 0; i < 10; i++ {
				fmt.Printf(".")
			}
		}
	}()

	// 創建另一個goroutine
	go func() {
		defer wg.Done()
		for count := 0; count < 3; count++ {
			for i := 0; i < 10; i++{
				fmt.Printf("=")
			}
		}
	}()

	fmt.Println("waiting to finish")
	wg.Wait()

	fmt.Println("\nfinish Program")
}