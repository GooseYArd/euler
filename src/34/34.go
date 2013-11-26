package main
import "fmt"
import "math"

func int_to_arr(n int) (x []int) {
	for n > 0 {
		x = append(x, n % 10)
		n = n / 10
	}
	return x
}

func arr_to_int(x []int) (n int) {
	for i,v := range x {
		n += int(math.Pow10(i)) * v
	}
	return n
}

func fact(n int) int {
	facts := [...]int{1, 1, 2, 6, 24, 120, 720, 5040, 40320, 362880}
	return facts[n]
}

func is_factorion(n int) bool {
	t := 0
	for _, v := range int_to_arr(n) {
		t += fact(v)
	}
	return t == n
}

func main() {
	// 
	// 
	// 145 is a curious number, as 1! + 4! + 5! = 1 + 24 + 120 = 145.
	// Find the sum of all numbers which are equal to the sum of the factorial of their digits.
	// Note: as 1! = 1 and 2! = 2 are not sums they are not included.
	// 
	// 
	//                0  1  2  3   4    5    6     7      8       9

	// things we know:
	// any number with an n in it cannot be in a factorion < n!

	t := 0
	for v := 3; v < 1000000; v++ {
		if is_factorion(v) {
			t += v
			fmt.Printf("%d\n", v)
		}
	}
	
	fmt.Printf("Total: %d\n", t)

}
 
