package core

import (
	"bytes"
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"encoding/gob"
	"encoding/hex"
	"io/ioutil"
	"log"
)

type Wallet struct {
	PublicKey  []byte
	PrivateKey *ecdsa.PrivateKey
}

func (w *Wallet) Address() []byte {
	return nil
}

func (w *Wallet) Save(fileName string) error {
	if fileName == "" {
		fileName = hex.EncodeToString(w.PublicKey)[:16] + ".wallet"
	}
	buffer := bytes.NewBuffer(nil)
	gob.Register(elliptic.P256())
	gob.NewEncoder(buffer).Encode(w)
	return ioutil.WriteFile(fileName, buffer.Bytes(), 0644)
}

func NewWallet(fileName string) *Wallet {
	wallet := &Wallet{}
	if fileName != "" {
		b, err := ioutil.ReadFile(fileName)
		if err != nil {
			log.Panic(err)
		}
		buffer := bytes.NewBuffer(b)
		gob.Register(elliptic.P256())
		err = gob.NewDecoder(buffer).Decode(wallet)
		if err != nil {
			log.Panic(err)
		}
	} else {
		curve := elliptic.P256()
		private, err := ecdsa.GenerateKey(curve, rand.Reader)
		if err != nil {
			log.Panic(err)
		}
		public := append(private.PublicKey.X.Bytes(), private.PublicKey.Y.Bytes()...)
		wallet.PrivateKey = private
		wallet.PublicKey = public
	}

	return wallet
}
