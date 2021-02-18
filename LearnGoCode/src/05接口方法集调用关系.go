package main

import (
	"fmt"
)

type Animal interface {
	GetName() string
}

type people struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func (p *people) GetName() string {
	return p.Name
}

func main05() {
	// 使用指针对象如果switch不存在 *people 这个类型选项则会输出 is animal,因为 people 实现了 Animal 接口
	var p interface{} = &people{Name: "黑化肥"}

	if o, ok := p.(Animal); ok {
		fmt.Println(o.GetName())
	}

	switch p.(type) {
	case people:
		fmt.Println("is people ")
	case *people:
		fmt.Println("is *people")
	case Animal:
		fmt.Println("is animal")
	default:
		fmt.Println("类型匹配失败")
	}

}
