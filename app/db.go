package main

import (
	"fmt"
	"lib"
)

// SQLiteDB - highlevel DB interface
type SQLiteDB interface {
	SQLInit() error
	Close()
	CheckDBversion() (int, error)
	GetLongLink(shortLink string, domain string, longLink *lib.LongLink) error
}

// NewDB - Create DB instance
func NewDB(dbPath string) SQLiteDB {
	db := new(lib.SQLiteDBImplementation)
	db.ConnString = fmt.Sprintf("%s?mode=ro", dbPath)
	return db
}

// OpenDB - Open new SQLite3 Database
func OpenDB(dbPath string) SQLiteDB {
	db := NewDB(dbPath)
	db.SQLInit()
	return db
}
