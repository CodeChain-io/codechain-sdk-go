package crypto

import (
	"bytes"
	"encoding/hex"
	"testing"
)

func TestEcdsa(t *testing.T) {

	msg, _ := hex.DecodeString("aabbccdd0011223344aabbccdd0011223344aabbccdd0011223344aabbccdd00")
	priv, _ := hex.DecodeString("ffffaaaaffffaaaaffffaaaaffffaaaaffffaaaaffffaaaaffffaaaaffffaaaa")
	signature, _ := hex.DecodeString("64b85af63afed5d344ce8b809cc964d7b4ecd35a8dae8ae07909f5daf7b578441c63bebfa4b1beec84db4cb14578a80d3e3b5919a6702838b2954fd9d50c92af00")
	a := SignEcdsa(msg, priv)

	if bytes.Compare(a, signature) != 0 {
		t.Fatal("SignEcdsa Error")
	}
}
