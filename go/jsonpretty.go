// JSON 格式化：读标准输入的 JSON，输出带缩进的版本。
//
// 运行：cat data.json | go run jsonpretty.go。
package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"os"
)

func main() {
	data, err := io.ReadAll(os.Stdin)
	if err != nil {
		fmt.Fprintln(os.Stderr, "读取失败:", err)
		os.Exit(1)
	}
	var buf bytes.Buffer
	if err := json.Indent(&buf, data, "", "  "); err != nil {
		fmt.Fprintln(os.Stderr, "不是合法 JSON:", err)
		os.Exit(1)
	}
	fmt.Println(buf.String())
}
