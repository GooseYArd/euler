package main
import "fmt"

func main() {

	// 
	// The fraction 49/98 is a curious fraction, as an inexperienced mathematician in attempting to simplify it may incorrectly believe that 49/98 = 4/8, which is correct, is obtained by cancelling the 9s.
	// We shall consider fractions like, 30/50 = 3/5, to be trivial examples.
	// There are exactly four non-trivial examples of this type of fraction, less than one in value, and containing two digits in the numerator and denominator.
	// If the product of these four fractions is given in its lowest common terms, find the value of the denominator.
	// 
	// 

	dp := 1
	np := 1
	
	for n := 1.0; n < 10; n++ {
		for d := n+1; d < 10; d++ {
			for m:=1.0; m < 10; m++ {
				a:=n/d
				b:=(n*10 + m) / (m*10+d)
				if a - b == 0 {
					fmt.Printf("%v/%v (%v) == %v/%v (%v), a - b = %v\n", n, d, n/d, n*10+m, m*10+d, (n*10+d) / (m*10+d), a - b)
					np *= int(n)
					dp *= int(d)
				}
			}
			
		}
	}

	fmt.Printf("dn/dp = %v/%v\n", np, dp)

}
 