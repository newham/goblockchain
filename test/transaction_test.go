package test

import (
	"blockchain/core"
	"fmt"
	"testing"
)

func TestNewCoinBaseTX(t *testing.T) {
	address := "jack"
	tx := core.NewCoinBaseTX(address, "")
	tx.Print()
	core.AddTransaction(tx)
	dbc.AddBlock(core.Serialize(core.LocalTransactions))
	fmt.Printf("%s: %f\n", address, core.Balance(dbc, address))
}

func TestTransaction(t *testing.T) {
	lastBlock := dbc.LastBlock() //末尾区块
	lastBlock.Print()
	//反序列化data到Transactions
	txs := core.UnSerializeTransactions(lastBlock.Data)
	//打印Transactions
	for _, tx := range txs {
		tx.Print()
	}
}

func TestNewTransaction(t *testing.T) {
	from := "tester"
	to := "jack"
	var amount float32 = 10.20101
	fmt.Printf("%s: %f\n", from, core.Balance(dbc, from))
	fmt.Printf("%s: %f\n", to, core.Balance(dbc, to))
	fmt.Printf("NewTransaction from %s to %s\n", from, to)
	core.NewTransaction(dbc, from, to, amount)
	dbc.AddBlock(core.Serialize(core.LocalTransactions))
	fmt.Printf("%s: %f\n", from, core.Balance(dbc, from))
	fmt.Printf("%s: %f\n", to, core.Balance(dbc, to))
}

func TestBalance(t *testing.T) {
	address1 := "tester"
	address2 := "jack"
	address3 := "lucy"
	fmt.Printf("%s: %f\n", address1, core.Balance(dbc, address1))
	fmt.Printf("%s: %f\n", address2, core.Balance(dbc, address2))
	fmt.Printf("%s: %f\n", address3, core.Balance(dbc, address3))
}
