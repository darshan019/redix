package main

import "fmt"

type DB struct {
	db map[string]any
}

func (db *DB) Set(key string, value any) {
	db.db[key] = value
}

func (db *DB) Get(key string) any {
	return db.db[key]
}

func NewDb() *DB {
	return &DB{
		db: make(map[string]any),
	}
}

func handleDBOps(data []any, db *DB) (any, error) {
	switch data[0] {
	case "SET":
		db.Set(data[1].(string), data[2])
		return "OK", nil
	case "GET":
		value := db.Get(data[1].(string))
		return value, nil
	case "DEL":
		key, ok := data[1].(string)
		if !ok {
			return nil, fmt.Errorf("invalid argument")
		}

		value := db.Get(key)
		if value == nil {
			return nil, fmt.Errorf("key not found")
		}

		delete(db.db, key)
		return "OK", nil
	default:
		return nil, fmt.Errorf("Invalid command")
	}
}
