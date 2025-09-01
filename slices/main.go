package main

import (
	"fmt"
	"slices"
)

func main() {
	var s []string
	fmt.Println("uninit:", s, s == nil, len(s) == 0)

	s = make([]string, 3)
	fmt.Println("emp:", s, "len:", len(s), "cap:", cap(s))

	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println("set:", s)
	fmt.Println("get:", s[2])

	fmt.Println("len:", len(s))
	fmt.Println("cap:", cap(s))

	s = make([]string, 3)
	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	s = append(s, "d")
	s = append(s, "e", "f")
	fmt.Println("apd:", s)
	fmt.Println("len:", len(s))
	fmt.Println("cap:", cap(s))

	c := make([]string, len(s))
	copy(c, s)
	fmt.Println("cpy:", c)

	l := s[2:5]
	fmt.Println("sl:", l)

	l = s[2:]
	fmt.Println("sl:", l)

	t := []string{"g", "h", "i"}
	fmt.Println("dcl:", t)

	u := []string{"g", "h", "i"}
	if slices.Equal(t, u) {
		fmt.Println("t === u")
	}

	vec := make([][]int, 3)
	fmt.Println("vec:", vec)
	for i := range 3 {
		j := i + 1
		vec[i] = make([]int, j)
		for k := range j {
			vec[i][k] = i + k
		}
	}
	fmt.Println("vec:", vec)
}
