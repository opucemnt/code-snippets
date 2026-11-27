// 简易 grep：在文件里查找包含关键字的行，输出行号和内容。
//
// 运行：go run grep_lite.go <关键字> <文件路径>
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("用法: grep_lite <关键字> <文件路径>")
		os.Exit(1)
	}
	keyword, path := os.Args[1], os.Args[2]
	f, err := os.Open(path)
	if err != nil {
		fmt.Println("打开失败:", err)
		os.Exit(1)
	}
	defer f.Close()
	sc := bufio.NewScanner(f)
	ln := 0
	for sc.Scan() {
		ln++
		if strings.Contains(sc.Text(), keyword) {
			fmt.Printf("%d: %s\n", ln, sc.Text())
		}
	}
}
