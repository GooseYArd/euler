package main
import "fmt"
import "math"

func main() {

	// 
	// Let d(n) be defined as the sum of proper divisors of n (numbers less than n which divide evenly into n).
	// If d(a) = b and d(b) = a, where a != b, then a and b are an amicable pair and each of a and b are called amicable numbers.
	// For example, the proper divisors of 
	// 220 are 1, 2, 4, 5, 10, 11, 20, 22, 44, 55 and 110; 
	// therefore d(220) = 284. The proper divisors of 284 are 1, 2, 4, 71 and 142; so d(284) = 220.
	// Evaluate the sum of all the amicable numbers under 10000.
	// 

	sum := 0
	for i := 1; i < 10000; i++ {
		cand := 1
		for j := 2; j < int(math.Sqrt(float64(i))); j++ {
			if i % j == 0 {
				cand += j + (i / j)
			}
		}

		p := 1
		for k := 2; k < int(math.Sqrt(float64(cand))); k++ {
			if cand % k == 0 {
				p += k + (cand / k)
			}
		}

		if (i == p && i != cand) {
			fmt.Printf("i: %v Cand %v Partner %v\n", i, cand, p)
			// we'll see p later
			sum += i
		}
	}

	fmt.Printf("Sum: %v\n", sum)

}
 