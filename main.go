package main

import (
	"time"
	"fmt"
)

func main() {
	hello := make(chan string, 1)
	name := make(chan string, 0)

	go func() {
		//time.Sleep(time.Second * 1)
		name <- "qinming"
		name <- "wudi"
	}()

	go func() {
		hello <- "hello "
	}()
	//time.Sleep(time.Second * 2)
	for {
		select {
		case s := <-hello:
			fmt.Println(s)
		case s := <-name:
			fmt.Println(s)
		case <-time.After(time.Second * 3):
			fmt.Println("time out")
			return
		}
	}
}
