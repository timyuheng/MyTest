package main

import (
	"fmt"
	"net"

	"github.com/timyuheng/MyTest/myutils"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:9090")
	defer conn.Close()
	if err != nil {
		fmt.Printf("connect failed, err : %v\n", err.Error())
		return
	}
	myutils.Encode2(conn, "hello world 0!!!")
}
