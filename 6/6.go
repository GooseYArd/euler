package main
import (
	"fmt"
)

func main() {

	//
	// The sum of the squares of the first ten natural numbers is,
	// 1^2 + 2^2 + ... + 102 = 385
	// The square of the sum of the first ten natural numbers is,
	// (1 + 2 + ... + 10)2 = 552 = 3025
	// Hence the difference between the sum of the squares of the first ten natural numbers and the square of the sum is 3025  385 = 2640.
	// Find the difference between the sum of the squares of the first one hundred natural numbers and the square of the sum.
	// 
	// 
		
	for k := int64(100); k <= 100; k++ {
		var sum_of_squares int64 = 0
		var square_of_sums int64 = 0

		for i := int64(0); i <= k; i++ {
			sum_of_squares += i * i
			square_of_sums += i
		}		
		square_of_sums *= square_of_sums		
		fmt.Printf("sum_of_squares: %v\n", sum_of_squares)
		fmt.Printf("square_of_sums: %v\n", square_of_sums)
		fmt.Printf("difference: %v\n", square_of_sums - sum_of_squares)		

		// these are from the #6 thread, hah!
		fmt.Printf("square_of_sums by alternate method: %v\n", ((k+1) * (k/2)) * ((k+1) * (k/2)))
		fmt.Printf("sum_of_squares by alternate method: %v\n", k * (k+1) * (2*k+1) * 1/6)

	}
}
 