// 随机密码生成器：默认 16 位，可指定长度。
//
// 运行：go run password.go [长度，默认 16]。
package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"os"
	"strconv"
)

const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789!@#$%^&*"

func main() {
	length := 16
	if len(os.Args) > 1 {
		n, err := strconv.Atoi(os.Args[1])
		if err == nil && n > 0 {
			length = n
		}
	}
	out := make([]byte, length)
	for i := range out {
		idx, err := rand.Int(rand.Reader, big.NewInt(int64(len(charset))))
		if err != nil {
			fmt.Fprintln(os.Stderr, "随机数失败:", err)
			os.Exit(1)
		}
		out[i] = charset[idx.Int64()]
	}
	fmt.Println(string(out))
}
