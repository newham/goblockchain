package core

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"time"
)

type Block struct {
	Timestamp     int64
	Data          []byte
	PrevBlockHash []byte
	Hash          []byte //上面三个数据的hash256
	Nonce         int    //挖矿所得的随机数
}

func (b Block) Print() {
	fmt.Printf("PrevHash: %x\n", b.PrevBlockHash)
	fmt.Printf("Data: %s\n", b.Data)
	fmt.Printf("Timestamp: %d\n", b.Timestamp)
	fmt.Printf("Nonce: %d\n", b.Nonce)
	fmt.Printf("Hash: %x\n\n", b.Hash)
}

func NewBlock(data string, preBlockHash []byte) *Block {
	block := &Block{
		Timestamp:     time.Now().UTC().Unix(), //UTC Time
		Data:          ToBytes(data),
		PrevBlockHash: preBlockHash,
	}

	//mine
	nonce, hash := NewProofOfWork(block).Work()
	block.Nonce = nonce
	block.Hash = hash

	return block
}

//func (b *Block) SetHash() {
//	timestamp := ToBytes(strconv.FormatInt(b.Timestamp, 10))
//	headers := BytesCombine(b.PrevBlockHash, b.Data, timestamp)
//	hash := sha256.Sum256(headers)
//	b.Hash = hash[:]
//}

func (b *Block) Serialize() []byte {
	buffer := bytes.NewBuffer(nil)
	encoder := gob.NewEncoder(buffer)
	if err := encoder.Encode(b); err != nil {
		return nil
	}
	return buffer.Bytes()
}

//创世块
func NewGenesisBlock() *Block {
	return NewBlock("Genesis Block", []byte{})
}
