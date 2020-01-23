package main

import (
	"fmt"
	"os"

	"errors"

	"encoding/json"

	"github.com/boltdb/bolt"
)

const (
	BUCKET_NAME = "autotiling"
)

type store struct {
	dbPath string
	dbConn *bolt.DB
}

func (s *store) put(key []byte, data interface{}) error {
	b, err := json.Marshal(data)
	if err != nil {
		return err
	}

	err = s.dbConn.Update(func(tx *bolt.Tx) error {
		bucket, err := tx.CreateBucketIfNotExists([]byte(BUCKET_NAME))
		if err != nil {
			return err
		}

		err = bucket.Put(key, b)
		if err != nil {
			return err
		}

		return nil
	})

	return err
}

func (s *store) get(key []byte) ([]byte, error) {
	var data []byte
	err := s.dbConn.View(func(tx *bolt.Tx) error {
		var err error

		bucket := tx.Bucket([]byte(BUCKET_NAME))
		if bucket == nil {
			err = fmt.Errorf("Bucket not created yet!")
		}
		data = bucket.Get(key)

		return err
	})

	return data, err
}

func newStore() (*store, error) {
	store := &store{}

	pathes := []string{os.Getenv("XDG_CONFIG_HOME"), os.Getenv("HOME")}
	for _, pathPrefix := range pathes {
		if pathPrefix != "" {
			store.dbPath = pathPrefix + "/.autotiling.bolt"
			break
		}
	}

	if store.dbPath == "" {
		return store, errors.New("")
	}

	if err := store.openDB(); err != nil {
		return store, err
	}

	return store, nil
}

func (s *store) openDB() error {
	err := s.createDB()
	if err != nil {
		return err
	}

	db, err := bolt.Open(s.dbPath, 0600, nil)
	if err != nil {
		return err
	}
	s.dbConn = db

	return nil
}

func (s *store) createDB() error {
	db, err := bolt.Open(s.dbPath, 0600, nil)
	if err != nil {
		return err
	}
	defer db.Close()

	err = db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists([]byte(BUCKET_NAME))
		if err != nil {
			return err
		}

		return nil
	})

	return err
}
