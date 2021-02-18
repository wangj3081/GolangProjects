package main

import "fmt"

type Person1 struct {
	Name string
	Age  int
}

type Student1 struct {
	Person1
	Score int
}

// Person 结构体的函数
func (p *Person1) SayHello() {
	fmt.Println(*p)
}

func (stu *Student1) SayHello() {
	fmt.Println(*stu)
}

func main() {
	stu := Student1{Person1{Name: "小花"}, 200}
	stu.Person1.SayHello()
	stu.SayHello()
}
