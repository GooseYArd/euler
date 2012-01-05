package main
import "fmt"
import "math"

func main() {

	// 
	// The sequence of triangle numbers is generated by adding the natural numbers. So the 7th triangle number would be 1 + 2 + 3 + 4 + 5 + 6 + 7 = 28. The first ten terms would be:
	// 1, 3, 6, 10, 15, 21, 28, 36, 45, 55, ...
	// Let us list the factors of the first seven triangle numbers:
	//  1: 1
	//  3: 1,3
	//  6: 1,2,3,6
	// 10: 1,2,5,10
	// 15: 1,3,5,15
	// 21: 1,3,7,21
	// 28: 1,2,4,7,14,28
	// We can see that 28 is the first triangle number to have over five divisors.
	// What is the value of the first triangle number to have over five hundred divisors?
	// 
	// 

	t := int64(1)	
	max := 0
	for i := int64(2); ; i++ {
		t += i		
		c := 0
		for k := int64(1); k < int64(math.Sqrt(float64(t)) +1); k++ {
	 		if t % k == 0 {
				c+=2
			}
		}
		
		if c > max  {
			max = c
			if c >  500 {
				break
			}
		}		
	}

	fmt.Printf("%v has %v factors\n", t, max)	

}
 