package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
)

func main() {
	db := NewDb()
	listener, err := net.Listen("tcp", ":6379")
	if err != nil {
		panic(err)
	}
	defer listener.Close()

	fmt.Println("Listening in port 6379")

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("protocol error", err)
			continue
		}

		go handleConn(conn, db)

	}
}

func handleConn(conn net.Conn, db *DB) {
	defer conn.Close()
	reader := bufio.NewReader(conn)

	for {
		prefix, err := reader.ReadByte()
		if err != nil {

			if err == io.EOF {
				fmt.Println("EOF, client connection closed")
			} else {
				fmt.Println("error reading byte", err)
			}
			return
		}

		data, err := handleParse(reader, prefix)
		if err != nil {
			fmt.Println(err)
			return
		}

		dataArray, ok := data.([]any)
		if !ok {
			fmt.Println("expected RESP array")
			_, _ = conn.Write([]byte("ERR expected array\r\n"))
			continue
		}

		value, err := handleCommands(dataArray, db)
		if err != nil {
			fmt.Println(err)
			return
		}
		// fmt.Println(db.db, value)
		response := fmt.Appendf(nil, "+%v\r\n", value)
		_, err = conn.Write(response)

		if err != nil {
			fmt.Println("write error:", err)
			return
		}
		fmt.Println(string(response))
	}
}
