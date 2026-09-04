package main

import "fmt"

func handleCommands(data []any, db *DB) (any, error) {
	if len(data) == 0 {
		return nil, fmt.Errorf("empty command")
	}

	switch data[0] {
	case "ECHO":
		if len(data) < 2 {
			return nil, fmt.Errorf("missing argument")
		}
		return data[1], nil

	case "PING":
		return "PONG", nil

	default:
		return handleDBOps(data, db)
	}
}