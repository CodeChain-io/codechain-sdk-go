package core

import (
	"encoding/hex"

	"github.com/CodeChain-io/codechain-sdk-go/primitives"
)

type AssetJSON struct {
	AssetType              string   `json:"assetType"`
	LockScriptHash         string   `json:"lockScriptHash"`
	Parameters             []string `json:"parameters"`
	Quantity               string   `json:"quantity"`
	ShardID                uint     `json:"shardId"`
	Tracker                string   `json:"tracker"`
	TransactionOutputIndex uint     `json:"transactionOutputIndex"`
}

type Asset struct {
	AssetType      primitives.H160
	ShardID        uint
	LockScriptHash primitives.H160
	Parameters     [][]byte
	Quantity       primitives.U64
	OutPoint       AssetOutPoint
}

func NewAsset(assetType primitives.H160,
	shardID uint,
	lockScriptHash primitives.H160,
	parameters []string,
	quantity primitives.U64,
	tracker primitives.H256,
	transactionOutputIndex uint) Asset {

	p := make([][]byte, len(parameters))
	for i, d := range parameters {
		res, _ := hex.DecodeString(d)
		p[i] = res
	}

	out := AssetOutPoint{tracker, transactionOutputIndex, assetType, shardID, quantity, lockScriptHash, p}

	return Asset{
		assetType,
		shardID,
		lockScriptHash,
		p,
		quantity,
		out}
}

func (a Asset) ToJSON() AssetJSON {
	parameters := make([]string, len(a.Parameters))
	for i, d := range a.Parameters {
		parameters[i] = hex.EncodeToString(d)
	}

	return AssetJSON{
		a.AssetType.ToJSON(),
		a.LockScriptHash.ToJSON(),
		parameters,
		a.Quantity.ToJSON(),
		a.ShardID,
		a.OutPoint.Tracker.ToJSON(),
		a.OutPoint.Index}
}

func (a Asset) CreateTransferInput() AssetTransferInput {
	return NewAssetTransferInput(a.OutPoint, nil, nil, nil)
}

func (a Asset) CreateTransferInputWithTimelock(timelock Timelock) AssetTransferInput {
	return NewAssetTransferInput(a.OutPoint, &timelock, nil, nil)
}
