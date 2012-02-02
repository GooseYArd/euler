package main
import (
	"fmt"
	"math/big"
)
	
func main() {

	// 
	// The Fibonacci sequence is defined by the recurrence relation:
	// Fn = Fn1 + Fn2, where F1 = 1 and F2 = 1.
	// Hence the first 12 terms will be:
	// F1 = 1
	// F2 = 1
	// F3 = 2
	// F4 = 3
	// F5 = 5
	// F6 = 8
	// F7 = 13
	// F8 = 21
	// F9 = 34
	// F10 = 55
	// F11 = 89
	// F12 = 144
	// The 12th term, F12, is the first term to contain three digits.
	// What is the first term in the Fibonacci sequence to contain 1000 digits?
	// 
	// 

	fn1 := big.NewInt(1)
	fn2 := big.NewInt(1)

	ten := big.NewInt(10)	
	limit := ten.Exp(ten, big.NewInt(999), nil)

	term := 0
	f := big.NewInt(0)
	for ; f.Cmp(limit) < 0; f = f.Add(fn1, fn2) {		
		term++
		fn1.Set(fn2)
		fn2.Set(f)		
	}
	
	fmt.Printf("Term: %v, %v", term, len(f.String()))
}
 