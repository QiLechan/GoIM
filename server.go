package main

import (
	"bufio"
	"fmt"
	"net"
	"strconv"
	"strings"
)

func server(port int){
	var addr string = ":" + strconv.Itoa(port)
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		fmt.Println("Error listening:", err.Error())
		return
	}
	defer listener.Close()
	fmt.Println("Server is listening on port:" + strconv.Itoa(port))
	for {
		// 接受一个来自客户端的连接请求
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting: ", err.Error())
			continue
		}
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	// defer conn.Close()

	reader := bufio.NewReader(conn)
	writer := bufio.NewWriter(conn)

	for {
		// 从客户端读取一行数据
		message, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading:", err.Error())
			break
		}

		// 去除末尾换行符并打印接收到的消息
		message = strings.TrimSpace(message)
		fmt.Printf("Received message from client: %s\n", message)

		// 将消息写回客户端
		writer.WriteString(message + "\n")
		writer.Flush()
	}
}