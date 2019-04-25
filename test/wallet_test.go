package test

import (
	"blockchain/core"
	"fmt"
	"testing"
)

func TestWallet(t *testing.T) {
	wallet := core.NewWallet("78dd859cbd77e30c.wallet")
	fmt.Printf("public:%x,private:%x\n", wallet.PublicKey, wallet.PrivateKey)
}

func TestWalletSave(t *testing.T) {
	wallet := core.NewWallet("")
	wallet.Save("")
}
