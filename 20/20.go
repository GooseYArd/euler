package main
import "fmt"
import "math/big"

func factorial(n *big.Int) *big.Int {
	one := big.NewInt(1)
	if n.Cmp(one) == 0 {	
		return one
	}

	nn := big.NewInt(0)
	nn = nn.Set(n)
	nn = nn.Sub(nn, one)

	r := big.NewInt(0)
	f := factorial(nn)
	r = r.Mul(n, f)

	return r
}

func main() {

	// 
	// n! means n  (n  1)  ...  3  2  1
	// For example, 10! = 10  9  ...  3  2  1 = 3628800,and the sum of the digits in the number 10! is 3 + 6 + 2 + 8 + 8 + 0 + 0 = 27.
	// Find the sum of the digits in the number 100!
	// 
	// 

	zero := big.NewInt(0)
	onehundred := big.NewInt(100)
	r := factorial(onehundred)
	fmt.Printf("100!: %v\n", r.String())
	m := big.NewInt(0)
	sum := big.NewInt(0)
	
	for ; r.Cmp(zero) > 0 ; {		
		ten := big.NewInt(10)
		r,m = r.DivMod(r, ten, ten)
		sum.Add(sum, m)
	}

	fmt.Printf("Sum of digits of 100!: %v\n", sum.String())
	
}
 