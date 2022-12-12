package main

import "fmt"

func main() {
	num1 := 1
	num2 := 2
	result1, result2 := calc(num1, num2)
	fmt.Println(result1, "\t", result2)
}

func calc(x int, y int) (int, int) {
	// func add(x, y int) (out1, out2 int) {
	var out1 = x + y
	var out2 = x - y
	return out1, out2 //return
}
