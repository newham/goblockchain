package core

import (
	"encoding/json"
	"fmt"
	"time"
)

type Block struct {
	Timestamp     int64
	Data          []byte
	PrevBlockHash []byte
	NBit          int
	Hash          []byte //上面三个数据的hash256
	Nonce         int    //挖矿所得的随机数
}

func (b *Block) Bytes() []byte {
	bts, _ := json.Marshal(b)
	return bts
}

func (b *Block) Print() {
	fmt.Printf("PrevHash: %x\n", b.PrevBlockHash)
	if len(b.Data) > 32 {
		fmt.Printf("Data: %x...\n", b.Data[:32])
	} else {
		fmt.Printf("Data: %x\n", b.Data)
	}
	fmt.Printf("Timestamp: %d\n", b.Timestamp)
	fmt.Printf("Nonce: %d\n", b.Nonce)
	fmt.Printf("Hash: %x\n", b.Hash)
}

func NewBlock(data []byte, preBlockHash []byte, nBit int) *Block {
	block := &Block{
		Timestamp:     time.Now().UTC().Unix(), //UTC Time
		Data:          data,
		PrevBlockHash: preBlockHash,
		NBit:          nBit,
	}

	//mine
	nonce, hash := NewProofOfWork(block).Work()
	block.Nonce = nonce
	block.Hash = hash

	return block
}

func UnSerializeBlock(data []byte) *Block {
	return UnSerialize(NewBlock(nil, nil, maxDifficulty), data).(*Block)
}

/**
序列化
*/
//func (b *Block) Serialize() []byte {
//	buffer := bytes.NewBuffer(nil)
//	encoder := gob.NewEncoder(buffer)
//	if err := encoder.Encode(b); err != nil {
//		return nil
//	}
//	return buffer.Bytes()
//}

//func (b *Block) UnSerialize(data []byte) *Block {
//	buffer := bytes.NewBuffer(data)
//	decoder := gob.NewDecoder(buffer)
//	if err := decoder.Decode(b); err != nil {
//		return nil
//	}
//	return b
//}

//创世块
func NewGenesisBlock(address string, nBit int) *Block {
	return NewBlock(Serialize([]Transaction{NewCoinBaseTX(address, "")}), nil, nBit)
}
