package crypto

import (
	"encoding/hex"
	"golang.org/x/crypto/blake2b"
)

func Blake256(data string) string {
	b, _ := blake2b.New(32, nil)
	s, _ := hex.DecodeString(data)
	b.Write(s)
	bs := b.Sum(nil)
	var result []byte = make([]byte, 64)
	hex.Encode(result, bs[:])
	return string(result)
}

func Blake256WithKey(data string, key []byte) string {
	b, _ := blake2b.New(32, key)
	s, _ := hex.DecodeString(data)
	b.Write(s)
	bs := b.Sum(nil)
	var result []byte = make([]byte, 64)
	hex.Encode(result, bs[:])
	return string(result)
}

func Blake128(data string) string {
	b, _ := blake2b.New(16, nil)
	s, _ := hex.DecodeString(data)
	b.Write(s)
	bs := b.Sum(nil)
	var result []byte = make([]byte, 32)
	hex.Encode(result, bs[:])
	return string(result)
}

func Blake128WithKey(data string, key []byte) string {
	b, _ := blake2b.New(16, key)
	s, _ := hex.DecodeString(data)
	b.Write(s)
	bs := b.Sum(nil)
	var result []byte = make([]byte, 32)
	hex.Encode(result, bs[:])
	return string(result)
}

func Blake160(data string) string {
	b, _ := blake2b.New(20, nil)
	s, _ := hex.DecodeString(data)
	b.Write(s)
	bs := b.Sum(nil)
	var result []byte = make([]byte, 40)
	hex.Encode(result, bs[:])
	return string(result)
}

func Blake160WithKey(data string, key []byte) string {
	b, _ := blake2b.New(20, key)
	s, _ := hex.DecodeString(data)
	b.Write(s)
	bs := b.Sum(nil)
	var result []byte = make([]byte, 40)
	hex.Encode(result, bs[:])
	return string(result)
}
