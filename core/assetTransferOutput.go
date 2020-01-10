package core

import (
	"encoding/hex"
	"errors"

	"github.com/CodeChain-io/codechain-sdk-go/key"
	"github.com/CodeChain-io/codechain-sdk-go/primitives"
)

type AssetTransferOutputJSON struct {
	LockScriptHash string   `json:"lockScriptHash"`
	Parameters     []string `json:"parameters"`
	AssetType      string   `json:"assetType"`
	ShardID        uint     `json:"shardId"`
	Quantity       string   `json:"quantity"`
}

type AssetTransferOutput struct {
	LockScriptHash primitives.H160
	Parameters     [][]byte
	AssetType      primitives.H160
	ShardID        uint
	Quantity       primitives.U64
}

func NewAssetTransferOutput(lockScriptHash primitives.H160,
	parameters [][]byte,
	assetType primitives.H160,
	shardID uint,
	quantity primitives.U64) AssetTransferOutput {
	return AssetTransferOutput{lockScriptHash, parameters, assetType, shardID, quantity}
}

func NewAssetTransferOutputWithAddress(recipient primitives.AssetAddress,
	assetType primitives.H160,
	shardID uint,
	quantity primitives.U64) (a AssetTransferOutput, err error) {
	assetAddressType := recipient.AssetType
	payload := recipient.Payload
	var lockScriptHash primitives.H160
	var parameters [][]byte

	switch assetAddressType {
	case 0x00:
		lockScriptHash = payload
		parameters = [][]byte{{}}
	case 0x01:
		lockScriptHash = key.GetP2PKHLockScriptHash()
		parameters = [][]byte{payload.Bytes()}
	case 0x02:
		lockScriptHash = key.GetP2PKHBurnLockScriptHash()
		parameters = [][]byte{payload.Bytes()}
	default:
		err = errors.New("Unexpected type of AssetAddress: " + string(assetAddressType))
		return
	}

	return AssetTransferOutput{lockScriptHash, parameters, assetType, shardID, quantity}, nil

}

func (a AssetTransferOutput) ToEncodeObject() []interface{} {

	return []interface{}{
		a.LockScriptHash.ToEncodeObject(),
		a.Parameters,
		a.AssetType.ToEncodeObject(),
		a.ShardID,
		a.Quantity.ToEncodeObject()}
}

func (a AssetTransferOutput) ToJSON() AssetTransferOutputJSON {
	parameters := make([]string, len(a.Parameters))
	for i, d := range a.Parameters {
		parameters[i] = hex.EncodeToString(d)
	}

	return AssetTransferOutputJSON{a.LockScriptHash.ToJSON(), parameters, a.AssetType.ToJSON(), a.ShardID, a.Quantity.ToJSON()}
}
