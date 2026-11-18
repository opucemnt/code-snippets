// 词频统计：读一个文本文件，按出现次数倒序列出前 10 个词。
// 用法: go run wordfreq.go <文件>
package main

import (
	"fmt"
	"os"
	"sort"
	"strings"
)

type pair struct {
	word  string
	count int
}

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintln(os.Stderr, "用法: go run wordfreq.go <文件>")
		os.Exit(1)
	}
	data, err := os.ReadFile(os.Args[1])
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	counts := map[string]int{}
	for _, w := range strings.Fields(string(data)) {
		counts[strings.ToLower(w)]++
	}
	pairs := make([]pair, 0, len(counts))
	for w, c := range counts {
		pairs = append(pairs, pair{w, c})
	}
	sort.Slice(pairs, func(i, j int) bool { return pairs[i].count > pairs[j].count })
	for i, p := range pairs {
		if i >= 10 {
			break
		}
		fmt.Printf("%-15s %d\n", p.word, p.count)
	}
}
