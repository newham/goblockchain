package core

type BlockChain interface {
	LastBlock() *Block
	AddBlock(data string)
	Blocks() []*Block
	Height() int
	List()
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
	if mbc.Height() < 1 {
		return nil
	}
	return mbc.Blocks()[mbc.Height()-1]
}

func (mbc *MemoryBlockChain) AddBlock(data string) {
	prevBlock := mbc.LastBlock()
	var newBlock *Block
	if prevBlock == nil {
		newBlock = NewGenesisBlock()
	} else {
		newBlock = NewBlock(data, prevBlock.Hash)
	}
	mbc.blocks = append(mbc.blocks, newBlock)
}

func (mbc *MemoryBlockChain) Blocks() []*Block {
	return mbc.blocks
}

func (mbc *MemoryBlockChain) Height() int {
	return len(mbc.Blocks())
}

func NewMemoryBlockChain() BlockChain {
	return &MemoryBlockChain{}
}

/**
Memory Block Chain
=>end
*/

type DbBlockChain struct {
}

func (dbc *DbBlockChain) LastBlock() *Block {
	panic("implement me")
}

func (dbc *DbBlockChain) AddBlock(data string) {
	panic("implement me")
}

func (dbc *DbBlockChain) Blocks() []*Block {
	panic("implement me")
}

func (dbc *DbBlockChain) Height() int {
	panic("implement me")
}

func (dbc *DbBlockChain) List() {
	panic("implement me")
}

func NewDbBlockChain() BlockChain {
	return &DbBlockChain{}
}