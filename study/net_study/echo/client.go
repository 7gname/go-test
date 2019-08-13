package echo

import (
	"bufio"
	"fmt"
	"net"
	"time"
)

func Client() {
	conn, err := net.Dial("tcp", "127.0.0.1:7001")
	if err != nil {
		fmt.Printf("err:%s\n", err.Error())
		return
	}
	defer conn.Close()

	for {
		str := fmt.Sprintf("hello, current timestamp:%d\n", time.Now().Unix())

		//n, err := conn.Write([]byte(str))
		//if err != nil {
		//	fmt.Printf("err:%s\n", err.Error())
		//	return
		//}
		//fmt.Printf("write text:%s success, length:%d\n", str, n)
		//
		//buff := make([]byte, 128)
		//len, err := conn.Read(buff)
		//if err != nil {
		//	fmt.Printf("err:%s\n", err.Error())
		//	return
		//}
		//
		//fmt.Println(string(buff[:len]))

		newWriter := bufio.NewWriter(conn)
		n, err := newWriter.WriteString(str)
		if err != nil {
			fmt.Printf("err:%s\n", err.Error())
			return
		}
		newWriter.Flush()
		fmt.Printf("write text:%s success, length:%d\n", str, n)

		rs, err := bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			fmt.Printf("err:%s\n", err.Error())
			return
		}
		fmt.Printf("return text:%s", rs)

		time.Sleep(time.Second * 3)
	}
}
