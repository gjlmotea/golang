package main

import (
	"fmt"
	"log"
	"math/rand"
	"sync"
	"time"
)

func main() {
	log.SetFlags(5)

	outC := make(chan int)
	//finC := make(chan struct{})
	wg := sync.WaitGroup{}
	wg.Add(10)
	for i := 0; i < 10; i++{
		go job(outC, &wg)
	}

	for i := 0; i < 10; i++{
		select {
		case c := <- outC:
			log.Println("...Get outC", c)
		}
	}
/*
	for i := 0; i < 10; i++{
		c, access := <- outC
		fmt.Println(c, access)
	}
//access偵測channel是否存在
 */
}

func job(outC chan int, wg *sync.WaitGroup)  {
	defer wg.Done()
	rand.Seed(time.Now().UnixNano())
	randNum := rand.Int31n(1000)
	time.Sleep(time.Duration(randNum) * time.Millisecond)
	log.Println("finished", randNum)

	outC <- int(randNum)
	close(outC)
}

/*
channel只傳出一次就被關閉了
2020/08/05 10:59:30.822739 finished 119
2020/08/05 10:59:30.822913 ...Get outC 119
2020/08/05 10:59:30.822917 ...Get outC 0
2020/08/05 10:59:30.822920 ...Get outC 0
2020/08/05 10:59:30.822922 ...Get outC 0
2020/08/05 10:59:30.822925 ...Get outC 0
2020/08/05 10:59:30.822927 ...Get outC 0
2020/08/05 10:59:30.822930 ...Get outC 0
2020/08/05 10:59:30.822931 ...Get outC 0
2020/08/05 10:59:30.822933 ...Get outC 0
2020/08/05 10:59:30.822934 ...Get outC 0
*/