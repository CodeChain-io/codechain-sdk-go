package key

import (
	"encoding/hex"
	"testing"

	"github.com/CodeChain-io/codechain-rpc-go/crypto"
)

func TestEcdsa(t *testing.T) {
	key, err := GenerateKey()
	if err != nil {
		t.Fatal("Fail to generate private key")
	}

	msg, err := hex.DecodeString("aabbccdd0011223344aabbccdd0011223344aabbccdd0011223344aabbccdd00")
	if err != nil {
		t.Fatal("Fail to decode msg")
	}

	signature := crypto.SignEcdsa(msg, key)
	if signature == nil {
		t.Fatal("Signature is empty")
	}
}
