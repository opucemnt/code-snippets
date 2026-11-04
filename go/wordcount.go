// 词频统计：读文件或标准输入，输出行数、单词数、字符数。
//
// 运行：go run wordcount.go [文件路径]，不给路径则读标准输入。
package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	var r io.Reader = os.Stdin
	if len(os.Args) > 1 {
		f, err := os.Open(os.Args[1])
		if err != nil {
			fmt.Fprintln(os.Stderr, "打开失败:", err)
			os.Exit(1)
		}
		defer f.Close()
		r = f
	}
	sc := bufio.NewScanner(r)
	lines, words, chars := 0, 0, 0
	for sc.Scan() {
		line := sc.Text()
		lines++
		chars += len([]rune(line)) + 1
		words += len(strings.Fields(line))
	}
	if err := sc.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "读取失败:", err)
		os.Exit(1)
	}
	fmt.Printf("行数: %d  单词: %d  字符: %d\n", lines, words, chars)
}
