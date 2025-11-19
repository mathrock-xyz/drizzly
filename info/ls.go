package info

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/mathrock-xyz/drizzly/db"
	"github.com/spf13/cobra"
	bolt "go.etcd.io/bbolt"
)

var ls = &cobra.Command{
	Use:   "ls",
	Short: "",
	RunE: func(cmd *cobra.Command, args []string) (err error) {
		entries, err := os.ReadDir(".")
		if err != nil {
			return
		}

		for _, entry := range entries {
			abs, err := filepath.Abs(entry.Name())
			if err != nil {
				return err
			}

			path := filepath.Join(abs, entry.Name())

			var desc string

			_ = db.DB.View(func(tx *bolt.Tx) (err error) {
				val := tx.Bucket(bucket).Get([]byte(path))

				if val == nil {
					desc = "Nooo"
				}

				desc = string(val)
				return
			})

			fmt.Println(entry.Name(), " -> ", desc)
		}

		return
	},
}
