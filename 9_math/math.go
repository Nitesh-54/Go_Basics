package main

import (
	"fmt"
	"math"
)

func main() {
	var num float64 = 7
	result := math.Sqrt(num)
	roundResult := math.Round(result)
	ceilResult := math.Ceil(result)
	floorResult := math.Floor(result)
	fmt.Printf("Basic result is %.3f\n", result)
	fmt.Printf("Basic result is %.3g\n", result)
	fmt.Printf("Round result is %.3f\n", roundResult)
	fmt.Printf("Ceil result is %.3f\n", ceilResult)
	fmt.Printf("Floor result is %.3f", floorResult)
}
