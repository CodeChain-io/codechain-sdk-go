package key

import (
	"encoding/hex"
	"testing"

	"github.com/CodeChain-io/codechain-sdk-go/crypto"
)

func TestEcdsa(t *testing.T) {
	key, err := GenerateEcdsa()
	if err != nil {
		t.Fatal("Fail to generate ecdsa")
	}
	privKey := key.GetPrivateKey()

	msg, err := hex.DecodeString("aabbccdd0011223344aabbccdd0011223344aabbccdd0011223344aabbccdd00")
	if err != nil {
		t.Fatal("Fail to decode msg")
	}

	signature := crypto.SignEcdsa(msg, privKey)
	if signature == nil {
		t.Fatal("Signature is empty")
	}
}
