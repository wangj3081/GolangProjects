package main

import (
	"fmt"
	"sync"
)

func main() {
	mute := sync.Mutex{}
	wg := sync.WaitGroup{}
	count := 0
	for i := 1; i <= 100; i++ {
		wg.Add(1)
		go func() {
			mute.Lock()
			defer mute.Unlock() // 最后解锁
			defer wg.Done()     // 业务处理完成后做确认操作
			// 业务处理
			count++
		}()
	}
	wg.Wait() // 等待任务完成
	fmt.Println(count)
}
