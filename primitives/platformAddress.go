package primitives

import (
	"encoding/hex"
	"errors"
	"golang.org/x/crypto/blake2b"
)

type PlatformAddress struct {
	AccountID H160
	Value     string
}

func PlatformAddressFromAccountID(accountID H160, networkID string) (p PlatformAddress, err error) {

	if len(networkID) != 2 {
		err = errors.New("The length networkID should be 2")
		return
	}

	words := toWords([]byte("01" + accountID.ToString()))
	address := PlatformAddress{AccountID: accountID, Value: bech32Encode(networkID+"c", words)}

	return address, nil
}

func PlatformAddressFromPublic(publicKey H512, networkID string) (PlatformAddress, error) {
	pubkey := H512.ToString(publicKey)
	accountID := GetAddcountIDFromPublic(pubkey)
	return PlatformAddressFromAccountID(accountID, networkID)
}

func GetAddcountIDFromPublic(publicKey string) H160 {
	return blake160(publicKey)
}

func blake160(data string) H160 {
	b, _ := blake2b.New(20, nil)
	s, _ := hex.DecodeString(data)
	b.Write(s)
	bs := b.Sum(nil)
	var result [20]byte
	copy(result[:], bs)
	return H160(result)
}
