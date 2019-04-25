package test

import (
	"blockchain/core"
	"fmt"
	"math/rand"
	"testing"
)

var dbc = core.NewDbBlockChain("tester", "", "")

func TestNewBlockChain(t *testing.T) {
	bc := core.NewMemoryBlockChain("tester")
	bc.AddBlock([]byte("Send 1 BTC to Tom"))
	bc.AddBlock([]byte("Send 2 BTC to Tom"))
	println()
	for _, b := range bc.Blocks() {
		b.Print()
	}
}

func TestPrint(t *testing.T) {
	dbc.Print()
}

func TestLastBlock(t *testing.T) {
	dbc.LastBlock().Print()
}

func TestAddBlock(t *testing.T) {
	dbc.AddBlock([]byte(fmt.Sprintf("%d-%s", rand.Intn(100), "test data")))
	dbc.Print()
}

func TestBlocks(t *testing.T) {
	for _, b := range dbc.Blocks() {
		b.Print()
	}
}

func TestNext(t *testing.T) {
	bci := dbc.Iterator()
	block := bci.Next()
	for block != nil {
		block.Print()
		txs := core.UnSerializeTransactions(block.Data)
		for _, tx := range txs {
			tx.Print()
		}
		println(">>>>>>>>>")
		block = bci.Next()
	}

}
