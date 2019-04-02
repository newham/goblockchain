package core

import (
	"crypto/sha256"
	"fmt"
)

/**
比特币的交易
*/
type Transaction struct {
	Id   []byte
	VIn  []TxInput  //所有输入
	VOut []TxOutput //所有输出
}

//输入
type TxInput struct {
	TxId      []byte
	VOut      float32
	ScriptSig string
}

//输出
type TxOutput struct {
	Value        float32
	ScriptPubKey string
}

func (t *Transaction) Print() {
	fmt.Printf("Id:%x\n", t.Id)
	fmt.Printf("VIn:%x\n", t.VIn)
	fmt.Printf("VOut:%x\n", t.VOut)
}

func (t *Transaction) SetId() {
	hash := sha256.Sum256(Serialize(t))
	t.Id = hash[:]
}

func (t *Transaction) Add(from, to string) {

}

var subsidy float32 = 10

func NewCoinBaseTX(to, data string) *Transaction { //to:目的用户地址，data:说明
	if data == "" {
		data = fmt.Sprintf("Reward to '%s'", to)
	}
	txInput := TxInput{[]byte{}, -1, data}
	txOutput := TxOutput{subsidy, to}
	tx := &Transaction{nil, []TxInput{txInput}, []TxOutput{txOutput}}
	tx.SetId()
	return tx
}

func UnSerializeTransactions(data []byte) []Transaction {
	var txs []Transaction
	UnSerialize(&txs, data) //千万别忘了这个取地址符【&】，否则无法反序列化
	return txs
}

var LocalTransactions []Transaction

func AddTransaction(transaction Transaction) {
	LocalTransactions = append(LocalTransactions, transaction)
}

func NewTransaction(from, to string, amount float32, bc BlockChain) *Transaction {
	//var inputs []TxInput
	//var outputs []TxOutput
	//for _, block := range bc.Blocks() {
	//
	//}
	return nil
}

func FindSpendableOutputs(address string, bc BlockChain) []TxOutput {
	var outputs []TxOutput
	bc.Find(func(block *Block) bool {
		transactions := UnSerializeTransactions(block.Data)
		for _, transaction := range transactions {
			for _, input := range transaction.VIn {
				if input.ScriptSig == address {
					outputs = append(outputs)
				} else {

				}
			}
		}
		return true
	})
	return nil
}

func Budget(bc BlockChain, address string) float32 {
	var budget float32 = 0
	for _, block := range bc.Blocks() {
		transactions := UnSerializeTransactions(block.Data)
		for _, transaction := range transactions {
			for _, out := range transaction.VOut {
				if out.ScriptPubKey == address {
					budget += out.Value
				}
			}
			for _, in := range transaction.VIn {
				if in.ScriptSig == address {
					budget -= in.VOut
				}
			}
		}
	}
	return budget
}
