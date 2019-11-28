package primitives

import (
	"encoding/hex"
	"testing"
)

func TestPlatformAddress(t *testing.T) {
	id, err := hex.DecodeString("7b5e0ee8644c6f585fc297364143280a45844502")

	if err != nil {
		t.Fatal("hex decode error")
	}

	accountID := NewH160(id)
	p, err := PlatformAddressFromAccountID(accountID, "cc")

	if err != nil {
		t.Fatal("hex decode error")
	}

	if p.Value != "cccq9a4urhgv3xx7kzlc2tnvs2r9q9ytpz9qgs7q0a7" {
		t.Fatal("PlatformAddressFromAccountID error")
	}
}
