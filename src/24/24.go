package main
import "fmt"
import "os"

var count int

func print(v []int) {
	if count == 999999 {
		for _,i := range v {
			fmt.Printf("%d", i);
		}
		fmt.Println()	
		os.Exit(0)
	}	
	count++
}

func swap(v []int, i int, j int) {	
	t := v[i];
	v[i] = v[j];
	v[j] = t;
}

func rotate(v []int, start int, n int) {
	tmp := v[start];
	for i := start; i < n-1; i++ {
		v[i] = v[i+1];
	}
	v[n-1] = tmp;
} 

func permute(v []int, start int, n int) {
	print(v);
	if start < n {
		for i := n-2; i >= start; i-- {
			for j := i + 1; j < n; j++ {
				swap(v, i, j);
				permute(v, i+1, n);
			}
			rotate(v, i, n);
		}
	}
}

func main() {

	// 
	// A permutation is an ordered arrangement of objects. For example, 3124 is one possible permutation of the digits 1, 2, 3 and 4. If all of the permutations are listed numerically or alphabetically, we call it lexicographic order. The lexicographic permutations of 0, 1 and 2 are:
	// 012   021   102   120   201   210
	// What is the millionth lexicographic permutation of the digits 0, 1, 2, 3, 4, 5, 6, 7, 8 and 9?
	// 
	//

	count = 0
	seq := []int{0,1,2,3,4,5,6,7,8,9}	
	permute(seq, 0, 10)

	// 0123456789 0
	// 0123456798 9
	// 0123456879 81
	// 0123456897 18
	// 0123456978 81
	// 0123456987 9

	// 0123457689 702
	// 0123457698 9
	// 0123457869 171
	// 0123457896 27
	// 0123457968 171
	// 0123457986 27

	// 0123458679 693
	// 0123458697 18
	// 0123458769 72
	// 0123458796 27
	// 0123458967 171
	// 0123458976 9
	
	// 0123459678 702

	// 012   
	// 021   
	// 102   
	// 120   
	// 201   
	// 210
 
}
 