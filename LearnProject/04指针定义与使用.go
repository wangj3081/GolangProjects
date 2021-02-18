package main

import (
	"fmt"
)

func main04() {
	// 定义指针，默认值为nil,直接赋值是没用的，因为没有开辟内存空间
	/*var p *int
	*p = 123
	fmt.Println(p)*/
	// 如果定义好的指针变量，需要赋值则要先开辟好内存
	var p *int
	p = new(int)
	*p = 123
	fmt.Println(*p)
	sliceVal := make([]int, 2)
	sliceVal[1] = 23
	fmt.Println(sliceVal)
}
