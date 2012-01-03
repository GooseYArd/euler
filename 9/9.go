package main
import (
	"fmt"
	"os"
)

func main() {

	// 
	// A Pythagorean triplet is a set of three natural numbers, a  b  c, for which,
	//  a2 + b2 = c2
	// For example, 32 + 42 = 9 + 16 = 25 = 52.
	// There exists exactly one Pythagorean triplet for which a + b + c = 1000.Find the product abc.
	// 
	// 
	
	for a := 1; a <= 998; a++ {
		for b := 1; b <= 998; b++ {
			c := 1000 - (a + b)
			if ((a*a) + (b*b) == (c*c)) {
				fmt.Printf("Found it: a=%v b=%v c=%v\n", a, b, c)
				os.Exit(0)
			}
		}
	}
	

}
 