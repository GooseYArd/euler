package main
import "fmt"

type lexicon [10][]string

func main() {
	
	// 
	// If the numbers 1 to 5 are written out in words: one, two, three, four, five, then there are 3 + 3 + 5 + 4 + 4 = 19 letters used in total.
	// If all the numbers from 1 to 1000 (one thousand) inclusive were written out in words, how many letters would be used? 
	// 
	// NOTE: Do not count spaces or hyphens. For example, 342 (three hundred and forty-two) contains 23 letters and 115 (one hundred and fifteen) contains 20 letters. The use of "and" when writing out numbers is in compliance with British usage.
	// 
		
	//hundred := "hundred"

	// replace these with lengths once i know its working
	// also, we only need to count 0-10 and 20-1000 once

	ones := []string{"","one","two","three","four","five","six","seven","eight","nine",}
	special := []string{"ten", "eleven", "twelve", "thirteen", "fourteen", "fifteen", "sixteen", "seventeen", "eighteen", "nineteen"}
	tens := []string{"","","twenty","thirty","fourty","fifty","sixty","seventy","eighty","ninety",}	
		
	total := 0
	for i:= 1; i <= 1000; i++ {		
		thousands := i / 1000
		if thousands > 0 {
			total += len(fmt.Sprintf("%vthousand", ones[thousands]))
			fmt.Printf("%vthousand", ones[thousands])
		}
		hundreds := i / 100 % 10
		if hundreds > 0 {
			fmt.Printf("%vhundred", ones[hundreds])
			total += len(fmt.Sprintf("%vhundred", ones[hundreds]))
		}
 		t := i / 10 % 10
 		if t > 0 {
 			fmt.Printf("%v", tens[t])
 			total += len(fmt.Sprintf("%v", tens[t]))
 		}
 		o := i % 10
 		if t == 1 {			
 			fmt.Printf("%v", special[o])
			total += len(fmt.Sprintf("%v", special[o]))
 		} else {
 			fmt.Printf("%v", ones[o])
 			total += len(fmt.Sprintf("%v", ones[o]))
 		}
		fmt.Printf(" %v\n", total)
	}
	
 	fmt.Printf("I think its %v\n", total)

  	oneslen := 0
  	for _, v := range ones {
  		oneslen += len(v)
  	}
 	fmt.Printf("oneslen: %v\n", oneslen)
	
  	tenslen := 0
  	for _, v := range tens {
  		tenslen += len(v)
  	}
 	fmt.Printf("tenslen: %v\n", tenslen)

  	speciallen := 0
  	for _, v := range special {
  		speciallen += len(v)
	}

	hundredlen := (9 * oneslen) + (tenslen * 10) + speciallen
	thousandlen := (10 * hundredlen) + (9 * oneslen)
 	fmt.Printf("thousandlen: %v\n", thousandlen)

	
}
