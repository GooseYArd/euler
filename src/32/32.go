package main
import "fmt"

type pdContext [10]int

func (pdctx *pdContext) reset() {
	for i := range(pdctx) {
		pdctx[i] = 0
	}
}

func (pdctx *pdContext) evaluate(values []int) (bool) {
	
	for _,v := range values {
		for v > 0 {
			tens := v % 10
			if tens == 0 {
				return false
			}
			if pdctx[tens] > 0 {
				return false
			}
			pdctx[tens] = 1
			v /= 10
		}
	}

	for _,v := range pdctx[1:] {
		if v == 0 {
			return false
		}
	}
	return true
}

func main() {

	// 
	// We shall say that an n-digit number is pandigital if it makes use of all the digits 1 to n exactly once; for example, the 5-digit number, 15234, is 1 through 5 pandigital.
	// 
	// The product 7254 is unusual, as the identity, 39  186 = 7254, containing multiplicand, multiplier, and product is 1 through 9 pandigital.
	// 
	// Find the sum of all products whose multiplicand/multiplier/product identity can be written as a 1 through 9 pandigital.
	// 
	// HINT: Some products can be obtained in more than one way so be sure to only include it once in your sum.
	// 
	// 

	seen := make(map[int]bool)
	pdctx := new(pdContext)
	sum := 0
	
	for prod := 1; prod < 10000; prod++ {
		for i := 1; (i * i) <= prod; i++ {
			if prod % i == 0 {
				vals := []int{prod, i, prod / i}
				if pdctx.evaluate(vals) {
					if seen[prod] == false {
						fmt.Printf("Found one: %v * %v = %v\n", i, prod/i, prod)
						sum += prod
						seen[prod] = true
					}

				}
				pdctx.reset()				
			}
		}

	}
	fmt.Printf("Sum is %v\n", sum)
	
}
 