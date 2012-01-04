package main
import (
	"fmt"
	"math"
)

func sum_primes (max int) (sum int64) {
	sum = 0
	var iscomp [2000001]bool
	maxsqrt := int(math.Sqrt(float64(max + 1)))
	for m := 2; m <= maxsqrt; m++ {
		if ! iscomp[m] {
			sum += int64(m)
			for k := m * m; k <= max; k+= m {
				iscomp[k] = true
			}
		}
	}
	
	for m := maxsqrt; m <= max; m++ {
		if ! iscomp[m] {
			sum +=int64(m)
		}
	}	
	return sum
}

func main() {

	//
	// The sum of the primes below 10 is 2 + 3 + 5 + 7 = 17.
	// Find the sum of all the primes below two million.
	// 
	
	fmt.Printf("Sum: %v", sum_primes(2000000))

}

