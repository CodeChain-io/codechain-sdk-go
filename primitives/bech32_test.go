package primitives

import (
	"bytes"
	"encoding/hex"
	"testing"
)

func TestBech32(t *testing.T) {
	hex, _ := hex.DecodeString("aabbccddeeff012a3b4d5e6f9922")
	r := bech32Encode("wcc", toWords(hex))
	if r != "wcc42aueh0wluqj5w6dtehejgsv7dgs8" {
		t.Fatal("Bech32 Encode Error")
	}

	dec := fromWords(bech32Decode(r, "wcc"))
	if bytes.Compare(dec, hex) != 0 {
		t.Fatal("Bech32 Decode Error")
	}
}
