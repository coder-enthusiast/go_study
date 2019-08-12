package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int,2)
	for i := 0; i < 10; i++ {
		select {
		case x := <-ch:
			fmt.Println(x)
		case ch <- i:
		}
	}

	time.Sleep(1000*time.Second)
}