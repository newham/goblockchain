package core

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"log"
)

const (
	Subsidy   float32 = 10      //挖矿的奖励
	MinAmount float32 = 0.00001 //交易的最低金额
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
	TxId      []byte //引用之前的交易Id，Transaction.Id
	VOutId    int    //Transaction.VOut的id索引
	ScriptSig string //输入地址
}

//输出
type TxOutput struct {
	Value        float32 //输出金额
	ScriptPubKey string  //目的地址的公钥，目的地址可以用私钥解开
}

func (in *TxInput) Unlock(address string) bool {
	return in.ScriptSig == address
}

func (out *TxOutput) Unlock(address string) bool {
	return out.ScriptPubKey == address
}

func (t *Transaction) Print() {
	fmt.Printf("Transactions:\n")
	fmt.Printf("\tId:%x\n", t.Id)
	fmt.Printf("\tVIn:%x\n", t.VIn)
	fmt.Printf("\tVOut:%x\n", t.VOut)
}

func (t *Transaction) SetId() {
	hash := sha256.Sum256(Serialize(t))
	t.Id = hash[:]
}

func (t *Transaction) Add(from, to string) {

}

func (t *Transaction) IsCoinBase() bool {
	return len(t.VIn) == 1 && len(t.VIn[0].TxId) == 0 && t.VIn[0].VOutId == -1
}

func NewCoinBaseTX(to, data string) Transaction { //to:目的用户地址，data:说明
	if data == "" {
		data = fmt.Sprintf("Reward to '%s'", to)
	}
	txInput := TxInput{[]byte{}, -1, data}
	txOutput := TxOutput{Subsidy, to}
	tx := Transaction{nil, []TxInput{txInput}, []TxOutput{txOutput}}
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

/**
创建交易
*/
func NewTransaction(bc *BlockChain, from, to string, amount float32) {
	if amount < MinAmount {
		log.Panic("Error: amount <", MinAmount)
	}
	var txInputs []TxInput                                              //输入
	var txOutputs []TxOutput                                            //输出
	account, spendableOutputs := FindSpendableOutputs(bc, from, amount) // 得到可用的输出
	if account < amount {
		log.Panic("Error: funds not enough")
	}
	for id, outs := range spendableOutputs {
		txId, err := hex.DecodeString(id)
		if err != nil {
			log.Panic(err)
		}
		for _, out := range outs {
			txInput := TxInput{txId, out, from}
			txInputs = append(txInputs, txInput)
		}
	}

	txOutputs = append(txOutputs, TxOutput{amount, to})

	//找零
	if account > amount {
		txOutputs = append(txOutputs, TxOutput{account - amount, from})
	}
	tx := Transaction{nil, txInputs, txOutputs}
	tx.SetId()

	AddTransaction(tx)
}

//寻找未花费的交易
func FindUnspentTransactions(bc *BlockChain, address string) []Transaction {
	var unspentTXs []Transaction
	spentTXOs := map[string][]int{} //已经花费的outId
	bci := bc.Iterator()
	//开始遍历整个链
	bci.While(func(block *Block) bool {
		for _, tx := range UnSerializeTransactions(block.Data) { //检查该区块中的每笔交易
			txId := hex.EncodeToString(tx.Id) //交易Id

		Outputs:
			for outId, out := range tx.VOut { //检查交易的输出是否已经作为输入（被花费）

				for _, spentOut := range spentTXOs[txId] {
					if spentOut == outId { //如果已经花费，则继续循环下一笔交易
						continue Outputs
					}
				}

				if out.Unlock(address) { //可用的交易
					unspentTXs = append(unspentTXs, tx)
				}
			}
			if !tx.IsCoinBase() { //将已经花费的TxId加入到
				for _, in := range tx.VIn {
					if in.Unlock(address) {
						inTxId := hex.EncodeToString(in.TxId)
						spentTXOs[inTxId] = append(spentTXOs[inTxId], in.VOutId)
					}
				}
			}
		}
		return true
	})
	return unspentTXs
}

func FindTxOutputs(bc *BlockChain, address string) []TxOutput {
	var txOutputs []TxOutput
	unspentTransactions := FindUnspentTransactions(bc, address)

	for _, tx := range unspentTransactions {
		for _, out := range tx.VOut {
			if out.Unlock(address) {
				txOutputs = append(txOutputs, out)
			}
		}
	}
	return txOutputs
}

func FindSpendableOutputs(bc *BlockChain, address string, amount float32) (float32, map[string][]int) {
	unspentOutputs := map[string][]int{}
	unspentTxs := FindUnspentTransactions(bc, address)
	var accumulated float32 = 0

Work:
	for _, tx := range unspentTxs {
		txId := hex.EncodeToString(tx.Id)
		for outId, out := range tx.VOut {
			if out.Unlock(address) && accumulated < amount {
				accumulated += out.Value
				unspentOutputs[txId] = append(unspentOutputs[txId], outId)
				if accumulated >= amount {
					break Work
				}
			}
		}
	}
	return accumulated, unspentOutputs
}

func Balance(bc *BlockChain, address string) float32 {
	var balance float32
	txOutputs := FindTxOutputs(bc, address) //找出所有未花费的输出
	for _, txOutput := range txOutputs {
		balance += txOutput.Value
	}
	return Decimal(balance)
}
