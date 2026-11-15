// 合并 JSON：把多个 JSON 对象文件合并成一个，后面的覆盖前面的。
// 用法: go run mergejson.go a.json b.json
package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Fprintln(os.Stderr, "用法: go run mergejson.go a.json b.json")
		os.Exit(1)
	}
	merged := map[string]any{}
	for _, path := range os.Args[1:] {
		data, err := os.ReadFile(path)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		var obj map[string]any
		if err := json.Unmarshal(data, &obj); err != nil {
			fmt.Fprintf(os.Stderr, "%s 不是 JSON 对象：%v\n", path, err)
			os.Exit(1)
		}
		for k, v := range obj {
			merged[k] = v
		}
	}
	out, _ := json.MarshalIndent(merged, "", "  ")
	fmt.Println(string(out))
}
