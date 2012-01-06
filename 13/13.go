package main
import (
	"fmt"
	"math/big"
	"os"
	"bufio"
	"strings"
)
type hundred_bigs [100]big.Int

func load_bigs(fn string, r []big.Int) (err error) {
	
	var (
		file *os.File
		line string
	)
	
	if file, err = os.Open(fn); err != nil {
		return err
	}

	reader := bufio.NewReader(file)

	for i := 0; i <100 ; i++  {
		if line, err = reader.ReadString('\n'); err != nil {
			break
		}
		line = strings.TrimRight(line, "\n")
		r[i].SetString(line, 10)
		if err != nil {
			return err
			// panic("couldn't set an element!")
		}
	}
	return err
}


func main() {

	// 
	// 
	// Work out the first ten digits of the sum of the following one-hundred 50-digit numbers.
	// 
	// 

	var biglist hundred_bigs	
	load_bigs("numbers.txt", biglist[:])

	sum := new(big.Int)

	for _, val := range biglist {
		sum.Add(sum, &val)		
	}

	fmt.Printf("%v\n", sum.String()[0:10]) 

}
 