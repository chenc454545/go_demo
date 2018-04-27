package main

import (
	"fmt"
)

func main() {
	p := newInt()
	fmt.Println(p)
	*p = 100
	fmt.Println(*p)
}

func newInt() *int {
	return new(int)
}
