package main

import (
	"fmt"
	"time"
)

func main07() {
	go func() {
		fmt.Println("第一个")
	}()
	go func() {
		fmt.Println("第二个")
	}()
	go func() {
		fmt.Println("第三个")
	}()

	time.Sleep(time.Second * 5)
}
