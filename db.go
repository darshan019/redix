package main

import (
	"fmt"
	"sync"
)

type DB struct {
	db map[string]any
	mu *sync.Mutex
}

func (db *DB) Set(key string, value any) {
	db.mu.Lock()
	db.db[key] = value
	db.mu.Unlock()
}

func (db *DB) Get(key string) any {
	db.mu.Lock()
	defer db.mu.Unlock()
	return db.db[key]
}

func (db *DB) Delete(key string) error {
	value := db.Get(key)
	if value == nil {
		return fmt.Errorf("key not found")
	}
	db.mu.Lock()
	defer db.mu.Unlock()
	delete(db.db, key)	
	return nil
}

func NewDb() *DB {
	return &DB{
		db: make(map[string]any),
		mu: &sync.Mutex{},
	}
}

func handleDBOps(data []any, db *DB) (any, error) {
	cmd, ok := data[0].(string)
	if !ok {
		return nil, fmt.Errorf("invalid command")
	}
	switch cmd {
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
		err := db.Delete(key)
		if err != nil {
			return nil, err
		}
		return "OK", nil
	default:
		return nil, fmt.Errorf("Invalid command")
	}
}
