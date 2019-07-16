package study

import (
	"os"
	"fmt"
	"bufio"
	"strings"
	"testing"
)

func TestBaseOpt(t *testing.T)  {
	fmt.Println(os.Getenv("NODE_ENV"))

	fileInfo, err := os.Stat("./main.go")
	fmt.Printf("file info[%#v], err[%#v]\n", fileInfo, err)

	fp, err := os.Open("./main.go")
	fmt.Printf("file[%#v], err[%#v]\n", fp, err)

	if err := os.Chdir("log"); err != nil {
		f, err := os.Open(".")
		if err != nil {
			fmt.Printf("err[%#v]\n", err)
		}
		fileInfo, err := f.Readdir(-1)
		if err != nil {
			fmt.Printf("err[%#v]\n", err)
		}
		for _, file := range fileInfo {
			fmt.Printf("file[%#v]\n", file)
		}
	}
}

func TestStdin(t *testing.T)  {
	stop := make(chan string)
	fmt.Printf("input \"stop\" to stop server\n")
	newReader := bufio.NewReader(os.Stdin)
	for {
		cmd, _ := newReader.ReadString('\n')
		//fmt.Printf("%+v\n", strings.TrimSpace(cmd) == "stop")
		if strings.TrimSpace(cmd) == "stop" {
			println(111)
			go func() {
				stop<- "stop"
			}()
			println(222)
			break
		}else{
			fmt.Println("cmd not defined")
			continue
		}
	}
	fmt.Println("cmd right")
	c := <-stop
	fmt.Println(c)
	fmt.Println("over")
}
