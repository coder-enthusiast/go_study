package main

import (
	"fmt"
	"sync"
	"time"
)

/**
并发
 */

//线程计数器
var wg sync.WaitGroup
var k  = 0
func main()  {
	//go hello()
	fmt.Println("main")
	time.Sleep(100)
	for i := 0 ;i<100 ;i++  {
		wg.Add(1)
		go hello()

	}
	//等待线程执行完毕
	wg.Wait()
	fmt.Println(k)


}

func hello()  {
	k++
	wg.Done()
}