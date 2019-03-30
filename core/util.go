package core

import (
	"bytes"
	"math/big"
)

func ToBytes(str string) []byte {
	return []byte(str)
}

//将int64转换为字节数组
func IntToHex(num int64) []byte {
	return big.NewInt(num).Bytes()
}

func BytesCombine(pBytes ...[]byte) []byte {
	return bytes.Join(pBytes, []byte{})
}
