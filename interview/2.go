package interview

import "fmt"

func defer_call() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("one=", err)
		}
	}()
	defer func() { fmt.Println("打印前") }()
	defer func() { fmt.Println("打印中") }()
	defer func() { fmt.Println("打印后") }()

	panic("触发异常")

}