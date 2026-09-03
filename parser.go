package main

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
	"strings"
)


func handleParse(reader *bufio.Reader, prefix byte) (any, error) {
	switch prefix {
	case '+':
		data, err := parseString(reader)
		if err != nil {
			return nil, err
		}
		return data, nil
	case ':':
		data, err := parseInteger(reader)
		if err != nil {
			return nil, err
		}
		return data, nil
	case '$':
		data, err := parseBulkString(reader)
		if err != nil {
			return nil, err
		}
		return data, nil
	case '*':
		data, err := parseArray(reader)
		if err != nil {
			return nil, err
		}
		return data, nil
	default:
		err := fmt.Errorf("Incorrect message format")
		return nil, err
	}
}

func parseString(reader *bufio.Reader) (any, error) {
	line, err := reader.ReadString('\n')
	if err != nil {
		return nil, err
	}
	line = strings.TrimSuffix(line, "\r\n")
	return line, nil
}

func parseInteger(reader *bufio.Reader) (int, error) {
	line, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("error reading input", err)
		return 0, err
	}
	line = strings.TrimSuffix(line, "\r\n")

	data, err := strconv.ParseInt(line, 10, 64)
	if err != nil {
		fmt.Println("Incorrect integer value")
		return 0, err
	}
	return int(data), nil
}

func parseBulkString(reader *bufio.Reader) (any, error) {
	line, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("error reading input", err)
		return "", err
	}
	line = strings.TrimSuffix(line, "\r\n")
	length, err := strconv.ParseInt(line, 10, 64)
	if err != nil {
		fmt.Println("Incorrect length value")
		return "", err
	}

	// string of len -1
	if length == -1 { 
		return nil, nil
	}

	buf := make([]byte, length)
	_, err = io.ReadFull(reader, buf)
	if err != nil {
		return "", err
	}

	cr, err := reader.ReadByte()
	if err != nil {
		return "", err
	}

	lf, err := reader.ReadByte()
	if err != nil {
		return "", err
	}

	if cr != '\r' || lf != '\n' {
		return "", fmt.Errorf("expected CRLF")
	}


	return string(buf), nil
}

func parseArray(reader *bufio.Reader) ([]any, error) {
	line, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("error reading input", err)
		return nil, err
	}
	line = strings.TrimSuffix(line, "\r\n")
	length, err := strconv.ParseInt(line, 10, 64)
	if err != nil {
		fmt.Println("Incorrect length value")
		return nil, err
	}

	if length == -1 {
		return nil, nil
	}

	data := make([]any, int(length))

	for i := 0; i < int(length); i++ {
		prefix, err := reader.ReadByte()
		if err != nil {
			return nil, err
		}
		subData, err := handleParse(reader, prefix)
		if err != nil {
			return nil, err
		}
		data[i] = subData
	}

	return data, nil
}