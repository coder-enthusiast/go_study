package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

/**
go 线程调度
 */
var gw sync.WaitGroup
func a()  {
	for i := 0;i<10;i++ {
		fmt.Println("A：",i)
		gw.Done()
		time.Sleep(1000*time.Millisecond)
	}
}

func b()  {
	for i := 0;i<10;i++ {
		fmt.Println("B：",i)
		gw.Done()
		time.Sleep(1000*time.Millisecond)
	}
}


func main() {
	//获取当前时间戳
	startTime:=time.Now().Second()
	//控制物理CPU
	runtime.GOMAXPROCS(2)

	gw.Add(20)
	go a()
	go b()
	gw.Wait()
	endTime:=time.Now().Second()
	fmt.Println("startTime:",startTime)
	fmt.Println("endTime:",endTime)

	fmt.Println(endTime-startTime)
}
