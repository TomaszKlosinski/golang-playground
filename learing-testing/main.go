package main

import "fmt"

func main() {
	sum := add(1, 2)
	fmt.Println(sum)
}

func add(l, r int) int {
	return l + r
}
