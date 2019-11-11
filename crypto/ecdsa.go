package crypto

import (
	"crypto/elliptic"

	"github.com/ethereum/go-ethereum/crypto/secp256k1"
)

func SignEcdsa(msg, priv []byte) []byte {
	signature, _ := secp256k1.Sign(msg, priv)
	return signature
}

func VerifyEcdsa(msg, signature, pubkey []byte) bool {
	result := secp256k1.VerifySignature(pubkey, msg, signature)
	return result
}

func RecoverEcdsa(message, signature []byte) []byte {
	result, _ := secp256k1.RecoverPubkey(message, signature)
	return result
}

func S256() elliptic.Curve {
	return secp256k1.S256()
}
