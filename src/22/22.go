package main
import (
	"fmt"
	"os"
	"bufio"
	"sort"
	"strings"
)

type names_t map[string] int

func load_names(fn string) (names names_t, err error) {
	var (
		file *os.File
		name []byte
	)

	names = make(map[string] int)
	if file, err = os.Open(fn); err != nil {
		return names, err
	}
	
	reader := bufio.NewReader(file)

	for {
		name, err = reader.ReadBytes(',')
		s := strings.Trim(string(name), "\",")		
		names[s] = 0
		for _,c := range(s) {
			names[s] += int(c) - 64
		}
		if err != nil {
			break
		}
	}	
	return names, err
}


func main() {

	// 
	// 
	// Using names.txt (right click and 'Save Link/Target As...'), a 46K text file containing over five-thousand first names, begin by sorting it into alphabetical order. Then working out the alphabetical value for each name, multiply this value by its alphabetical position in the list to obtain a name score.
	// For example, when the list is sorted into alphabetical order, COLIN, which is worth 3 + 15 + 12 + 9 + 14 = 53, is the 938th name in the list. So, COLIN would obtain a score of 938  53 = 49714.
	// What is the total of all the name scores in the file?
	// 
	// 
	
	names, err := load_names("names.txt")	
	if err != nil {
	//	panic("couldn't load name list\n")
	}

	mk := make([]string, len(names))
	i := 0
	for k, _ := range names {
		mk[i] = k
		i++
	}

	sort.Strings(mk)	
	sum := int64(0)

	for i,v := range mk {
		fmt.Printf("Pos: %v Name %v Val %v\n", i+1, v, names[v])
		sum += int64((i+1) * names[v])
	}
	
	fmt.Printf("Total: %v\n", sum)

}
 