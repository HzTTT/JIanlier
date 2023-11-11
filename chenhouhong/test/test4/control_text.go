package main

import "fmt"

func main() {
	//for range
	x := [...]int{1,2,3}
	for index, v := range x {
		fmt.Printf("index: %v\n", index)
		fmt.Printf("v: %v\n", v)
	}


	//
}