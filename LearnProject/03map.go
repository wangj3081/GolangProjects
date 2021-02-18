package main

import (
	"fmt"
)

func main() {
	map_val := map[int]string{1001: "哈哈哈", 1002: "呵呵呵", 1003: "哇哇哇"}
	map_val[1004] = "嘻嘻嘻"
	fmt.Println(map_val)
	for key, value := range map_val {
		fmt.Printf("%d\t%s\n", key, value)
	}
	// 获取对应键是否存在
	_, ok := map_val[1004]
	fmt.Println(ok)
}
