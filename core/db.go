package core

import (
	"encoding/hex"
	"errors"
	"github.com/boltdb/bolt"
	"log"
	"time"
)

type DB interface {
	Put(data Data)
	Puts(dataList ...Data)
	GetValue(key []byte) []byte
	DeleteValue(key []byte)
	Close()
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

type mapDB struct {
	dataMap map[string][]byte
}

func (mdb *mapDB) Put(data Data) {
	mdb.dataMap[hex.EncodeToString(data.Key)] = data.Value
}

func (mdb *mapDB) Puts(dataList ...Data) {
	for _, data := range dataList {
		mdb.dataMap[hex.EncodeToString(data.Key)] = data.Value
	}
}

func (mdb *mapDB) GetValue(key []byte) []byte {
	return mdb.dataMap[hex.EncodeToString(key)]
}

func (mdb *mapDB) DeleteValue(key []byte) {
	mdb.dataMap[hex.EncodeToString(key)] = nil
}

func (mdb *mapDB) Close() {
	panic("implement me")
}

type boltDB struct {
	dbFile     string
	bucketName []byte
	db         *bolt.DB
}

func (bdb *boltDB) Close() {
	err := bdb.db.Close()
	if err != nil {
		log.Panic(err)
	}
}

func NewMapDB() DB {
	return &mapDB{dataMap: map[string][]byte{}}
}

/*
>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>
*/

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
	bdb := &boltDB{dbFile: dbFile, bucketName: []byte(bucketName)}
	var err error
	bdb.db, err = bolt.Open(dbFile, 0600, &bolt.Options{Timeout: 1 * time.Second})
	if err != nil {
		panic(err)
	}
	return bdb
}

func (bdb *boltDB) Put(data Data) {
	err := bdb.db.Update(func(tx *bolt.Tx) error {
		bucket, err := tx.CreateBucketIfNotExists(bdb.bucketName)
		if err != nil {
			return err
		}
		err = bucket.Put(data.Key, data.Value)
		return err
	})
	if err != nil {
		log.Panic(err)
	}
}

func (bdb *boltDB) Puts(dataList ...Data) {
	err := bdb.db.Update(func(tx *bolt.Tx) error {
		bucket, err := tx.CreateBucketIfNotExists(bdb.bucketName)
		if err != nil {
			return err
		}
		for _, data := range dataList {
			err = bucket.Put(data.Key, data.Value)
			if err != nil {
				return err
			}
		}
		return nil
	})
	if err != nil {
		log.Panic(err)
	}
}

func (bdb *boltDB) GetValue(key []byte) []byte {
	var value []byte
	err := bdb.db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket(bdb.bucketName)
		if bucket != nil {
			value = bucket.Get(key)
		} else {
			return errors.New("no exist")
		}
		return nil
	})
	if err != nil {
		return nil
	}
	return value
}

func (bdb *boltDB) DeleteValue(key []byte) {
}
