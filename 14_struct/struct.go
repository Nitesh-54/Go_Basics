// struct is a user-defined data type

package main

import "fmt"

type Score struct {
	name    string
	jersey_number int
	runs   int
}

func main() {
	s1 := Score{"Jos", 23, 101}
	fmt.Println(s1.name, "whose jersey_number is", s1.jersey_number, "scored", s1.runs)
	s2 := Score{jersey_number: 23, runs: 536}
	fmt.Println(s2)
	fmt.Println(s2.name)	
}
