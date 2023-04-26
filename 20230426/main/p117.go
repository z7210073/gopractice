package main

import "fmt"

func test(n int) int {
	//斐波那契
	if n == 1 || n == 2 {
		return 1
	} else {
		return test(n-1) + test(n-2)
	}
}
func num(n int) int {
	if n == 1 {
		return 3
	} else {
		return 2*num(n-1) + 1
	}
}

func main() {

	fmt.Println(num(2))
	//num(2)

}
