package main

import "fmt"

func main() {
	a := 5
	func() {
		fmt.Println("a=", a)
	}()

	fun := appendStr()
	mun := appendStr()
	fmt.Println(fun("World"))
	fmt.Println(fun("Baby"))

	fmt.Println(mun("Gopher"))
	fmt.Println(mun("!"))
}

func appendStr() func(string) string {
	t := "Hello"
	c := func(b string) string {
		t = t + " " + b
		return t
	}
	return c
}
