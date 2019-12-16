package primitives

import (
	"errors"
)

type AssetAddress struct {
	AssetType byte
	Payload   H160 // should consider MultisigValue
	Value     string
}

func AssetAddressFromTypeAndPayload(assetType byte, payload H160, networkID string) (a AssetAddress, err error) {
	if assetType < 0 || assetType > 3 {
		err = errors.New("AssetType should be 1 or 2 or 3")
		return
	}

	words := toWords(append([]byte{1, assetType}, payload[:]...))
	address := bech32Encode(networkID+"a", words)

	return AssetAddress{AssetType: assetType, Payload: payload, Value: address}, nil
}

func AssetAddressFromString(address string) (a AssetAddress, err error) {
	if address[2] != 97 {
		err = errors.New("Not an AssetAddress")
		return
	}

	words := bech32Decode(address, address[0:3])
	bytes := fromWords(words)
	version := bytes[0]

	if version != 1 {
		err = errors.New("Version of AssetAddress should be 1")
		return
	}

	assetType := bytes[1]

	if assetType < 0 || assetType > 3 {
		err = errors.New("AssetType should be 1 or 2 or 3")
		return
	}
	// Should consider type <0x03 or type==3
	payload, _ := NewH160FromSlice(bytes[2:])
	return AssetAddress{AssetType: assetType, Payload: payload, Value: address}, nil
}
