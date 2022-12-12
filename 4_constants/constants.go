package main

import "fmt"

func main() {

	var num int = 1
	num = 2

	fmt.Print("variable value is ", num)

	const num_c = 3 //cannot change value once declared

	fmt.Print("\nconstant value is ", num_c)

}
