package main

import (
	"fmt"
	"flag"
	"os"
	"strconv"
	"path/filepath"
)


func main() {
	
	var eroot string
	flag.StringVar(&eroot, "root", ".", "euler root")
	flag.Parse()
	fmt.Printf("euler root %s", eroot)

	d, err := os.Open(eroot)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fi, err := d.Readdir(-1)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	next := 0
	for _, fi := range fi {
		if fi.IsDir() {
			s, err := strconv.Atoi(fi.Name())
			if err == nil {
				if s > next {
					next = s
				}
			}			
		}
	}
	
	next += 1
	fmt.Printf("next project is %d\n", next)
	nextdir := filepath.Join(eroot, strconv.Itoa(next))
	err = os.Mkdir(nextdir, 0755)
	if err != nil {
		panic("Unable to make nextdir!")
	}

	mkfile := filepath.Join(nextdir, "Makefile")
	f, err := os.Create(mkfile)
	if err != nil {
		panic("Unable to open a new makefile")
		os.RemoveAll(nextdir)
	}

	f.WriteString("include $(GOROOT)/src/Make.inc")
	f.WriteString("\n")
	f.WriteString(fmt.Sprintf("TARG=%d\n", next))
	f.WriteString("GOFILES=\\\n")
	f.WriteString(fmt.Sprintf("\t%d.go\n", next))
	f.WriteString("\n")
	f.WriteString("include $(GOROOT)/src/Make.cmd\n")
	f.Close()
}