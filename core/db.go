package core

import (
	"github.com/boltdb/bolt"
	"log"
	"time"
)

type DB interface {
	Put(data Data) error
	Puts(dataList ...Data) error
	GetValue(key []byte) []byte
	DeleteValue(key []byte) error
	Close() error
}

type Data struct {
	Key   []byte
	Value []byte
}

func NewData(key []byte, value []byte) Data {
	return Data{key, value}
}

func NewStringData(key, value string) Data {
	return Data{[]byte(key), []byte(value)}
}

type BoltDB struct {
	dbFile     string
	bucketName []byte
	db         *bolt.DB
}

func (bdb *BoltDB) Close() error {
	return bdb.db.Close()
}

const (
	defaultDbFile     = "../db/blockchain.db"
	defaultBucketName = "blocks"
)

func NewBoltDB(dbFile string, bucketName string) DB {
	if dbFile == "" {
		dbFile = defaultDbFile
	}
	if bucketName == "" {
		bucketName = defaultBucketName
	}
	bdb := &BoltDB{dbFile: dbFile, bucketName: []byte(bucketName)}
	var err error
	bdb.db, err = bolt.Open(dbFile, 0600, &bolt.Options{Timeout: 1 * time.Second})
	if err != nil {
		panic(err)
	}
	return bdb
}

func (bdb *BoltDB) Put(data Data) error {
	return bdb.db.Update(func(tx *bolt.Tx) error {
		bucket, err := tx.CreateBucketIfNotExists(bdb.bucketName)
		if err != nil {
			log.Panic(err)
		}
		err = bucket.Put(data.Key, data.Value)
		if err != nil {
			log.Panic(err)
		}
		return err
	})
}

func (bdb *BoltDB) Puts(dataList ...Data) error {
	return bdb.db.Update(func(tx *bolt.Tx) error {
		bucket, err := tx.CreateBucketIfNotExists(bdb.bucketName)
		if err != nil {
			log.Panic(err)
		}
		for _, data := range dataList {
			err = bucket.Put(data.Key, data.Value)
			if err != nil {
				log.Panic(err)
			}
		}
		return err
	})
}

func (bdb *BoltDB) GetValue(key []byte) []byte {
	var value []byte
	bdb.db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket(bdb.bucketName)
		if bucket != nil {
			value = bucket.Get(key)
		} else {
			value = nil
		}
		return nil
	})
	return value
}

func (bdb *BoltDB) DeleteValue(key []byte) error {
	return nil
}
