package test

import (
	"blockchain/core"
	"testing"
)

func TestNewBlockChain(t *testing.T) {
	bc := core.NewMemoryBlockChain()
	bc.AddBlock("Send 1 BTC to Tom")
	bc.AddBlock("Send 2 BTC to Tom")
	println()
	for _, b := range bc.Blocks() {
		b.Print()
	}

}
