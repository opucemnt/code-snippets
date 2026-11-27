// 快速排序：取中间元素做基准，递归分治。
// 用法: go run quicksort.go
package main

import (
	"fmt"
	"math/rand"
)

func quicksort(items []int) []int {
	if len(items) <= 1 {
		return items
	}
	pivot := items[len(items)/2]
	var left, mid, right []int
	for _, x := range items {
		switch {
		case x < pivot:
			left = append(left, x)
		case x == pivot:
			mid = append(mid, x)
		default:
			right = append(right, x)
		}
	}
	return append(append(quicksort(left), mid...), quicksort(right)...)
}

func main() {
	data := make([]int, 12)
	for i := range data {
		data[i] = rand.Intn(100)
	}
	fmt.Println("before:", data)
	fmt.Println("after: ", quicksort(data))
}
