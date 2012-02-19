package main
import "fmt"

func combinations(target int, coins []int) (n int) {		

	if target == 0 {
		return 0
	}
	
	if len(coins) == 0 {
		return 0
	}
	
	// if only one value of coin is left, and that value divides evenly
	// into the remaining value, then exactly one combination remains
	if len(coins) == 1 {
		if target % coins[0] > 0 {
			return 0
		}
		return 1
	}
	
	for i := 0; i * coins[0] < target; i++ {
		v := combinations(target - (i * coins[0]), coins[1:])				
		n+=v
	}	

	// in case i * coins[0] == target
	if target % coins[0] == 0 {
		n++
	}

	return n
}


func main() {

	// 
	// 
	// In England the currency is made up of pound, £, and pence, p, and there are eight coins in general circulation:
	// 1p, 2p, 5p, 10p, 20p, 50p, £1 (100p) and £2 (200p).
	// It is possible to make £2 in the following way:
	// 1£1 + 150p + 220p + 15p + 12p + 31p
	// How many different ways can £2 be made using any number of coins?
	// 
	// 
	
	coins := []int{200, 100, 50, 20, 10, 5, 2, 1}
	//coins := []int{200, 100, 50, 20}

	// 200
	// 100, 100
	// 100, 50, 50
	// 50, 50, 50, 50

	target := 200
	
	total := combinations(target, coins)

	fmt.Printf("Total combinations is: %v\n", total)
}
 