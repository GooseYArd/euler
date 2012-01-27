package main
import (
	"fmt"
)

func main() {

	// 
	// 
	// A unit fraction contains 1 in the numerator. The decimal representation of the unit fractions with denominators 2 to 10 are given:
	// 
	// 1/2= 0.5
	// 1/3= 0.(3)
	// 1/4= 0.25
	// 1/5= 0.2
	// 1/6= 0.1(6)
	// 1/7= 0.(142857)
	// 1/8= 0.125
	// 1/9= 0.(1)
	// 1/10= 0.1
	// 
	// Where 0.1(6) means 0.166666..., and has a 1-digit recurring cycle. It can be seen that 1/7 has a 6-digit recurring cycle.
	// Find the value of d  1000 for which 1/d contains the longest recurring cycle in its decimal fraction part.
	// 
	// 
		
	// The period of 1/k for integer k is always ≤ k − 1.

	max := 0
	maxd := 0
	var bigseq []int

	for i := 1; i < 1000; i++ {
		divisor := i
		dividend := 1	

		leading_zeros := 0
		in_zeros := true

		l := make([]int, i*2)
		for j:=0; j<(i*2); j++ {
			quotient := dividend / divisor
			dividend -= (quotient * divisor)
			dividend *= 10		
			if quotient == 0 {
				if in_zeros {
					leading_zeros++
				}
			} else {
				in_zeros = false
			}
			
			if in_zeros == false {
				l[j-leading_zeros] = quotient
			}			
		}		
		
		longest := 0
		for b := i; b > 1; b-- {
			match := true
			length := 0
			for ; length < b ; length++ {
				if l[1+length] != l[b+length] {
					match = false
					break
				}
			}

			if match {
				//fmt.Printf("%v: found match of length %v\n", i, b)
				longest = b-1
				break
			}
		}
		
		for j := 1; j < longest; j++ {
			ok := true
			for k := 1 ;k < longest; k++ {
				if l[k] != l[j+k] {
					ok = false
					break;
				}
			}
			if ok {
				//fmt.Printf("%v: new length is %v\n", i, j)
				longest = j
				break
			}
		}
		
		if longest > max {				
			// fmt.Printf("The new longest is d=%v (len %v)\n", i, longest)
			max = longest
			maxd = i
			bigseq = l[0:longest]
		}
	}

	fmt.Printf("longest cycle d: %v, length %v\n", maxd, max)

	for _,i := range bigseq {
		fmt.Printf("%d", i)
	}

}
 