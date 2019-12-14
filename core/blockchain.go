package core

import "log"

type BlockChain struct {
	tip []byte
	db  DB
}

func (bc *BlockChain) LastBlock() *Block {
	return UnSerializeBlock(bc.db.GetValue(bc.tip))
}

func (bc *BlockChain) AddBlock(data []byte) {
	newBlock := NewBlock(data, bc.LastBlock().Hash, maxDifficulty)
	bc.db.Put(NewData(newBlock.Hash, Serialize(newBlock))) // save this block to DB
	bc.setTip(newBlock.Hash)                               //set last hash = this block's hash
}

func (bc *BlockChain) Blocks() []*Block {
	var blocks []*Block
	block := bc.LastBlock()
	for true {
		blocks = append(blocks, block)
		if block.PrevBlockHash != nil {
			block = UnSerializeBlock(bc.db.GetValue(block.PrevBlockHash))
		} else {
			break
		}
	}
	return blocks
}

func (bc *BlockChain) Print() {
	for _, block := range bc.Blocks() {
		block.Print()
	}
}

func (bc *BlockChain) Iterator() *BlockChainIterator {
	return &BlockChainIterator{bc.tip, bc}
}

const lastHashKey = "lastHash"

func (bc *BlockChain) setTip(hash []byte) {
	bc.db.Put(NewData([]byte(lastHashKey), hash))
	bc.tip = hash
}

func (bc *BlockChain) getTip() []byte {
	return bc.db.GetValue([]byte(lastHashKey))
}

func NewMemoryBlockChain(address string) *BlockChain {
	bc := &BlockChain{db: NewMapDB()} //创建区块链时，创建创世区块
	genesisBlock := NewGenesisBlock(address, maxDifficulty)
	bc.db.Put(NewData(genesisBlock.Hash, Serialize(genesisBlock)))
	bc.setTip(genesisBlock.Hash)
	return bc
}

func NewDbBlockChain(address, dbFile, bucketName string) *BlockChain {
	if address == "" {
		log.Panic("Need a address to create a block chain")
	}
	bc := &BlockChain{db: NewBoltDB(dbFile, bucketName)}
	lastHash := bc.getTip()
	if lastHash == nil { //如果区块链为空，则创建创世块
		genesisBlock := NewGenesisBlock(address, maxDifficulty)
		//l保存最后一个区块的hash
		bc.db.Put(NewData(genesisBlock.Hash, Serialize(genesisBlock)))
		bc.setTip(genesisBlock.Hash)
	} else {
		bc.tip = lastHash
	}
	return bc
}

//区块链遍历器
type BlockChainIterator struct {
	currentHash []byte
	bc          *BlockChain
}

func (dci *BlockChainIterator) Next() *Block {
	if dci.currentHash == nil {
		return nil
	}
	block := UnSerializeBlock(dci.bc.db.GetValue(dci.currentHash))
	dci.currentHash = block.PrevBlockHash
	return block
}

//While循环遍历
func (dci *BlockChainIterator) While(f func(*Block) bool) {
	for {
		b := dci.Next()
		if b == nil || !f(b) {
			break
		}
	}

}
