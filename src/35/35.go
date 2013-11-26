package main
import "fmt"
import "math"

func permute(v []int) (r []int) {
	if len(v) == 1 {
		r = append(r,v[0])
		return r
	}

	for i := 0; i < len(v); i++ {
		tmp := v[0]		
		for j := 0; j < (len(v) - 1); j++ {			
			v[j] = v[j+1]
		}
		v[len(v)-1] = tmp
		r = append(r, arr_to_int(v))
	}	
	return r
}

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

func generate() chan int {
	ch := make(chan int)
	go func() {
		for i := 2; ; i++ {
			ch <- i
		}
	}()
	return ch
}

func filter(in chan int, prime int) chan int {
	out := make(chan int)
	go func() {
		for {
			if i := <-in; i%prime != 0 {
				out <- i
			}
		}
	}()
	return out
}

func sieve() chan int {
	out := make(chan int)
	go func() {
		ch := generate()
		for {
			prime := <-ch
			out <- prime
			ch = filter(ch, prime)
		}
	}()
	return out
}

func gen_primes (c []bool) {	
	limit := len(c) 
	c[1] = true              // 1 not considered prime
	p := 2
	for {
		p2 := p * p
		if p2 >= limit {
			break
		}
		for i := p2; i < limit; i += p {
			c[i] = true // it's a composite
			
		}
		for {
			p++
			if !c[p] {
				break
			}
		}
	}
}

func test_rotation(v []int, iscomp []bool) bool {
	for _, j := range(permute(v)) {
		if iscomp[j] == true {
			return false
		}		
	}
	return true
}

func main() {

	// 
	// 
	// The number, 197, is called a circular prime because all rotations of the digits: 197, 971, and 719, are themselves prime.
	// There are thirteen such primes below 100: 2, 3, 5, 7, 11, 13, 17, 31, 37, 71, 73, 79, and 97.
	// How many circular primes are there below one million?
	// 
	// 

	iscomp := make([]bool, 1000000)	
	gen_primes(iscomp)

	for i, comp := range iscomp {
		if !comp {
			v := int_to_arr(i)
			if test_rotation(v, iscomp) {
				fmt.Printf("All rotations of %d are prime\n", i)				
			}			
		}
	}


}
 
