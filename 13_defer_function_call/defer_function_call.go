// function calling is done with the help of stack which follows last in, first out.
// defer also follows the same

package main

import "fmt"

func main() {
	for i := 0; i < 7; i++ {
		defer fmt.Println(i)
	}
	defer c()
	fmt.Println("main function execution started")
}

func c() {
	fmt.Println("function c called by main")
	defer b()
	fmt.Println("function c execution completed")
}

func b() {
	fmt.Println("function b is called inside a")
}