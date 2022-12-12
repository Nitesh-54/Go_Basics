package main

import "fmt"

func main() {

	num := 67
	if num < 7 {
		fmt.Println(num, "is less than 7")
	} else if num > 7 {
		fmt.Println(num, "is greater than 7")
	} else {
		fmt.Println("entered number is equal to 7")
	}
}
