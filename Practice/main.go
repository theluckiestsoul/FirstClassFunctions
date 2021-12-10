package main

import "fmt"

type student struct {
	firstName string
	lastName  string
	grade     string
	country   string
}

func filter(s []student, f func(student) bool) []student {
	var r []student
	for _, v := range s {
		if f(v) {
			r = append(r, v)
		}
	}
	return r
}

func main() {
	s1 := student{
		firstName: "Kiran",
		lastName:  "Mohanty",
		grade:     "A",
		country:   "India",
	}
	s2 := student{
		firstName: "Smauel",
		lastName:  "Johnson",
		grade:     "B",
		country:   "USA",
	}

	s := []student{s1, s2}
	f := filter(s, func(s student) bool {
		return s.grade == "B"
	})

	fmt.Println(f)

	a:=[]int{5,6,7,8,9}
	r:=iMap(a,func(i int) int {
		return i*5
	})
	fmt.Println(r)
}

func iMap(s []int, f func(int) int) []int {
	var r []int
	for _, v := range s {
		r = append(r, f(v))
	}
	return r
}
