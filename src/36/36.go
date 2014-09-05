package main

import (
	"fmt"
)

func reverse_bits(input uint32) uint32 {  
	if input == 0 {
		return 0
	}
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

func check(x uint32) bool{		
	return (x & reverse_bits(x) == x)
}

func showifpal(x uint32) {
	if check(x) {
		fmt.Printf("%v\n", x)
	}
	return
}

func main() {

	// 
	// 
	// The decimal number, 585 = 10010010012 (binary), is palindromic in both bases.
	// Find the sum of all numbers, less than one million, which are palindromic in base 10 and base 2.
	// (Please note that the palindromic number, in either base, may not include leading zeros.)
	// 
	// 
	
	for c := uint32(0); c < 1000; c++ {
		number := c
		upper := c
		lower := uint32(0)
		digit := uint8(0)
		base := uint32(10)
		power := uint32(1)
		
		for number > 0 {
			upper = upper * base
			digit = uint8(number % base)
			lower = lower * base + uint32(digit)
			number = number / base
			power *= base
		}
		
		showifpal(upper + lower)
		last := upper * base + lower			
		showifpal(last)
		
		for i := uint32(1); i < base; i++ {			
			showifpal(last + i * power)
		}
	}
}
 
