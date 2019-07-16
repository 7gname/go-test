package study

import (
	"io/ioutil"
	"fmt"
)

var filepath = "./main.go"
var path = "."

func ReadFileTest()  {
	data, err := ioutil.ReadFile(filepath)
	if err != nil {
		fmt.Printf("err:[%v]\n", err)
	}
	fmt.Printf("file content:[\n%s\n]\n", data)
}

func ReadPathTest(arg... string)  {
	if len(arg) > 0 {
		path = arg[0]
	}
	files, err := ioutil.ReadDir(path)
	if err != nil {
		fmt.Printf("err:[%v]", err)
	}
	for _, file := range files {
		fn := file.Name()
		if fn == "." && fn == ".." {
			continue
		}
		if file.IsDir() {
			fmt.Printf("+[%s]\n", file.Name())
			//ReadPathTest(file.Name())
		}else {
			fmt.Printf("-[%s]\n", file.Name())
		}
	}
}

func ReadLenTest()  {

}
