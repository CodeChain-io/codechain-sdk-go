package primitives

import (
	"bytes"
	"testing"
)

func TestHexstring(t *testing.T) {
	a := NewH128Zero()
	b := NewH128Zero()
	c := H128([16]byte{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0})
	if a.Cmp(b) != true {
		t.Fatal("Hexstring Error")
	}
	if c.Cmp(b) != false {
		t.Fatal("Hexstring Error")
	}
	if c.ToString() != "01000000000000000000000000000000" {
		t.Fatal("Hexstring ToString() Error")
	}

	if bytes.Compare(a.RlpBytes(), []byte{144, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 00}) != 0 {
		t.Fatal("Hexstring RlpBytes() Error")
	}
	if c.ToJSON() != "0x01000000000000000000000000000000" {
		t.Fatal("Hexstring ToJSON() Error")
	}

	z := "0x1010101010101010101010101010101010101010"
	x, y := StringToH160(z)

	if y != nil || bytes.Compare(x.Bytes(), []byte{16, 16, 16, 16, 16, 16, 16, 16, 16, 16, 16, 16, 16, 16, 16, 16, 16, 16, 16, 16}) != 0 {
		t.Fatal("StringToH160 Error")
	}
}
