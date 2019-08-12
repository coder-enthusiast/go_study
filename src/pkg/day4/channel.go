package main

import (
	"fmt"
	"time"
)

/**
通道
 */
func main() {
	var ch chan int
	ch = make(chan int)
	go recv(ch)
	time.Sleep(time.Second)
	//会一直堵塞子线程 直至有线程接收 否则造成死锁
	ch  <- 10

	fmt.Println("main")
}

//接收者
func recv(ch chan int)  {
	for i:= range ch {
		fmt.Println("接收者", i)

	}
}