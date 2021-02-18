package main

import (
	"fmt"
)

// 定义接口类型
type Humaner interface {
	SayHi(str string) string
}

type Teacher struct {
	Name string
	Age  int
}

type School struct {
	Name string
}

func (t *Teacher) SayHi(str string) string {
	fmt.Println(*t)
	return t.Name + ": Hi、" + str
}

func (school *School) SayHi(str string) string {
	fmt.Println(*school)
	return str + ": Hi、" + school.Name
}

// 多态实现, 将接口类型作为参数
func SayHi(h Humaner, str string) string {
	return h.SayHi(str)
}

func main01() {
	t := Teacher{"安西", 27}
	// 创建接口类型变量
	var h Humaner
	// 将对象的地址赋值给接口类型
	h = &t
	// 通过接口调用
	fmt.Println(h.SayHi("樱木"))
}

// 多态函数实现
func main04() {
	t := Teacher{"三井", 28}
	fmt.Println(SayHi(&t, "樱木"))
	s := School{Name: "湘北高中"}
	fmt.Println(SayHi(&s, "樱木"))
	// 判断是否实现对应接口
	var teacher interface{} = &Teacher{"宫城", 22}
	if te, ok := teacher.(Humaner); ok {
		fmt.Println(te.SayHi("赤木"))
	}
}
