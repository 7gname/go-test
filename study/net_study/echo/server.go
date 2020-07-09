package echo

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"strings"
)

func processTelnetCommand(str string) bool {
	switch {
	case strings.HasPrefix(str, "@close"):
		fmt.Printf("session closed\n")
		return false
	default:
		fmt.Println(str)
		return true
	}
}

func handler(conn net.Conn) {
	defer conn.Close()
	fmt.Printf("session started\n")
	reader := bufio.NewReader(conn)
	for {
		str, err := reader.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				break
			}
			fmt.Printf("read err:%s, session closed\n", err.Error())
			continue
		}

		str = strings.TrimSpace(str)

		if !processTelnetCommand(str) {
			break
		}

		conn.Write([]byte(str + "\r\n"))
	}
}

func Server() {
	stop := make(chan string)
	address := "127.0.0.1:7001"

	listen, err := net.Listen("tcp", address)
	if err != nil {
		fmt.Printf("server start err:%s\n", err.Error())
		stop <- "stop"
	}

	fmt.Printf("server start success, address:%s\n", address)

	defer listen.Close()
	go func() {
		for {
			conn, err := listen.Accept()
			if err != nil {
				fmt.Printf("accept err:%s\n", err.Error())
				continue
			}

			go handler(conn)
		}
	}()

	<-stop
}
