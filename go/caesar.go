// 凯撒密码：对字母按指定位数移位加密/解密（只处理英文字母）。
//
// 运行：go run caesar.go <文本> <位数> [d 解密，默认加密]
package main

import (
	"fmt"
	"os"
	"strconv"
)

func shiftChar(ch rune, n int) rune {
	switch {
	case ch >= 'a' && ch <= 'z':
		return 'a' + (ch-'a'+rune(n)+26)%26
	case ch >= 'A' && ch <= 'Z':
		return 'A' + (ch-'A'+rune(n)+26)%26
	default:
		return ch
	}
}

func main() {
	if len(os.Args) < 3 {
		fmt.Println("用法: caesar <文本> <位数> [d 解密]")
		os.Exit(1)
	}
	n, err := strconv.Atoi(os.Args[2])
	if err != nil {
		fmt.Println("位数必须是整数")
		os.Exit(1)
	}
	if len(os.Args) > 3 && os.Args[3] == "d" {
		n = -n
	}
	out := []rune(os.Args[1])
	for i, ch := range out {
		out[i] = shiftChar(ch, n)
	}
	fmt.Println(string(out))
}
