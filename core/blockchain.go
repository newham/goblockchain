package core

import "log"

type BlockChain interface {
	LastBlock() *Block
	AddBlock(data []byte)
	Blocks() []*Block
	List()
	Find(f func(*Block) bool)
}

/**
Memory Block Chain
=>start
*/
type MemoryBlockChain struct {
	BlockChain
	blocks []*Block
}

func (mbc *MemoryBlockChain) List() {
	for _, block := range mbc.blocks {
		block.Print()
	}
}

func (mbc *MemoryBlockChain) LastBlock() *Block {
	height := len(mbc.blocks)
	if height < 1 {
		return nil

	}
	return mbc.Blocks()[height-1]
}

func (mbc *MemoryBlockChain) AddBlock(data []byte) {
	//添加新区块
	mbc.blocks = append(mbc.blocks, NewBlock(data, mbc.LastBlock().Hash))
}

func (mbc *MemoryBlockChain) Blocks() []*Block {
	return mbc.blocks
}

func (mbc *MemoryBlockChain) Find(f func(*Block) bool) {

}

func NewMemoryBlockChain(address string) BlockChain {
	return &MemoryBlockChain{blocks: []*Block{NewGenesisBlock(address)}} //创建区块链时，创建创世区块
}

/**
Memory Block Chain
=>end
*/

/**
DB Block Chain
=>start
*/

type DbBlockChain struct {
	tip []byte
	db  DB
}

func (dbc *DbBlockChain) Find(f func(*Block) bool) {
	block := dbc.LastBlock()
	for f(block) {
		if block.PrevBlockHash != nil {
			block = UnSerializeBlock(dbc.db.GetValue(block.PrevBlockHash))
		} else {
			break
		}
	}
}

func (dbc *DbBlockChain) LastBlock() *Block {
	return UnSerializeBlock(dbc.db.GetValue(dbc.tip))
}

func (dbc *DbBlockChain) AddBlock(data []byte) {
	newBlock := NewBlock(data, dbc.LastBlock().Hash)
	dbc.db.Put(NewData(newBlock.Hash, Serialize(newBlock))) // save this block to DB
	dbc.setTip(newBlock.Hash)                               //set last hash = this block's hash
}

func (dbc *DbBlockChain) Blocks() []*Block {
	var blocks []*Block
	block := dbc.LastBlock()
	for true {
		blocks = append(blocks, block)
		if block.PrevBlockHash != nil {
			block = UnSerializeBlock(dbc.db.GetValue(block.PrevBlockHash))
		} else {
			break
		}
	}
	return blocks
}

func (dbc *DbBlockChain) List() {
	for _, block := range dbc.Blocks() {
		block.Print()
	}
}

const lastHashKey = "lastHash"

func (dbc *DbBlockChain) setTip(hash []byte) {
	dbc.db.Put(NewData([]byte(lastHashKey), hash))
	dbc.tip = hash
}

func (dbc *DbBlockChain) getTip() []byte {
	return dbc.db.GetValue([]byte(lastHashKey))
}

func NewDbBlockChain(address, dbFile, bucketName string) BlockChain {
	if address == "" {
		log.Panic("Need a address to create a block chain")
	}
	dbc := &DbBlockChain{db: NewBoltDB(dbFile, bucketName)}
	lastHash := dbc.getTip()
	if lastHash == nil { //如果区块链为空，则创建创世块
		genesisBlock := NewGenesisBlock(address)
		//l保存最后一个区块的hash
		dbc.db.Put(NewData(genesisBlock.Hash, Serialize(genesisBlock)))
		dbc.setTip(genesisBlock.Hash)
	} else {
		dbc.tip = lastHash
	}
	return dbc
}

/**
DB Block Chain
=>end
*/
