package example

//并发聊天服务器

import (
	"fmt"
	"net"
	"strings"
	"time"
)

type Client struct {
	//用户有三个属性：用户名，用户地址，C
	C    chan string // 该通道是用来接收需要发送给客户端的数据
	Name string      // 用户名
	Addr string      // 客户端地址
}

//  存储在线用户的键值数据库，用map模拟
var onlin_client_Map = make(map[string]Client)

//  用于广播消息给用户的通道
var message = make(chan string)

//  广播消息给所有在线用户的方法
func Messager() {
	for {
		//  循环监听读取message中的数据
		msg := <-message

		//  遍历所有在线的用户，再将消息写入用户用于接收消息的通道中
		for _, cli := range onlin_client_Map {
			cli.C <- msg
		}
	}
}

//  用户将自己通道的消息写给用户客户端的方法
func WriteMsgToClient(cli Client, conn net.Conn) {
	for msg := range cli.C { // 遍历出通道中的信息
		conn.Write([]byte(msg + "\n")) //   利用通信socket，将信息传输给客户端
	}
}

//  消息生产方法
func MakeMessage(cli Client, msg string) string {
	str := "[" + cli.Addr + "] " + cli.Name + ": " + msg
	return str
}

//  这是程序的核心，上线提醒、发送消息、更换用户名、查看所有在线用户列表 的功能都在这个函数实现
func HandleConnect(conn net.Conn) {
	defer conn.Close() //不要忘记关闭

	//  获取 客户端 地址
	cli_addr := conn.RemoteAddr().String() //“.String()” 作用是：转成string类型
	//  初始化新用户
	cli := Client{make(chan string), cli_addr, cli_addr}

	//  添加新上线用户到Map中
	onlin_client_Map[cli_addr] = cli

	//  往全局通道中写入 登录信息
	message <- MakeMessage(cli, "login")

	//  用户将信息发送到客户端的go程
	go WriteMsgToClient(cli, conn)

	//  这两个通道是用来控制客户端的在线时间的
	isQuit := make(chan bool)
	hasData := make(chan bool)

	go func() {
		buf := make([]byte, 4096)

		for {

			n, err := conn.Read(buf) // 读取客户端发来的信息
			fmt.Println(n, err)
			if n == 0 {
				fmt.Printf("客户端%s断开\n", cli.Name)
				isQuit <- true
				return
			}
			if err != nil {
				fmt.Println("conn.Read err:", err)
				return
			}
			msg := string(buf[:n-1]) //去掉空格行
			msg = strings.TrimSpace(msg)
			if msg == "who" && len(msg) == 3 { //   查看在线用户列表
				//  不需要广播，所以不用message，直接往conn中写数据
				conn.Write([]byte("user list:\n"))
				for _, cli := range onlin_client_Map { //   遍历map中所有在线客户，将信息组织好后再发送
					msg = cli.Addr + ":" + cli.Name + "\n" //   这里没有使用MakeMessage,直接自己生成
					conn.Write([]byte(msg))
				}
			} else if len(msg) >= 8 && msg[:7] == "rename|" { //    更换用户名功能

				cli.Name = strings.Split(msg, "|")[1]
				onlin_client_Map[cli_addr] = cli
				conn.Write([]byte("rename success!\n"))
			} else {
				message <- MakeMessage(cli, msg)
			}
			hasData <- true
		}
	}()
	for {
		select {
		case <-isQuit:
			delete(onlin_client_Map, cli.Addr)
			message <- MakeMessage(cli, "log out")
		case <-hasData:

		case <-time.After(time.Second * 30): //  如果30s不发言，系统将强制将用户退出
			delete(onlin_client_Map, cli.Addr)
			message <- MakeMessage(cli, "time out leave")
			return
		}
	}
}
func StartChatServer() {
	//创建与客户端的连接地址
	listener, err := net.Listen("tcp", "127.0.0.1:10001") //利用的是tcp通信
	if err != nil {
		fmt.Println("net.Listen err:", err)
		return
	}
	defer listener.Close()

	//  启动go程 广播方法，等待接收数据，并广播   到每个客户
	go Messager()

	//  循环监听客户端连接
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Accept err:", err)
			return
		}
		go HandleConnect(conn) //收到一个客户连接，启动一个go程
	}
}
