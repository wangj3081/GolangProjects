package main

import (
	"fmt"
	"time"
)

func main09() {
	// 空结构体不占用空间，做通知作用
	ch := make(chan struct{})

	go func() {
		time.Sleep(time.Second * 2)
		fmt.Println("携程执行完成。。。")
		close(ch)
	}()
	fmt.Println("主线程开始执行")
	<-ch
	fmt.Println("主线程执行完成。。。")

}
