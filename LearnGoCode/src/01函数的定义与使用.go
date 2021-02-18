package main

import "fmt"

func Demo(i int) {
	var arr [2]int
	defer func() {
		// 设置拦截错误信息
		err := recover()
		// 产生 panic 打印异常日志
		if err != nil {
			fmt.Println(err)
		}
	}()
	// 如果i的数字超过定义的数组长度，会报异常
	arr[i] = i
}

/**
有返回值有入参的函数
*/
func SayHello(name string) string {
	return "【Echo】Hello：" + name
}

func main() {
	Demo(1)
	fmt.Println("程序继续向下执行")
	Demo(2)
	fmt.Println("程序继续向下执行")
	fmt.Println(SayHello("小明"))
}
