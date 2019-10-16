package crypto

import (
	"testing"
)

func TestBlake(t *testing.T) {
	r1 := Blake256("abcd")
	if r1 != "9606e52f00c679e548b5155af5026f5af4130d7a15c990a791fff8d652c464f5" {
		t.Fatal("Blake256 function output error")
	}

	key1 := []byte{1, 2, 3, 4}
	r2 := Blake128WithKey("abcd", key1)
	if r2 != "cc98decf76abbbdf3a2027552bec09c8" {
		t.Fatal("Blake128WithKey function output error")
	}

	key2 := []byte{33, 65, 97, 34, 35, 36}
	r3 := Blake160WithKey("aabbccddeeff0011223344", key2)
	if r3 != "68e2afd80441be243c755c2184eb4c9ce9f9f17c" {
		t.Fatal("Blake160WithKey function output error")
	}

	r4 := Ripemd160("abcdef12345678")
	if r4 != "6b1f9162346b8962edde6e28ee7599541def053e" {
		t.Fatal("Ripemd160 function error")
	}
}
