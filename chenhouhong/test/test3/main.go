package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Printf("%s\n", runtime.Version())
	const Ln2 = 0.693147180559945309417232121458176568075500134360255254120680009
	const Log2E = 1 / Ln2
	const Billion = 1e9
	const hardEight = (1 << 100) >> 97
	fmt.Println(Log2E)

	const beef, two, c = "eat", 2, "veg"
	const (
		Monday, Tuesday, Wednesday = 1, 2, 3
		Thursday, Friday, Saturday = 4, 5, 6
	)
	const (
		a = iota
		b = iota
		cc = iota
	)
	const (
		first, second = iota+1,iota+2
		three,four
		five,six
	)
	fmt.Println(five)
}
