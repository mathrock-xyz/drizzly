package info

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"

	"github.com/mathrock-xyz/drizzly/db"
	"github.com/spf13/cobra"
	bolt "go.etcd.io/bbolt"
)

var cat = &cobra.Command{
	Use:   "cat [name]",
	Short: "Display the description of a file or directory",
	RunE: func(cmd *cobra.Command, args []string) (err error) {
		if len(args) < 1 {
			return errors.New("noo")
		}

		name := args[0]
		if name == "" {
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

		var desc string

		_ = db.DB.View(func(tx *bolt.Tx) (err error) {
			val := tx.Bucket(bucket).Get([]byte(path))

			if val == nil {
				desc = "Nooo"
			}

			desc = string(val)
			return
		})

		_, err = fmt.Println(desc)
		return
	},
}
