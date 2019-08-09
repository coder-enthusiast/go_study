package main

import "fmt"

func f1() int {
	x := 5
	defer func() {
		x++
	}()
	return x
}

func f2() (x int) {
	defer func() {
		x++
	}()
	return 5
}

func f3() (y int) {
	x := 5
	defer func() {
		x++
	}()
	return x
}
func f4() (x int) {
	defer func(x int) {
		x++
	}(x)
	return 5
}
func main() {
	/*fmt.Println(f1())
	fmt.Println(f2())
	fmt.Println(f3())
	fmt.Println(f4())*/

	fmt.Println(s1())
	fmt.Println(s2())
	fmt.Println(s3())
	fmt.Println(s4())
}

func s1() int {
	defer fmt.Println("s1")
	return 1
}

func s2() int {
	defer fmt.Println("s2")
	return 2
}

func s3() int {
	defer fmt.Println("s3")
	return 3
}

func s4() int {
	defer fmt.Println("s4")
	return 4
}