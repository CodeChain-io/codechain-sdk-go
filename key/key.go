package key

import (
	"crypto/ecdsa"
	"crypto/rand"

	"github.com/CodeChain-io/codechain-sdk-go/crypto"
)

func GenerateKey() ([]byte, error) {
	privKey, err := ecdsa.GenerateKey(crypto.S256(), rand.Reader)
	if err != nil {
		return nil, err
	}

	rawPrivKey := make([]byte, 32)
	blob := privKey.D.Bytes()
	copy(rawPrivKey[32-len(blob):], blob)

	return rawPrivKey, nil
}
