package db

import "fmt"

var (
	_ IDB = (*DB)(nil)
)

type IDB interface {
	Query(string) string
}

type DB struct {
	driver string
}

func ProvideDb() IDB {
	return &DB{}
}

func (d *DB) Query(q string) string {
	fmt.Println("[Database Mock] Quering... ", q)

	return q
}
