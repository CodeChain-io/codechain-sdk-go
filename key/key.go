package key

import (
	"crypto/ecdsa"
	"crypto/rand"
	"errors"
	"math/big"

	"github.com/CodeChain-io/codechain-sdk-go/crypto"
	"github.com/CodeChain-io/codechain-sdk-go/primitives"
)

type EcdsaKey ecdsa.PrivateKey

func GenerateEcdsa() (privateKey EcdsaKey, err error) {
	privKey, err := ecdsa.GenerateKey(crypto.S256(), rand.Reader)
	if err != nil {
		return
	}
	return EcdsaKey(*privKey), nil
}

func (t EcdsaKey) GetPrivateKey() []byte {
	rawPrivKey := make([]byte, 32)
	blob := t.D.Bytes()
	copy(rawPrivKey[32-len(blob):], blob)

	return rawPrivKey
}

func (t EcdsaKey) GetPublicKey() []byte {
	blobX := t.PublicKey.X.Bytes()
	blobY := t.PublicKey.Y.Bytes()

	rawPubKey := make([]byte, 64)
	copy(rawPubKey[32-len(blobX):], blobX)
	copy(rawPubKey[64-len(blobY):], blobY)

	return rawPubKey
}

func (t EcdsaKey) CreateKey() (h primitives.H160, err error) {
	hash, err := crypto.Blake160(t.GetPublicKey())
	if err != nil {
		return
	}
	return primitives.NewH160FromSlice(hash)
}

func CreateAssetAddress(key EcdsaKey, networkID string) (a primitives.AssetAddress, err error) {
	hash, err := key.CreateKey()
	if err != nil {
		return
	}
	return primitives.AssetAddressFromTypeAndPayload(byte(1), hash, networkID)
}

func CreatePlatformAddress(key EcdsaKey, networkID string) (a primitives.PlatformAddress, err error) {
	hash, err := key.CreateKey()
	if err != nil {
		return
	}

	return primitives.PlatformAddressFromAccountID(hash, networkID)
}

func ToECDSA(privateKey []byte) (*EcdsaKey, error) {
	priv := new(EcdsaKey)
	priv.PublicKey.Curve = crypto.S256()
	priv.D = new(big.Int).SetBytes(privateKey)

	priv.PublicKey.X, priv.PublicKey.Y = priv.PublicKey.Curve.ScalarBaseMult(privateKey)

	if priv.PublicKey.X == nil {
		return nil, errors.New("Invalid private key")
	}

	return priv, nil
}
