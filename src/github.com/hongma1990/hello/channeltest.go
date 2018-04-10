package main

import (
	"fmt"
	"runtime"
	// "reflect"
)

func sum(a []int, c chan int) {
	total := 0
	for _, v := range a {
		total += v
	}
	c <- total
}

func fibonacci(n int, c chan int) {
	x, y := 1, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}
	close(c)
}

func main() {
	a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	c := make(chan int, 2)
	go sum(a[:len(a)/2], c)
	go sum(a[len(a)/2:], c)
	x, y := <-c, <-c
	fmt.Printf("x: %d, y:%d", x, y)

	fc := make(chan int, 10)
	go fibonacci(cap(fc), fc)
	go func() {
		for v := range fc {
			fmt.Println(v)
		}
	}()
	fmt.Println("\nnumbers of goroutine:", runtime.NumGoroutine())
}
