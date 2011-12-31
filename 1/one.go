package main

import fmt "fmt"

func main() {

	limit := 1000
	sum := 0	
	j := 0

	for i := 3; i < limit; i += 3 {
		if j < 4 { 
			sum += i
			j++
		} else {
			j = 0
		}
	}	

	for i := 5; i < limit; i += 5 {
		sum += i
	}	
	
	fmt.Printf("the sum of all natural numbers between 1 and 1000 which are multiples of 3 or 5 is %d\n", sum)
	

}