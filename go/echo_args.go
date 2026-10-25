// 命令行参数回显: 打印程序收到的所有参数及其个数。
//
// 运行: go run echo_args.go a b c
package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args[1:]
	fmt.Printf("共 %d 个参数:\n", len(args))
	for i, a := range args {
		fmt.Printf("  [%d] %s\n", i, a)
	}
}
