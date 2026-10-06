// WaitGroup 并发示例: 5 个 goroutine 并发执行任务, 全部完成后汇总。
//
// 运行: go run wg_demo.go
package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(id int, wg *sync.WaitGroup) {
	defer wg.Done()
	time.Sleep(time.Duration(id) * 100 * time.Millisecond)
	fmt.Printf("worker %d done\n", id)
}

func main() {
	var wg sync.WaitGroup
	for i := 1; i <= 5; i++ {
		wg.Add(1)
		go worker(i, &wg)
	}
	wg.Wait()
	fmt.Println("all workers done")
}
