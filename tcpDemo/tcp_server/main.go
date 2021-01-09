package main

import (
	"net"

	myutils "github.com/timyuheng/MyTest/myutils"
)

func main() {

	myutils.Encode()
}

func process(conn net.Conn) {
	defer conn.Close()
	myutils.Decode(conn)
}
