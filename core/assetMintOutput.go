package core

import (
	"encoding/hex"
	"errors"

	"github.com/CodeChain-io/codechain-sdk-go/key"
	"github.com/CodeChain-io/codechain-sdk-go/primitives"
)

type AssetMintOutputJSON struct {
	LockscriptHash string   `json:"lockScriptHash"`
	Parameters     []string `json:"parameters"`
	Supply         string   `json:"supply"`
}

type AssetMintOutput struct {
	LockScriptHash primitives.H160
	Parameters     [][]byte
	Supply         primitives.U64
}

func AssetMinttOutputFromJSON(j AssetMintOutputJSON) AssetMintOutput {

	parameters := make([][]byte, len(j.Parameters))
	for i, d := range j.Parameters {
		res, _ := hex.DecodeString(d)
		parameters[i] = res
	}

	hash, _ := primitives.StringToH160(j.LockscriptHash)
	supply := primitives.NewU64(j.Supply)

	return NewAssetMintOutput(hash, parameters, supply)
}

func NewAssetMintOutput(lockScriptHash primitives.H160, parameters [][]byte, supply primitives.U64) AssetMintOutput {
	return AssetMintOutput{lockScriptHash, parameters, supply}
}

func NewAssetMintOutputWithRecipient(recipient primitives.AssetAddress, supply primitives.U64) (output AssetMintOutput, err error) {
	assetType := recipient.AssetType
	if assetType == 0 {
		return AssetMintOutput{recipient.Payload, [][]byte{}, supply}, nil
	} else if assetType == 1 {
		lockScriptHash := key.GetP2PKHLockScriptHash()
		parameters := [][]byte{recipient.Payload.Bytes()}
		return AssetMintOutput{lockScriptHash, parameters, supply}, nil
	} else if assetType == 2 {
		lockScriptHash := key.GetP2PKHBurnLockScriptHash()
		parameters := [][]byte{recipient.Payload.Bytes()}
		return AssetMintOutput{lockScriptHash, parameters, supply}, nil
	} else {
		err = errors.New("Unexpected type of AssetAddress")
		return
	}
}

func (output AssetMintOutput) ToJSON() AssetMintOutputJSON {

	parameterInByte := output.Parameters
	var paramterInString = make([]string, len(parameterInByte))
	for i, d := range parameterInByte {
		paramterInString[i] = hex.EncodeToString(d)
	}

	return AssetMintOutputJSON{
		output.LockScriptHash.ToJSON(),
		paramterInString,
		output.Supply.ToJSON()}
}
