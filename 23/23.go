package main
import (
	"fmt"
	"math"
	"sort"
)

type Sequence []int

func (s Sequence) Len() int {
	return len(s)
}
func (s Sequence) Less(i, j int) bool {
	return s[i] < s[j]
}
func (s Sequence) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func disp(c int) int {
	
	sum := 1	
	for i := 2; i <= int(math.Sqrt(float64(c))); i++ {
		if c % i == 0 {
			if i == (c/i) {
				sum += i
			} else {
				sum += i + (c / i)
			}
		}
	}
	
	if sum < c {
		return -1
	} 
	if sum > c {
		return 1
	}
	return 0
}

func main() {

	// 
	// A perfect number is a number for which the sum of its proper divisors is exactly equal to the number. For example, the sum of the proper divisors of 28 would be 1 + 2 + 4 + 7 + 14 = 28, which means that 28 is a perfect number.
	// A number n is called deficient if the sum of its proper divisors is less than n and it is called abundant if this sum exceeds n.
	// 
	
	// As 12 is the smallest abundant number, 1 + 2 + 3 + 4 + 6 = 16, the smallest number that can be written as the sum of two abundant numbers is 24. By mathematical analysis, it can be shown that all integers greater than 28123 can be written as the sum of two abundant numbers. However, this upper limit cannot be reduced any further by analysis even though it is known that the greatest number that cannot be expressed as the sum of two abundant numbers is less than this limit.
	// Find the sum of all the positive integers which cannot be written as the sum of two abundant numbers.
	// 
	// 

	abunds := make(map[int]bool)

	for i := 5; i < 28123; i++ {
		if disp(i) > 0 {
			abunds[i] = true
		}
	}
	
	aseq := make(Sequence, len(abunds))

	i := 0
	for k, _ := range abunds {
		aseq[i] = k
		i++
	}
	
	sort.Sort(aseq)

	sum := 0
	for i := 1; i < 28123; i++ {
		found := false
		for _, j := range aseq {			
			if j >= i {
				break
			}
			if abunds[i - j] {
				found = true
				break
			}
		}

		if found == false {
			// fmt.Printf("%v couldn't be expressed as the sum of two abundants\n", i)
			sum += i
		}
	}

	fmt.Printf("Sum was: %v\n", sum)

}
 