package test

import (
	"blockchain/core"
	"fmt"
	"testing"
)

func TestNewBlock(t *testing.T) {
	b := core.NewBlock([]byte(""), core.ToBytes(""))
	println("hash", b.Hash)
}

func TestSerialize(t *testing.T) {
	block1 := core.NewBlock([]byte("hi"), nil)
	data := core.Serialize(block1)
	fmt.Printf("Serialize data:%x\n", data)
	block2 := &core.Block{}
	core.UnSerialize(block2, data)
	block2.Print()
}
