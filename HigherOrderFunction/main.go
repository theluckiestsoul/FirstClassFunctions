package main

import "fmt"

func simple(a func(b, c int) int) {
	fmt.Println(a(60, 7))
}

func main() {
	f := func(a, b int) int {
		return a + b
	}

	simple(f)
	fmt.Println(simple2()(12, 34))
}

func simple2() func(a, b int) int {
	return func(a, b int) int {
		return a + b
	}
}
