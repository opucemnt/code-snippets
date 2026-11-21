// 计算文件的 MD5 值（流式读取，大文件也不怕）。
//
// 运行：go run md5sum.go <文件路径> [文件路径 ...]
package main

import (
	"crypto/md5"
	"fmt"
	"io"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("用法: md5sum <文件路径> [文件路径 ...]")
		os.Exit(1)
	}
	for _, path := range os.Args[1:] {
		f, err := os.Open(path)
		if err != nil {
			fmt.Printf("读取失败 %s: %v\n", path, err)
			continue
		}
		h := md5.New()
		if _, err := io.Copy(h, f); err != nil {
			fmt.Printf("读取失败 %s: %v\n", path, err)
			f.Close()
			continue
		}
		f.Close()
		fmt.Printf("%x  %s\n", h.Sum(nil), path)
	}
}
