package key

import (
	"crypto/ecdsa"
	"crypto/rand"

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

func CreateAssetAddress(networkID string) (a primitives.AssetAddress, err error) {
	key, err := GenerateEcdsa()
	if err != nil {
		return
	}
	hash, err := key.CreateKey()
	if err != nil {
		return
	}
	return primitives.AssetAddressFromTypeAndPayload(byte(1), hash, networkID)
}

func CreatePlatformAddress(networkID string) (a primitives.PlatformAddress, secret []byte, err error) {
	key, err := GenerateEcdsa()
	if err != nil {
		return
	}
	hash, err := key.CreateKey()
	if err != nil {
		return
	}

	add, err := primitives.PlatformAddressFromAccountID(hash, networkID)
	return add, key.GetPrivateKey(), err
}
