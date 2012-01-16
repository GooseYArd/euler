package main
import "fmt"

type Node struct {
	val int
}

func NewNode(leftpar *Node, rightpar *Node, n int) *Node {
	leftval := 0
	rightval := 0

	if leftpar != nil {
		leftval = leftpar.val
	}

        if rightpar != nil {
		rightval = rightpar.val
	}
	r := new(Node)
	r.val = n + Max(leftval, rightval)
	return r
}

func Max(v1, v2 int) int { 
        if v1 >= v2 { 
                return v1 
        } 
        return v2 
} 

func main() {

	// 
	// By starting at the top of the triangle below and moving to adjacent numbers on the row below, the maximum total from top to bottom is 23.
	// 3
	// 7 4
	// 2 4 6
	// 8 5 9 3
	// That is, 3 + 7 + 4 + 9 = 23.
	// Find the maximum total from top to bottom of the triangle below:

	v := []int{
		75,
		95,64, 
		17,47,82,
		18,35,87,10,
		20, 4,82,47,65,
		19, 1,23,75, 3,34,
		88, 2,77,73, 7,63,67,
		99,65, 4,28, 6,16,70,92,
		41,41,26,56,83,40,80,70,33,
		41,48,72,33,47,32,37,16,94,29,
		53,71,44,65,25,43,91,52,97,51,14,
		70,11,33,28,77,73,17,78,39,68,17,57,
		91,71,52,38,17,14,91,43,58,50,27,29,48,
		63,66, 4,68,89,53,67,30,73,16,69,87,40,31,
		 4,62,98,27,23, 9,70,98,73,93,38,53,60, 4,23,
	}
	
	p := make([][]int, 15) 
	
	c := 0	
	for i := 0; i < 15; i++ {
		p[i] = make([]int, i+1)		
		for j := range p[i] {
			p[i][j] = v[c]
			c++
		}
	}

	branches := [][]*Node{}
	for ridx, row := range p {
		branch := []*Node{}
		for cidx := range row {
			var leftpar *Node
			var rightpar *Node
			if ridx > 0 {
				parentbr := branches[ridx-1]
				if cidx < len(row) - 1 {
					rightpar = parentbr[cidx]
				}				
				if cidx > 0 {
					leftpar = parentbr[cidx - 1]
				}
			}			
			node := NewNode(leftpar, rightpar, row[cidx])
			branch = append(branch, node)
		}
		branches = append(branches, branch)
	}
	max := 0
	for _, n := range branches[len(branches) - 1] {
		if n.val > max {
			max = n.val
		}
	}
	
	fmt.Printf("Max: %v\n", max)
}