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
	sqrtn := int64(math.Ceil(math.Sqrt(float64(N))))

	primes := sieve()

	var factors [2]int64

	for i := 0; i < 5000; i++ {
		next := <-primes
		fmt.Printf("checking %v\n", next)
		if int64(next) > int64(sqrtn) {
			fmt.Println("giving up, last prime was greater than the sqrt of the target")
			break
		}
		if N % int64(next) == 0 {			
			fmt.Printf("%v is a prime factor\n", next)
			factors[0] = factors[1]
			factors[1] = int64(next)
		}
	}	
	
}