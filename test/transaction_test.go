package test

import (
	"blockchain/core"
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

func TestBudget(t *testing.T) {
	//budget := core.Budget(dbc, "tester")
	//fmt.Printf("tester's Budget:%f\n", budget)
}
