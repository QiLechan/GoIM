package main

import (
	"fmt"
)

func main() {
	var ip string
	var port int
	var input int
	fmt.Println("客户端模式输入1，服务端模式输入2")
	fmt.Scan(&input)
	if (input == 1){
		fmt.Println("请输入IP")
		fmt.Scan(&ip)
		fmt.Println("请输入端口")
		fmt.Scan(&port)
		client(ip, port)
	}else if (input == 2){
		fmt.Println("请输入端口")
		fmt.Scan(&port)
		server(port)
	}else {
		fmt.Println("输入错误")
	}
}