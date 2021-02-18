package main

import (
	"fmt"
)

func main02() {
	slice_val := []int{1, 2, 3, 4, 5}
	fmt.Println(slice_val)
	// 浅拷贝
	slice_val2 := slice_val
	slice_val2[2] = 2
	fmt.Println(slice_val2)
	fmt.Println(slice_val)
	// 深拷贝,需要开辟新的内存空间
	slice_val3 := make([]int, len(slice_val))
	copy(slice_val3, slice_val)
	slice_val3[1] = 444
	fmt.Println(slice_val)
	fmt.Println(slice_val3)

	slice_b := []int{}
	fmt.Printf("%T", slice_b)

}
