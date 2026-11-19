// ROT13：字母表轮换 13 位，加一次是加密，再加一次就解回来了。
// 用法: go run rot13.go <文本>
package main

import (
	"fmt"
	"os"
	"strings"
)

func rot13(s string) string {
	var b strings.Builder
	for _, ch := range s {
		switch {
		case ch >= 'a' && ch <= 'z':
			b.WriteRune('a' + (ch-'a'+13)%26)
		case ch >= 'A' && ch <= 'Z':
			b.WriteRune('A' + (ch-'A'+13)%26)
		default:
			b.WriteRune(ch)
		}
	}
	return b.String()
}

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintln(os.Stderr, "用法: go run rot13.go <文本>")
		os.Exit(1)
	}
	fmt.Println(rot13(strings.Join(os.Args[1:], " ")))
}
