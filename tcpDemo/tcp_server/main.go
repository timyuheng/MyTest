package main

import (
	"fmt"
	"net"

	"github.com/timyuheng/MyTest/myutils"
)

func main() {

	listener, err := net.Listen("tcp", "0.0.0.0:9090")
	if err != nil {
		fmt.Println("Listen err:", err)
		return
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Accept err:", err)
			continue
		}

		go process(conn)

	}

	fmt.Println("tcp server end...")
}

func process(conn net.Conn) {
	defer conn.Close()
	fmt.Println("get conn")
	for {

		bt, err := myutils.Decode2(conn)
		if err != nil {
			fmt.Println("myutils.Decode err:", err)
			break
		}

		fmt.Println("receice data:", string(bt))
	}

}
