package main

import "fmt"

func main() {
	var a add = func(a, b int) int {
		return a + b
	}
	s := a(5, 6)
	fmt.Println(s)
}

type add func(a int, b int) int
