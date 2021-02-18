package main

import (
	"fmt"
	"time"
)

func main08() {
	// chan 定义方式, 有缓冲表示可接收 N 个值，属于异步操作;
	// 无缓冲需要发送和接收方都准备好，才可以完成，在接收时会阻塞等待，属于同步操作
	//var ch  = make(chan 数据类型, 缓冲大小)
	//var ch  = make(chan 数据类型)
	//ch := make(chan int,6)
	//fmt.Println(unsafe.Sizeof(ch))

	// 单向chan
	//ch := make(chan <- int) // 发送
	//ch := make(<-chan int) // 接收
	// 异步的channel必须在 gorouting 中接收
	ch := make(chan int)
	go func() {
		value := <-ch
		fmt.Println(value)
	}()
	ch <- 1
	time.Sleep(time.Second * 2)
}
