package main

import (
	"fmt"
	"flag"
	"os"
	"strconv"
	"path/filepath"
	"net/http"
	"html"
	"bytes"
	"strings"
)

func fetch_description (num int) (n string, err error) {
	
	n = "ok"
	
	url := fmt.Sprintf("http://projecteuler.net/problem=5")
	resp,err := http.Get(url)
	if err != nil {
		fmt.Printf("Error fetching description: %v",err.Error())
		return n, err
	}
	
	var desc bytes.Buffer
	//buf.ReadFrom(resp.Body)
	z := html.NewTokenizer(resp.Body)
	
	in_desc := false
	depth := 0
	for {
		tt := z.Next()
		switch tt {
		case html.ErrorToken:
			fmt.Printf("returning ErrorToken, captured %v", string(desc.Bytes()))
			return string(desc.Bytes()), err
			//return z.Err()
		case html.TextToken:
			if in_desc {
				desc.Write(z.Text())
			}
		case html.StartTagToken, html.EndTagToken:
			tn, _ := z.TagName()
			stn := string(tn)
			if stn == "div" {
				if tt == html.StartTagToken {
					key, val, _ := z.TagAttr()
					if string(key) == "class" && string(val) == "problem_content" {
						in_desc = true
					}
				} else {
					if in_desc {
						return string(desc.Bytes()), err
					}
				}
			}
			
			if len(tn) == 1 && tn[0] == 'a' {
				if tt == html.StartTagToken {
					depth++
				} else {
					depth--
				}
			}
		}
	}

	n = string(desc.Bytes())
	resp.Body.Close()	
	return n, err
}

func main() {
	
	var eroot string
	flag.StringVar(&eroot, "root", ".", "euler root")
	flag.Parse()
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
		os.RemoveAll(nextdir)
		panic("Unable to open a new makefile")
	}

	f.WriteString("include $(GOROOT)/src/Make.inc")
	f.WriteString("\n")
	f.WriteString(fmt.Sprintf("TARG=%d\n", next))
	f.WriteString("GOFILES=\\\n")
	f.WriteString(fmt.Sprintf("\t%d.go\n", next))
	f.WriteString("\n")
	f.WriteString("include $(GOROOT)/src/Make.cmd\n")
	f.Close()

	srcfile := filepath.Join(nextdir, fmt.Sprintf("%d.go", next))
	f, err = os.Create(srcfile)
	if err != nil {
		os.RemoveAll(nextdir)
		panic("Unable to open a new source file")
	}

	desc, err := fetch_description(next)
	if err != nil {
		os.RemoveAll(nextdir)		
		panic("Unable to fetch project description")
	}
	
	desc = strings.Replace(desc, "\n", "\n\t// ", -1)

	f.WriteString("package main")
	f.WriteString("\n")
	f.WriteString("import \"fmt\"\n")
	f.WriteString("\n")
	f.WriteString("func main() {\n")
	f.WriteString("\n")
	f.WriteString("\t// ")
	f.WriteString(desc)
	f.WriteString("\n")
	f.WriteString("}\n ")

	f.Close()

}