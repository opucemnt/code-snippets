// 倒序输出文件：把文件的每一行按相反顺序打印。
//
// 运行：go run reverse_lines.go <文件路径>
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("用法: reverse_lines <文件路径>")
		os.Exit(1)
	}
	f, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Println("打开失败:", err)
		os.Exit(1)
	}
	defer f.Close()
	var lines []string
	sc := bufio.NewScanner(f)
	for sc.Scan() {
		lines = append(lines, sc.Text())
	}
	for i := len(lines) - 1; i >= 0; i-- {
		fmt.Println(lines[i])
	}
}
