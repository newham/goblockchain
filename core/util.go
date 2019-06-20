package core

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"math/big"
	"strconv"
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

//序列化struct
func Serialize(obj interface{}) []byte {
	buffer := bytes.NewBuffer(nil)
	encoder := gob.NewEncoder(buffer)
	if err := encoder.Encode(obj); err != nil {
		return nil
	}
	return buffer.Bytes()
}

//反序列化struct，使用.(*XXX [struct])将interface转成struct
func UnSerialize(obj interface{}, data []byte) interface{} {
	buffer := bytes.NewBuffer(data)
	decoder := gob.NewDecoder(buffer)
	if err := decoder.Decode(obj); err != nil {
		return nil
	}
	return obj
}

//保留5位小数
func Decimal(value float32) float32 {
	v, _ := strconv.ParseFloat(fmt.Sprintf("%.5f", value), 32)
	return float32(v)
}
