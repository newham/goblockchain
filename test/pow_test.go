package test

import (
	"blockchain/core"
	"testing"
)

func TestProofOfWork(t *testing.T) {
	pow := core.NewProofOfWork(core.NewGenesisBlock("tester"))
	pow.Work()
}

func TestLsh(t *testing.T) {
	core.Lsh()
}
