package study

import (
	"fmt"
	"strings"
	"io"
)

func toLowerOrUpper () {
	var s = "Abcd"
	fmt.Println(strings.ToLower(s))
	fmt.Println(strings.ToUpper(s))
}

func mMap()  {
	var s = "abc"
	fmt.Println(strings.Map(func(r rune) rune {
		fmt.Println(r)
		return r
	},s))
}

func reader()  {
	var s = "hello world!"
	r := strings.NewReader(s)
	fmt.Println(r.Size())
	fmt.Println(r.Len())
	b := make([]byte, 5)
	for {
		n,err := r.Read(b)
		if err != nil && err != io.EOF {
			panic("err")
		}
		if n < 1 {
			break
		}
		fmt.Println(n)
		fmt.Println(string(b))
	}
	fmt.Println(s)
}

func StudyString () {
	//toLowerOrUpper()
	//mMap()
	reader()
}
