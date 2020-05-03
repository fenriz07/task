package bd

import (
	"fmt"

	bolt "go.etcd.io/bbolt"
)

func setupDB() (*bolt.DB, error) {
	db, err := bolt.Open("test.db", 0600, nil)
	if err != nil {
		return nil, fmt.Errorf("could not open db, %v", err)
	}
	err = db.Update(func(tx *bolt.Tx) error {

		_, err := tx.CreateBucketIfNotExists([]byte("TODO"))

		if err != nil {
			return fmt.Errorf("could not create todo bucket: %v", err)
		}

		return nil
	})
	if err != nil {
		return nil, fmt.Errorf("could not set up buckets, %v", err)
	}
	return db, nil
}

func AddorUpdate(task string, active string) {

	db, _ := setupDB()

	defer db.Close()

	_ = db.Update(func(tx *bolt.Tx) error {
		err := tx.Bucket([]byte("TODO")).Put([]byte(task), []byte(active))
		if err != nil {
			return fmt.Errorf("could not insert weight: %v", err)
		}
		return nil
	})
}

func Show() map[string]string {

	list := make(map[string]string)
	db, _ := setupDB()

	defer db.Close()

	db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("TODO"))
		b.ForEach(func(k, v []byte) error {
			//fmt.Println(string(k), string(v))
			list[string(k)] = string(v)
			return nil
		})
		return nil
	})

	return list
}
