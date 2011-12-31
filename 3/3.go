package main

import (
	"fmt"
	"math"
)

func is_square(n int64) bool {
	sq := math.Sqrt(float64(n))
	return sq - math.Floor(sq) > 0		
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

func main() {

	// The prime factors of 13195 are 5, 7, 13 and 29.
	// What is the largest prime factor of the number 600851475143 ?
	
	N := int64(600851475143)
// 	a := int64(math.Ceil(math.Sqrt(float64(N))))
// 	b2 := int64(math.Pow(float64(a), 2)) - N
	
// 	for {
// 		if ! is_square(b2) {
// 			a += 1
// 			b2 = int64(math.Pow(float64(a), 2)) - N
// 		} else {
// 			break
// 		}		
// 	}
	
	primes := sieve()

	var factors [2]int64

	for i := 0; i < 5000; i++ {
		next := <-primes
		if N % int64(next) == 0 {			
			fmt.Printf("%v is a prime factor\n", next)
			factors[0] = factors[1]
			factors[1] = int64(next)
// 			fmt.Printf("testing other factor: %v\n", N / int64(next))
// 			guess := N / int64(next)
// 			terms := sieve()
// 			prime := true
// 			for j := 0; j < 10000; j++ { 
// 				t := <- terms				
// 				if guess % int64(t) == 0 {
// 					fmt.Printf("whoops, %v is not prime\n", guess)
// 					prime = false
// 					break
// 				}
// 			}			
// 			if prime {
// 				fmt.Printf("I think %v is prime\n", guess)
// 			}
				
		}
	}
	
	newbig := factors[0] * factors[1]

	if N % newbig == 0 {
		fmt.Printf("also try %v\n", N / newbig)
	}

	
}