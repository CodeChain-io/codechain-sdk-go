package core

import (
	"encoding/hex"

	"github.com/CodeChain-io/codechain-sdk-go/primitives"
)

type AssetOutPointJSON struct {
	Tracker        string   `json:"tracker"`
	Index          uint     `json:"index"`
	AssetType      string   `json:"assetType"`
	ShardID        uint     `json:"shardId"`
	Quantity       string   `json:"quantity"`
	LockScriptHash string   `json:"lockScriptHash,omitempty"`
	Parameters     []string `json:"parameters,omitempty"`
}

type AssetOutPoint struct {
	Tracker        primitives.H256
	Index          uint
	AssetType      primitives.H160
	ShardID        uint
	Quantity       primitives.U64
	LockScriptHash primitives.H160
	Parameters     [][]byte
}

func AssetOutPointFromJSON(j AssetOutPointJSON) AssetOutPoint {

	lockScriptHash, _ := primitives.StringToH160(j.LockScriptHash)

	parameters := make([][]byte, len(j.Parameters))
	for i, d := range j.Parameters {
		res, _ := hex.DecodeString(d)
		parameters[i] = res
	}

	tracker, _ := primitives.StringToH256(j.Tracker)
	assetType, _ := primitives.StringToH160(j.AssetType)

	return AssetOutPoint{
		tracker,
		j.Index,
		assetType,
		j.ShardID,
		primitives.NewU64(j.Quantity),
		lockScriptHash,
		parameters}
}

func (p AssetOutPoint) ToEncodeObject() []interface{} {
	return []interface{}{
		p.Tracker.ToEncodeObject(),
		p.Index,
		p.AssetType.ToEncodeObject(),
		p.ShardID,
		p.Quantity.ToEncodeObject()}
}

func (p AssetOutPoint) ToJSON() AssetOutPointJSON {
	parameters := make([]string, len(p.Parameters))
	for i, d := range p.Parameters {
		parameters[i] = hex.EncodeToString(d)
	}

	return AssetOutPointJSON{p.Tracker.ToJSON(),
		p.Index,
		p.AssetType.ToJSON(),
		p.ShardID,
		p.Quantity.ToJSON(),
		p.LockScriptHash.ToJSON(),
		parameters}
}
