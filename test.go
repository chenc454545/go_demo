package main

import (
	"fmt"
)

func main() {
	p := newInt()
	q := newInt()
	fmt.Println(*p,*q, p, q)
	fmt.Println(p == q)
	var pp *int
	*p = 100
	*q = 100
	fmt.Println(*p,*q, p, q)
	fmt.Println(p == q)
	fmt.Println(pp == p)
}

func newInt() *int {
	return new(int)
}
