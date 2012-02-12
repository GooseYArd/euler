package main
import (
	"fmt"
)

func gen_primes (c []bool) {
	
	limit := len(c) 
	c[1] = true              // 1 not considered prime
	p := 2
	for {
		p2 := p * p
		if p2 >= limit {
			break
		}
		for i := p2; i < limit; i += p {
			c[i] = true // it's a composite
			
		}
		for {
			p++
			if !c[p] {
				break
			}
		}
	}
}

func main() {

	// 
	// Euler published the remarkable quadratic formula:
	// n² + n + 41
	// It turns out that the formula will produce 40 primes for the consecutive values n = 0 to 39. 
        // However, when n = 40, 40^2 + 40 + 41 = 40(40 + 1) + 41 is divisible by 41, and 
        // certainly when n = 41, 41² + 41 + 41 is clearly divisible by 41.
	// Using computers, the incredible formula n²  79n + 1601 was discovered, 
        // which produces 80 primes for the consecutive values n = 0 to 79. 
        // The product of the coefficients, 79 and 1601, is 126479.
	// Considering quadratics of the form:
	// 
	// n² + an + b, where |a|  1000 and |b|  1000
	// where |n| is the modulus/absolute value of ne.g. |11| = 11 and |4| = 4
	//
	// Find the product of the coefficients, a and b, for the quadratic expression 
        // that produces the maximum number of primes for consecutive values of n, starting with n = 0.
	//
	// 
	
	iscomp := make([]bool, 2000001)	
	gen_primes(iscomp)
	longest := 0
	longa := int64(0)
	longb := int64(0)
	
	for a := int64(-1000); a <= 1000; a++ {
		for b := int64(-1000); b <= 1000; b++ {
			nprime := 0
			for n := int64(0); n < 80; n++ {
				idx := n*n + a*n + b
				if idx < 0 {
					idx = -idx
				}
				if ! iscomp[idx] {
					nprime++					
				} else {
					break
				}				
			}
			if nprime > longest {
				longest = nprime
				longa = a
				longb = b				
				fmt.Printf("a: %v b: %v yielded %v primes\n", longa, longb, nprime) 
			}
		}
	}

}
 