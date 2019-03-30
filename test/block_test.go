package test

import (
	"blockchain/core"
	"fmt"
	"testing"
)

func TestNewBlock(t *testing.T) {
	b := core.NewBlock("", core.ToBytes(""))
	println("hash", b.Hash)
}

func TestSerialize(t *testing.T) {
	b := core.NewBlock("hi", nil)
	fmt.Printf("%x", b.Serialize())
}
