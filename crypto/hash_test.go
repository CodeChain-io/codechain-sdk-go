package crypto

import (
	"bytes"
	"encoding/hex"
	"testing"
)

func TestBlake(t *testing.T) {
	input, _ := hex.DecodeString("abcd")
	r1, _ := Blake256(input)
	a1, _ := hex.DecodeString("9606e52f00c679e548b5155af5026f5af4130d7a15c990a791fff8d652c464f5")
	if 0 != bytes.Compare(r1, a1) {
		t.Fatal("Blake256 function output error")
	}

	key1 := []byte{1, 2, 3, 4}
	r2, _ := Blake128WithKey(input, key1)
	a2, _ := hex.DecodeString("cc98decf76abbbdf3a2027552bec09c8")
	if 0 != bytes.Compare(r2, a2) {
		t.Fatal("Blake128WithKey function output error")
	}

	input2, _ := hex.DecodeString("aabbccddeeff0011223344")
	key2 := []byte{33, 65, 97, 34, 35, 36}
	r3, _ := Blake160WithKey(input2, key2)
	a3, _ := hex.DecodeString("68e2afd80441be243c755c2184eb4c9ce9f9f17c")
	if 0 != bytes.Compare(r3, a3) {
		t.Fatal("Blake160WithKey function output error")
	}

}

func TestRipemd160(t *testing.T) {
	input, _ := hex.DecodeString("abcdef12345678")
	rawHash, _ := Ripemd160(input)
	r := hex.EncodeToString(rawHash)
	a := "6b1f9162346b8962edde6e28ee7599541def053e"
	if r != a {
		t.Fatal("Ripemd160 function error")
	}
}
