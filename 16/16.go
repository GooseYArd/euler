package main
import "fmt"
import "math/big"

func main() {

	// 
	// 
	// 2^15 = 32768 and the sum of its digits is 3 + 2 + 7 + 6 + 8 = 26.
	// What is the sum of the digits of the number 2^1000?
	// 
	// 
	

	zero := big.NewInt(0)
	p := big.NewInt(2)
	exp := big.NewInt(1000)
	ten := big.NewInt(10)
	m := big.NewInt(0)
	sum := int64(0)
	p = p.Exp(p, exp, nil)

	fmt.Printf("p: %v m: %v\n", p.String(), m.String())
	for {
		fmt.Printf("p: %v m: %v\n", p.String(), m.String())
		ten.SetInt64(int64(10))
		if p.Cmp(zero) < 1 {
			break
		}
		p, m = p.DivMod(p, ten, ten)
		sum += m.Int64()
	}
	fmt.Printf("%v\n", sum)
}
 