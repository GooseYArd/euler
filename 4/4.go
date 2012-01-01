package main

import (
	"fmt"
	"math"
	"os"
)

func rotator(start int) chan int {

	// for a number with 6 digits there can be at most 10 ^ (log10(n) / 2) palindromes	
	maxpal := int(math.Pow10(int(math.Log10(1000000) / 2)))

	// 001100
	// 010010
	// 100001
	ch := make(chan int)
	base := []int{1100, 10010, 100001}	
	go func() {		
		for i := 0; i < maxpal; i++ {
			v := (i % 10) * base[0]
			v += ((i % 100) / 10) * base[1]
			v += ((i % 1000) / 100) * base[2]
			ch <- start - v
		}
	}()
	return ch
}

func main() {
	
	// largest possible 3 digit numbers
	// N := 999 * 999 == 998001
	N := 999999	
	fmt.Printf("Working backwards from %v\n", N)	
	rot := rotator(N)	
	for {
		next := <-rot
		if next < 1 {
			break
		}
		for i := 999; i >= 100; i-- {
			d := next / i
			r := math.Remainder(float64(next), float64(i))			
			if d > 99 && d < 1000 && r == 0 {				
				fmt.Printf("N %v A %v B %v\n", next, d,i )
				os.Exit(0)
			}
		}

	}
	
	
}