package main

import (
	"fmt"
	"math"
)

func is_square(n int64) bool {
	sq := math.Sqrt(float64(n))
	return sq - math.Floor(sq) > 0		
}

func largest_prime_factor (comp int64) (lpf int64) {
	
	maxsqrt := int64(math.Sqrt(float64(comp)) + 1)
	iscomp := make([]bool, maxsqrt, maxsqrt)

	for m := int64(2); m < maxsqrt; m++ {
		if ! iscomp[m] {			
			for k := m * m; k <= maxsqrt; k+= m {
				iscomp[k] = true
			}
		}
	}
	
	for m := maxsqrt - 1; m >= 2; m-- {
		if ! iscomp[m] {
			if comp % int64(m) == 0 {
				return int64(m)
			}
		}
	}	
	return 1
}

func main() {

	// The prime factors of 13195 are 5, 7, 13 and 29.
	// What is the largest prime factor of the number 600851475143 ?
	
	N := int64(600851475143)
	fmt.Printf("largest prime factor is %v\n", largest_prime_factor(N))

	
}