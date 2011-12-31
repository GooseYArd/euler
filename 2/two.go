package main

import "fmt"

func main() {

	// 1, 2, 3, 5, 8, 13, 21, 34, 55, 89, ...

	sum_of_evens := 0
	last := 1
	for i := 1; i <= 4000000; {
		tmp := i
		i += last
		last = tmp
		if (i % 2) == 0 {
			sum_of_evens += i
		}
	}

	fmt.Printf("sum %d\n", sum_of_evens)
}