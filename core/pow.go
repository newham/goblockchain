package core

import (
	"crypto/sha256"
	"math"
	"math/big"
)

/**
proof of work
*/
type PoW struct {
	block  *Block
	target *big.Int
}

var maxNonce = math.MaxInt64

var difficulty = 8

func (pow *PoW) Work() (int, []byte) {

	nonce := 0
	var hash []byte
	isValidate := false
	for nonce < maxNonce {
		isValidate, hash = pow.Validate(nonce)
		if isValidate {
			break
		} else {
			nonce++
		}
	}
	return nonce, hash
}

func (pow *PoW) Validate(nonce int) (bool, []byte) {
	var hashInt big.Int
	hash := pow.BlockDataHash(nonce)
	hashInt.SetBytes(hash)
	isValid := hashInt.Cmp(pow.target) == -1 // hashInt <  target ，即hash的前targetBits位为0
	return isValid, hash
}

func (pow *PoW) prepareData(nonce int) []byte {
	return BytesCombine(
		pow.block.PrevBlockHash,
		pow.block.Data,
		IntToHex(pow.block.Timestamp),
		IntToHex(int64(difficulty)),
		IntToHex(int64(nonce)),
	)
}

func (pow *PoW) BlockDataHash(nonce int) []byte {
	hash := sha256.Sum256(pow.prepareData(nonce))
	return hash[:]
}

func NewProofOfWork(block *Block) *PoW {
	target := big.NewInt(0)
	target.SetBit(target, 256-difficulty, 1) // 将第256-difficulty 位设为1，之后为0
	//target.Lsh(target, uint(256-difficulty)) // 左移（256-difficulty）位，得到一个前targetBits为0，形如00000...010000000...00的字节码
	return &PoW{block, target}
}

func Lsh() {
	x := big.NewInt(0)
	x.SetBit(x, 256, 1)
}
