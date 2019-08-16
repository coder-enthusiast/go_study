package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func a(n int64) int64 {
	sum := int64(0)
	for i := int64(0);i<=n;i++ {
		sum+=i
	}
	return sum
}

func b(star int64, end int64, c chan int64, group *sync.WaitGroup) {
	sum := int64(0)
	for i := star;i<=end;i++ {
		sum+=i
	}
	c <- sum
	group.Done()
}

func sumA(n int64)  {
	a(n)
}

func sumB(n int64, cpu int64, ch chan int64, group *sync.WaitGroup) {
	a:=n/cpu
	data := make(map[int64]int64)
	for i:=int64(0);i<cpu;i++ {
		if i==0 {
			data[0]=a
			continue
		}
		if i==cpu-1 {
			data[(a*i)+1]=n
			continue
		}

		data[(a*i)+1]=a*(i+1)
	}
	for k,v := range data{
		go b(k, v, ch,group)
	}
	group.Wait()
	close(ch)
}

func main() {

	fmt.Println("CPU数量：",runtime.NumCPU())
	//满血
	runtime.GOMAXPROCS(runtime.NumCPU())
	startTime := time.Now().Unix()
	var num int64 = 1000000;

	cpu:=2
	ch := make(chan int64,cpu)
	sy := sync.WaitGroup{}
	sy.Add(cpu)
	sumB(num, int64(cpu),ch,&sy)
	sum := int64(0)
	for i := range ch{
		sum+=i
	}
	fmt.Println(sum)
	fmt.Println("线程数",cpu)
	fmt.Println("计算",num,"的累加")
	fmt.Println("消耗时间",time.Now().Unix()-startTime)


	fmt.Println("-------------------")
	fmt.Println("线程数",1)
	fmt.Println("计算",num,"的累加")
	startTime1 := time.Now().Unix()
	fmt.Println(a(num))
	fmt.Println("消耗时间",time.Now().Unix()-startTime1)
}
