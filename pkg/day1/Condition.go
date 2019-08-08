/**
条件语句
 */
package main

import "fmt"

func main()  {
	a:=false
	if a{
		fmt.Println(a)
	}else{
		fmt.Println(!a)
	}

	b:="bb"
	switch b {

	case "bb":
		fmt.Println("bb")
	case "aa":
		fmt.Println("aa")
	default:
		fmt.Println("default")
	}

	var c1, c2, c3 chan int
	var i1, i2 int
	select {
	case i1 = <-c1:
		fmt.Printf("received ", i1, " from c1\n")
	case c2 <- i2:
		fmt.Printf("sent ", i2, " to c2\n")
	case i3, ok := (<-c3):  // same as: i3, ok := <-c3
		if ok {
			fmt.Printf("received ", i3, " from c3\n")
		} else {
			fmt.Printf("c3 is closed\n")
		}
	default:
		fmt.Printf("no communication\n")
	}
}
