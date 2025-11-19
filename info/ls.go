package info

import (
	"fmt"
	"os"
	"path/filepath"

	"charm.land/lipgloss/v2"
	"charm.land/lipgloss/v2/table"
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

		base := lipgloss.NewStyle().Padding(0, 1)
		header := base.Foreground(lipgloss.Color("252")).Bold(true)

		objects := [][]string{}

		for _, entry := range entries {
			abs, err := filepath.Abs(entry.Name())
			if err != nil {
				return err
			}

			path := filepath.Join(abs, entry.Name())

			var desc string

			_ = db.DB.View(func(tx *bolt.Tx) (err error) {
				val := tx.Bucket(bucket).Get([]byte(path))

				desc = string(val)

				return
			})

			if desc == "" {
				continue
			}

			objects = append(objects, []string{entry.Name(), desc})
		}

		t := table.New().
			Border(lipgloss.NormalBorder()).
			BorderStyle(lipgloss.NewStyle().Foreground(lipgloss.Color("238"))).
			Headers("Name", "Description").
			Width(40).
			Rows(objects...).
			StyleFunc(func(row, col int) lipgloss.Style {
				if row == table.HeaderRow {
					return header
				}
				return base.Foreground(lipgloss.Color("252"))
			})

		_, err = fmt.Println(t)
		return
	},
}
