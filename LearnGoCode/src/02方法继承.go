package main

import "fmt"

/**
Go 语言的继承是通过在结构体内设置对应的匿名对象
*/

type Person struct {
	Name string
	Age  int
}

type Student struct {
	Person
	Score int
}

// Person 结构体的函数
func (p Person) SayHello() {
	fmt.Println(p)
}

func main() {
	p := Person{"小明", 10}
	p.SayHello()
	// 可以直接调用父类的函数，但是只能使用到父类的属性
	stu := Student{Person{"小红", 18}, 100}
	stu.SayHello()
}
