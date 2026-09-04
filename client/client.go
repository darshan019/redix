package main

import (
	"bufio"
	"fmt"
	"net"
)

func send(conn net.Conn, payload string, name string) {
	fmt.Printf("Sending %s\n", name)

	_, err := conn.Write([]byte(payload))
	if err != nil {
		fmt.Println("write error:", err)
		return
	}

	response, _ := bufio.NewReader(conn).ReadString('\n')
	fmt.Println(response)
}

func main() {
	conn, err := net.Dial("tcp", "localhost:6379")
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	send(conn, "*1\r\n+PING\r\n", "PING")
	send(conn, "*2\r\n+ECHO\r\n$5\r\nhello\r\n", "ECHO")

	send(
		conn,
		"*3\r\n+SET\r\n+foo\r\n+bar\r\n",
		"SET",
	)

	send(
		conn,
		"*3\r\n+SET\r\n+lob\r\n+bobo\r\n",
		"SET",
	)

	send(
		conn,
		"*2\r\n+GET\r\n+foo\r\n",
		"GET",
	)

	send(
		conn,
		"*2\r\n+DEL\r\n+foo\r\n",
		"DELETE",
	)
}

func test(conn net.Conn) {

	// + Simple String
	send(
		conn,
		"+HELLO\r\n",
		"Simple String",
	)

	// : Integer
	send(
		conn,
		":42\r\n",
		"Integer",
	)

	// $ Bulk String
	send(
		conn,
		"$5\r\nhello\r\n",
		"Bulk String",
	)

	// Empty Bulk String
	send(
		conn,
		"$0\r\n\r\n",
		"Empty Bulk String",
	)

	// * Array
	send(
		conn,
		"*3\r\n:1\r\n:2\r\n:3\r\n",
		"Array of Integers",
	)

	// Nested Array
	send(
		conn,
		"*3\r\n:1\r\n$5\r\nhello\r\n*2\r\n:2\r\n:3\r\n",
		"Nested Array",
	)

	fmt.Println("Done")
}
