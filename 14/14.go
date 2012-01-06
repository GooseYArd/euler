package main
import (
	"fmt"
)

type Cache map[int]int
var cache Cache

func next_k(k int64) int64 {
	if k % int64(2) == 0 {
		return k/int64(2)
	}
	return k * int64(3) + int64(1)
}

func make_list(start int64) (a int64) {
	// fmt.Printf("called with %v\n", start)
	//if cache[start] > 0 {
	//	return cache[start]
	//}
	
	a = 1
	if start > 1 {
		if start % 2 == 0 {
			//fmt.Println("found an even, recursing")			
			a += make_list(start / 2)
		} else {
			//fmt.Println("found an odd, recursing")
			a += make_list(start * 3 + 1)
		}
	}
	//cache[start] = a
	return a
}

func main() {
	
	// 
	// The following iterative sequence is defined for the set of positive integers:
	// n  n/2 (n is even)
	// n  3n + 1 (n is odd)
	// Using the rule above and starting with 13, we generate the following sequence:
	// 13  40  20  10  5  16  8  4  2  1
	// It can be seen that this sequence (starting at 13 and finishing at 1) contains 10 terms. Although it has not been proved yet (Collatz Problem), it is thought that all starting numbers finish at 1.
	// Which starting number, under one million, produces the longest chain?
	// NOTE: Once the chain starts the terms are allowed to go above one million.
	// 
	// 

	// RAB NOTES:
	// the recursive solution was about 50% slower
	// with recursion plus a history, it was about 75% slower- the array indexing is slower than the integer math
	// also, at least one of the members of the longest set wraps int, I kept getting the answer wrong until
	// I switched to int64

	cache = make(Cache)	
	
	max := int64(0)
	max_start := int64(0)

	if false {
		for i := int64(1); i < 1000000; i++ {
			c:= int64(make_list(i))
			if (c > max) {
				max = c
				max_start = i
			}
		}
	} else {
		
		for i := int64(1); i < 1000000; i++ {
			c := int64(0)
			for k := i; k > 1; k = next_k(k) {
				c++
			}
			if (c > max) {
				max = c
				max_start = i
			}
		}
	}
	fmt.Printf("longest chain started with %v, length %v\n", max_start, max)

}
 