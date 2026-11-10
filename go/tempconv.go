// 温度换算：摄氏度和华氏度互转。
//
// 运行：go run tempconv.go 25C 或 go run tempconv.go 77F。
package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("用法: tempconv <数值C|数值F>，例如 25C 或 77F")
		os.Exit(1)
	}
	arg := os.Args[1]
	unit := arg[len(arg)-1]
	num, err := strconv.ParseFloat(arg[:len(arg)-1], 64)
	if err != nil || (unit != 'C' && unit != 'F') {
		fmt.Println("参数格式错误，示例: 25C 或 77F")
		os.Exit(1)
	}
	if unit == 'C' {
		fmt.Printf("%.1fC = %.1fF\n", num, num*9/5+32)
	} else {
		fmt.Printf("%.1fF = %.1fC\n", num, (num-32)*5/9)
	}
}
