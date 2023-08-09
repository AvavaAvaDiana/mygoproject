package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	service := "127.0.0.1:8086"
	tcpAddr, err := net.ResolveTCPAddr("tcp", service)
	checkError1(err)

	listener, err := net.ListenTCP("tcp", tcpAddr)
	checkError1(err)
	defer listener.Close()

	fmt.Println("Server is listening on", service)

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting connection:", err)
			continue
		}
		go handleClient(conn)
	}
}

func handleClient(conn net.Conn) {
	defer conn.Close()

	fmt.Println("Accepted connection from", conn.RemoteAddr())

	// 在这里可以处理连接并回复数据
}

func checkError1(err error) {
	if err != nil {
		fmt.Println("Error:", err.Error())
		os.Exit(1)
	}
}
