package db

import (
	"os"
)

type Database struct {
	Store map[string]uint32
}

func (d *Database) openDatabase(dbPath string) {

}
