// 逐行读取文件并打印行号。
//
// 运行: go run read_lines.go <文件路径>
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("用法: read_lines <文件路径>")
		os.Exit(1)
	}
	f, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Println("打开文件失败:", err)
		os.Exit(1)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	n := 0
	for scanner.Scan() {
		n++
		fmt.Printf("%4d: %s\n", n, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		fmt.Println("读取失败:", err)
	}
}
