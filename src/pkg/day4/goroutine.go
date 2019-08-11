package main

import (
	"fmt"
	"time"
)

/**
并发
 */
func main()  {
	go hello()
	fmt.Println("main")
	time.Sleep(100)
}

func hello()  {
	fmt.Println("hello")
}