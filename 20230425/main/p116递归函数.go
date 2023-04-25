package main

import "fmt"

func test(n int) {

	if n < -10 {
		fmt.Println(123)
	} else {
		n--
		test(n)
	}
}
func main() {

	test(10)
}
