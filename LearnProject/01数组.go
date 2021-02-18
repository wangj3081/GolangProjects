package main

import "fmt"

func main01() {
	arr := [...]int{1, 2, 3, 4, 5}
	var b [5]int
	fmt.Println(arr)
	fmt.Printf("%T", b)
	fmt.Printf("%T", arr)
}
