package main

import (
	"fmt"
	"net"
	"strings"
	"strconv"
	"bufio"
)

func client(ip string, port int) {
	var addr string = ip + ":" + strconv.Itoa(port)
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		fmt.Println("Error connecting:", err.Error())
		return
	}
	defer conn.Close()

	// 发送一条消息给服务器
	messageToSend := "Hello from client!\n"
	_, err = conn.Write([]byte(messageToSend))
	if err != nil {
		fmt.Println("Error sending message:", err.Error())
		return
	}

	// 如果需要接收服务器响应（例如Echo服务器会返回相同的消息），可以使用bufio读取：
	for{
		reader := bufio.NewReader(conn)
		response, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading response:", err.Error())
		} else {
			response = strings.TrimSpace(response)
			fmt.Printf("Received response from server: %s\n", response)
		}
		fmt.Scan(&messageToSend)
		_, err = conn.Write([]byte(messageToSend))
		if err != nil {
			fmt.Println("Error sending message:", err.Error())
			return
		}
	}
}