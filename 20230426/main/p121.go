package main

import "fmt"

// 支持多个参数
// arg... 是切片
func calc2(args ...int) {
	var sum int
	for i := 0; i < len(args); i++ {
		sum += args[i]

	}
	fmt.Println(sum)
}

func main() {
	calc2(1, 2, 66, 7)
}
