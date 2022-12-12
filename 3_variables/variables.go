package main

import "fmt"

func main() {

	var num int
	//default value is 0

	var num1 int = 1
	//type is optional when a value is assigned

	var num2 int = 2

	var num3, num4 int
	//variables of the same type can be declared on the same line

	num3, num4 = 3, 4
	//Assigning values to defined variables of same type

	var num5, num6 = 5, 6

	num7 := 7
	//colon takes care of creation and assignment

	var result = num + num1 + num2 + num3 + num4 + num5 + num6 + num7

	fmt.Print(result)

}
