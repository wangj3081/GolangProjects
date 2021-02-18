package main

import (
	"fmt"
	"time"
)

func main() {
	intch := make(chan int, 1)
	strch := make(chan string, 1)

	intch <- 1
	strch <- "一段内容拉"
	/*
		1、select case 如果未准备好，会阻塞
		2、如果多个case都满足条件就会随机执行一个
		3、case必须为IO操作
		4、可以做超时和退出判断
		5、select 如果对应的是异步处理需要再for循环中使用
		6、select 可以监控 chan 数据流向
	*/
	select {
	case value := <-intch:
		fmt.Println(value)
	case value := <-strch:
		fmt.Println(value)
	case <-time.After(time.Second * 5): //  超时处理
		fmt.Println("超时了。。")
	}
}
