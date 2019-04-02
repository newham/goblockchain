package test

import (
	"blockchain/core"
	"fmt"
	"testing"
)

func TestDB(t *testing.T) {
	db := core.NewBoltDB("", "")
	data1 := core.NewStringData("key", "Hello")
	db.Put(data1)
	fmt.Printf("k:%s,v:%s\n", data1.Key, db.GetValue(data1.Key))
	data2 := core.NewStringData("Hello", "World")
	db.Puts(data1, data2)
	fmt.Printf("k:%s,v:%s\n", data1.Key, db.GetValue(db.GetValue(data1.Key)))
}
