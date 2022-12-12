package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(math.Pi)
	// Pi is exported from the math package. So, it should begin with a capital letter.
	// The same is applicable to Print which is exported from fmt package.
	// When importing a package, you can refer only to its exported names.
	// Any "unexported" names are not accessible from outside the package.
}

func Demo() {

}
