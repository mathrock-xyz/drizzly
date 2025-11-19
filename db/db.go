package db

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/adrg/xdg"
	bolt "go.etcd.io/bbolt"
)

func init() {
	if err := setup(); err != nil {
		_, _ = fmt.Println(err.Error())
		os.Exit(1)
		return
	}
}

func setup() (err error) {
	path := filepath.Join(xdg.Home, ".config", "drizzly")

	err = os.MkdirAll(path, 0777)
	if err != nil {
		return
	}

	DB, err = bolt.Open(filepath.Join(path, "dz.db"), 0600, bolt.DefaultOptions)
	return
}

var DB *bolt.DB
