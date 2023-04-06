package main

func main() {
	//go 语言没有while  , do while
	//通过 for 实现while
	var a int
	for {
		if a == 1 {
			break
		}
		a++
	}
	//通过 for 实现do while

	a = 0
	for {
		a++
		if a == 1 {
			break
		}
	}
}
