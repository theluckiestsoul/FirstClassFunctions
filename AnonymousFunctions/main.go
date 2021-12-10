package main

import "fmt"

func main() {
	a := func() {
		fmt.Println("Hello World from first class function")
	}

	func(d int){
		fmt.Println(d)
	}(1)

	func(){
		fmt.Println("Anonymous function")
	}()

	a()
	fmt.Printf("%T",a)
}


