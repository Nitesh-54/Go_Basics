package main

import "fmt"

func main() {
	// for {
	// 	fmt.Print("alex")
	// }
	// infinite loop
	i := 1
	for i <= 5 {
		fmt.Print("alex", i, "\n")
		i++
	}
	fmt.Print("\n")
	for i := 1; i < 5; i++ {
		fmt.Print("alex", i, "\n")
	}

}
