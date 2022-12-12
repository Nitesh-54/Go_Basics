package main

import "fmt"

func main() {
	num := 1
	switch num {
	case 1:
		fmt.Println("entered value is 1")
	case 2:
		fmt.Println("entered value is 2")
	default:
		fmt.Println("entered value is bigger than 2")
	}
}
