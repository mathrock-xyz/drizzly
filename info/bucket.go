package info

import (
	"github.com/mathrock-xyz/drizzly/db"
	"go.etcd.io/bbolt"
)

var bucket = []byte("info")

func init() {
	db.DB.Update(func(tx *bbolt.Tx) (err error) {
		_, err = tx.CreateBucketIfNotExists(bucket)
		return
	})
}
