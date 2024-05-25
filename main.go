package main

import "fmt"

func add(x int, y int) int {
	return x + y
}

// Code example from https://golang.org/doc/code.html
func main() {
	fmt.Println(add(42, 13))
}
