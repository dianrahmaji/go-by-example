package main

import "fmt"

func add(a int, b int) int {
	return a + b
}

func veryAdd(a, b, c int) int {
	return a + b + c
}

func main() {
	res := add(1, 2)
	fmt.Println("1+2 =", res)

	res = veryAdd(1, 2, 3)
	fmt.Println("1+2+3 =", res)
}
