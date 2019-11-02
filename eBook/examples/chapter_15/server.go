package main

import (
	"fmt"
	"net"
)

func main() {
	fmt.Println("Starting the server ...")
	// 建立 listener
	listener, err := net.Listen("tcp", "localhost:50000")
	if err != nil {
		fmt.Println("Error listening", err.Error())
		return //終止程式
	}
	// 監聽並接受來自客户端的連線
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting", err.Error())
			return // 終止程式
		}
		go doServerStuff(conn)
	}
}

func doServerStuff(conn net.Conn) {
	for {
		buf := make([]byte, 512)
		len, err := conn.Read(buf)
		if err != nil {
			fmt.Println("Error reading", err.Error())
			return //終止程式
		}
		fmt.Printf("Received data: %v", string(buf[:len]))
	}
}