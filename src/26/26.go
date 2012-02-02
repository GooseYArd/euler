package main
import (
	"stree"
	"strconv"
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

	for i := 987; i < 988; i++ {
		divisor := i
		dividend := 1

		// l := make([]int, i*2)
		var l string

		for j:=0; j<(i*2); j++ {
			quotient := dividend / divisor
			dividend -= (quotient * divisor)
			dividend *= 10
			// l[j] = quotient
			l += strconv.Itoa(quotient)
		}

		fmt.Printf("1/%v: %v\n", i, l)
		t, _ := stree.CreateTree(l)
		t.PrintTree()
	}
}
