package main
/**
遍历通道里数据
 */
import "fmt"

func main() {
	ch1:=make(chan int)
	ch2:=make(chan int)
	go func() {
		for i := 0;i < 10; i++ {
			ch1<-i
		}
		close(ch1)
	}()

	go func() {
		for  {
			i,ok := <-ch1
			if !ok{
				break
			}
			ch2 <- i*i
		}
		close(ch2)
	}()


	for i := range ch2{
		fmt.Println(i)
	}

	fmt.Println("main")
}
