package main

import (
	"fmt"
	"maps"
)

func main() {
	m := make(map[string]int)

	m["a"] = 7
	m["b"] = 13

	fmt.Println("map:", m)

	v := m["a"]
	fmt.Println("v:", v)

	v = m["c"]
	fmt.Println("v:", v)

	fmt.Println("len:", len(m))

	delete(m, "a")
	fmt.Println("map:", m)

	clear(m)
	fmt.Println("map:", m)

	_, exists := m["x"]
	fmt.Println("exi:", exists)

	n := map[string]int{"foo": 1, "bar": 2}
	fmt.Println("map:", n)

	p := map[string]int{"foo": 1, "bar": 2}
	if maps.Equal(n, p) {
		fmt.Println("n === p")
	}
}
