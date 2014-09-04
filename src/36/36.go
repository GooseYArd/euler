package main

import (
	"fmt"
	"math"
)

func rotator(start int) chan int {

	// for a number with 6 digits there can be at most 10 ^ (log10(n) / 2) palindromes	
	maxpal := int(math.Pow10(int(math.Log10(1000000) / 2)))

	// 001100
	// 010010
	// 100001
	ch := make(chan int)

	// pow10(4) + pow10(3)
	// pow10(5) + pow10(2)
	// pow10(6) + pow10(1)
	// etc
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

func reverse_bits(input uint32) uint32 {  
	input = (input & 0x55555555) <<  1 | (input & 0xaaaaaaaa) >>  1
	input = (input & 0x33333333) <<  2 | (input & 0xcccccccc) >>  2
	input = (input & 0x0f0f0f0f) <<  4 | (input & 0xf0f0f0f0) >>  4
	input = (input & 0x00ff00ff) <<  8 | (input & 0xff00ff00) >>  8
	input = (input & 0x0000ffff) << 16 | (input & 0xffff0000) >> 16
	
	for (input % 2) == 0 {
		input = input >> 1;
	}
	return input;
}

func main() {

	// 
	// 
	// The decimal number, 585 = 10010010012 (binary), is palindromic in both bases.
	// Find the sum of all numbers, less than one million, which are palindromic in base 10 and base 2.
	// (Please note that the palindromic number, in either base, may not include leading zeros.)
	// 
	// 
	
	N := 999999
	fmt.Printf("Working backwards from %v\n", N)	
	rot := rotator(N)	
	for {
		next := <-rot
		if next < 1 {
			break
		}
		//fmt.Printf("next %v/%v\n", next, reverse_bits(uint32(next)));
		if uint32(next) & reverse_bits(uint32(next)) == uint32(next) {
			fmt.Printf("next %v\n", next);
		}		
	}
}
 
