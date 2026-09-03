package main

import (
	"net"
)

func main() {
	conn, _ := net.Dial("tcp", "localhost:6379")
	defer conn.Close()

	conn.Write([]byte(
        "*2\r\n" +
        ":1\r\n" +
        "*2\r\n" +
        "+OK\r\n" +
        "$5\r\nhello\r\n",
    ))
}
