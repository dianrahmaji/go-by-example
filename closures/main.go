package main

import "fmt"

func intSeq() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func main() {
	nextInt := intSeq()

	for range 10 {
		fmt.Println(nextInt())
	}

	nextInt = intSeq()
	fmt.Println(nextInt())
}
