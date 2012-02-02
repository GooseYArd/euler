package main
import (
	"fmt"
	"math"
)

func main() {

	//
	// 2520 is the smallest number that can be divided by each of the numbers from 1 to 10 without any remainder.
	// What is the smallest positive number that is evenly divisible by all of the numbers from 1 to 20?
	// 
	// 
	
	var res int64 = 20

	// save some division later by skipping composite values or squares
	for i := 19; i > 2; i-- {
		prime := true
		for j := i-1; j > 1; j-- {
			if i % j == 0 {
				prime = false
			}			
		}		
		if prime || math.Remainder(math.Sqrt(float64(i)), 1.0) == 0 { 
			res *= int64(i)			
		}
	}

	fmt.Println()	
	fmt.Println("Testing solution...")
	if ! valid(res) {
		panic("The algorithm is incorrect")		
	}

	fmt.Println("value is valid, continuing")
	
	res = reduce(res)	
	fmt.Printf("Final answer: %v\n", res)
	if (valid(res)) {
		fmt.Printf("and its ok\n")
	} else {
		fmt.Printf("NOT OK\n")
	}
	
}
 
func valid(n int64) bool {
	var i int64
	for i = 20; i > 1; i-- {
		if n % int64(i) > 0{
			return false
		}
	}
	return true
}

func reduce(n int64) (res int64) {
	res = n	
	for i := int64(20); i > 1; i-- {
		//fmt.Printf("attempting with %v\n", i)
		for {
			testval := res / i
			if valid(testval) {
				//fmt.Printf("reducing by %v\n", i)
				res = testval
			} else {
				//fmt.Printf("Can't divide by %v anymore\n", i)
				break
			}
		}
	}
	return res
}
