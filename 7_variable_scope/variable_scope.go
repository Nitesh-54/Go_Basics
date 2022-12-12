package main

import "fmt"

var a = 9

func demo() {
	var a = 6
	fmt.Println("in demo", a)
}

func main() {
	demo()
	fmt.Println("in main", a)
}
