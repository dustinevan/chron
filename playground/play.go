package main

import (
	"fmt"
	"math"
)

func main() {
	sum1 := 0.0
	sum2 := 0.0
	for n := 0; n < 337; n++ {
		sum1 += 2032 / math.Pow(1.0+.035/12, float64(n))
		sum2 += 2032 / math.Pow(1.0+.055/12, float64(n))
	}
	fmt.Printf("%.2f \n%.2f \n%.2f", sum1, sum2, sum1-sum2)
}
