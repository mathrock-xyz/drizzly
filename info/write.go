package info

import (
	"errors"
	"os"
	"path/filepath"

	"github.com/mathrock-xyz/drizzly/db"
	"github.com/spf13/cobra"
	bolt "go.etcd.io/bbolt"
)

var write = &cobra.Command{
	Use:   "write [name] [desc]",
	Short: "Add or update a description for a file or directory",
	RunE: func(cmd *cobra.Command, args []string) (err error) {
		if len(args) < 1 {
			return errors.New("noo")
		}

		name, desc := args[0], args[1]
		if name == "" && desc == "" {
			return errors.New("")
		}

		obj, err := os.Stat(name)
		if err != nil {
			return
		}

		abs, err := filepath.Abs(obj.Name())
		if err != nil {
			return
		}

		path := filepath.Join(abs, obj.Name())
		return db.DB.Update(func(tx *bolt.Tx) error {
			return tx.Bucket(bucket).Put([]byte(path), []byte(desc))
		})
	},
}
