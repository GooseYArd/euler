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
	// Starting in the top left corner of a 22 grid, there are 6 routes (without backtracking) to the bottom right corner.
	// 
	// 
	// 
	// How many routes are there through a 2020 grid?
	// 
	// 
	// 1x2 = 4 edges, 2 paths
	// 2x2 = 12 edges, 6 paths
	// 3x3 = 24 edges, 
	// 4x4 = 40 edges
	// 5x5 = 60 edges
	// 6x6 = 
	
	// (4 * 1) + (0 * 1) + (0 * 1) 4   2 paths
	// (4 * 2) + (1 * 2) + (1 * 2) 12  6 paths
	// (4 * 3) + (2 * 3) + (2 * 3) 24 
	// (4 * 4) + (3 * 4) + (3 * 4) 40
	// (4 * 5) + (4 * 5) + (4 * 5) 60
	// (4 * 6) + (5 * 6) + (5 * 6) 84
	
	// ( m + n ) ! / ( m! * n! )

	for i := int64(1); i < 21; i++ {
		fii := factorial(big.NewInt(i+i))
		fi := factorial(big.NewInt(i))

		fi = fi.Mul(fi,fi)
		
		r := big.NewInt(1)

		r = r.Div(fii, fi)

		fmt.Printf("%v: %v / %v = %v\n", i, fii.String(), fi.String(), r.String())
	}
			
}
 