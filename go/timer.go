// 命令行倒计时：每秒打印剩余秒数，到 0 提示完成。
//
// 运行：go run timer.go [秒数，默认 10]。
package main

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

func main() {
	secs := 10
	if len(os.Args) > 1 {
		n, err := strconv.Atoi(os.Args[1])
		if err == nil && n > 0 {
			secs = n
		}
	}
	for i := secs; i > 0; i-- {
		fmt.Printf("\r剩余 %d 秒 ", i)
		time.Sleep(time.Second)
	}
	fmt.Println("\r时间到！")
}
