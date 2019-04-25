package test

import (
	"blockchain/core"
	"fmt"
	"testing"
)

func TestNewCoinBaseTX(t *testing.T) {
	tx := core.NewCoinBaseTX("tester", "")
	tx.Print()
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
	to := "lucy"
	var amount float32 = 0.1
	fmt.Printf("%s: %f\n", from, core.Balance(dbc, from))
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
