package main
import (
	"fmt"
	"bufio"
	"os"
	"strings"
	"strconv"
)

type grid [20][20]int

func load_grid(fn string, r *grid) (lines []string, err error) {

	var (
		file *os.File
		line string
	)
	
	if file, err = os.Open(fn); err != nil {
		return
	}

	reader := bufio.NewReader(file)
	// buffer := bytes.NewBuffer(make([]byte, 1024))

	for i := 0; i <20 ; i++  {
		if line, err = reader.ReadString('\n'); err != nil {
			break
		}
		line = strings.TrimRight(line, "\n")
		for j, e := range strings.SplitN(line, " ", 20) {
			r[i][j], err = strconv.Atoi(e)
		}
	}

	//if err == os.EOF {
	//	err = nil
	//}
	return lines, err
}

func main() {

	//
	// In the 2020 grid below, four numbers along a diagonal line have been marked in red.
	//
	//
	// The product of these numbers is 26  63  78  14 = 1788696.
	// What is the greatest product of four adjacent numbers in any direction (up, down, left, right, or diagonally) in the 2020 grid?
	// 
	// 

	var g grid
	max := 0
	load_grid("grid.txt", &g)

	// horizontal
	for i := 0; i < 20; i++ {
		for j := 0; j < 16 ; j++ {
			prod := g[i][j]
			for m := 1; m < 4; m++ {
				prod *= g[i][j+m]
				if prod > max {
					max = prod
				}
			}
		}
	}

	// vertical
	for i := 0; i < 20; i++ {
		for j := 0; j < 16 ; j++ {
			prod := g[j][i]
			for m := 1; m < 4; m++ {
				prod *= g[j+m][i]
				if prod > max {
					max = prod
				}
			}
		}
	}

	// descending
	for i := 0; i < 17; i++ {
		for j := 0; j < 17 ; j++ {
			prod := g[i][j]
			for m := 1; m < 4; m++ {
				prod *= g[i+m][j+m]
				if prod > max {
					max = prod
				}
			}
		}
	}

	// ascending
	for i := 3; i < 20; i++ {
		for j := 0; j < 17 ; j++ {
			prod := g[i][j]
			for m := 1; m < 4; m++ {
				prod *= g[i-m][j+m]
				if prod > max {
					max = prod
				}
			}
		}
	}
	fmt.Printf("ascending scan yielded: %v\n", max)



}



 