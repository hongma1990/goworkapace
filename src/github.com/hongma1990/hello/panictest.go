package main

import (
	"fmt"
	"runtime"
	"strconv"
	"time"
)

func main() {
	fmt.Println("hi")
	defer func() {
		v := recover()
		value, ok := v.(string)
		fmt.Printf("in defer,isOk:%s,value:%s", strconv.FormatBool(ok), value)
	}()
	//panic("bye!")
	fmt.Println("no end")

	fmt.Println("hi!")

	go func() {
		defer func() {
			v := recover()
			fmt.Printf("in def111er,isOk:%s,value:%s", v)
		}()
		time.Sleep(time.Second)
		panic(123)
	}()

	/*for {
	    time.Sleep(time.Second)
	}*/

	fmt.Println("goruntine:", runtime.NumGoroutine())
}
