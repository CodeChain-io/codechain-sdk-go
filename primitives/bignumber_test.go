package primitives

import (
	"bytes"
	"testing"
)

func TestBigNumber(t *testing.T) {
	a := NewU128("1000112341234123413241324234")
	b := NewU128("5")
	c := NewU256("0")
	if bytes.Compare(a.Bytes(), []byte{3, 59, 70, 6, 167, 159, 170, 29, 70, 13, 226, 202}) != 0 {
		t.Fatal("bignumber creation Error")
	}

	if bytes.Compare(b.ToEncodeObject(), []byte{5}) != 0 {
		t.Fatal("ToEncodeObject Error")
	}
	if bytes.Compare(c.ToEncodeObject(), []byte{}) != 0 {
		t.Fatal("ToEncodeObject Error")
	}
}
