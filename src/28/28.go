package main
import "fmt"

// func printGrid(grid [][]int) {
// 	for y,row := range grid {
// 		for x := range row {
// 			fmt.Printf("%2d ", grid[x][y])
// 		}
// 		fmt.Println()
// 	}
// 	fmt.Println()
// }

func main() {

	// 
	// Starting with the number 1 and moving to the right in a clockwise direction a 5 by 5 spiral is formed as follows:
	// 21 22 23 24 25
	// 20  7  8  9 10
	// 19  6  1  2 11
	// 18  5  4  3 12
	// 17 16 15 14 13
	// It can be verified that the sum of the numbers on the diagonals is 101.
	// What is the sum of the numbers on the diagonals in a 1001 by 1001 spiral formed in the same way?
	// 
	// 
	
	sum := int64(1)
	next := int64(3)
	k := int64(2)
	for k <= 1001 {
		for i := 0; i < 4; i++ {
			fmt.Printf("adding next of %v\n", next)
			sum += next
			next = next + k			
		}		
		fmt.Printf("length %v sum %v\n", k+1, sum)
		next += 2
		k+=2
	}	
	fmt.Printf("Sum of diagonals is %v\n", sum)
}


	// 21 22 23 24 25
	// 20  7  8  9 10
	// 19  6  1  2 11
	// 18  5  4  3 12
	// 17 16 15 14 13

 