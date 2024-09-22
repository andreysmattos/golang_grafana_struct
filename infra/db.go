package db

import "fmt"

type DB struct {
	driver string
}

func (d *DB) Query(q string) string {
	fmt.Println("[Database Mock] Quering... ", q)

	return q
}
