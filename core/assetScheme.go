package core

import (
	"github.com/CodeChain-io/codechain-sdk-go/primitives"
)

type AssetPool struct {
	AssetType primitives.H160
	Quantity  primitives.U64
}

type AssetPoolJSON struct {
	AssetType string `json:"assetType"`
	Quantity  string `json:"quantity"`
}

type AssetSchemeJSON struct {
	Metadata            string          `json:"metaData"`
	Supply              string          `json:"supply"`
	Approver            string          `json:"approver"`
	Registrar           string          `json:"registrar"`
	AllowedScriptHashes []string        `json:"allowedScriptHashes"`
	Pool                []AssetPoolJSON `json:"pool"`
	Seq                 uint            `json:"seq"`
}

type AssetScheme struct {
	NetworkID           string
	ShardID             uint
	Metadata            string
	Supply              primitives.U64
	Approver            primitives.PlatformAddress
	Registrar           primitives.PlatformAddress
	AllowedScriptHashes []primitives.H160
	Pool                []AssetPool
	Seq                 uint
}

func NewAssetScheme(networkID string, shardID uint, metadata string, supply primitives.U64, approver primitives.PlatformAddress, registrar primitives.PlatformAddress, allowedScriptHashes []primitives.H160, pool []AssetPool, seq uint) AssetScheme {
	return AssetScheme{networkID, shardID, metadata, supply, approver, registrar, allowedScriptHashes, pool, seq}
}

func (s AssetScheme) ToJSON() AssetSchemeJSON {

	hashes := make([]string, len(s.AllowedScriptHashes))
	for i, d := range s.AllowedScriptHashes {
		res := d.ToJSON()
		hashes[i] = res
	}

	pool := make([]AssetPoolJSON, len(s.Pool))
	for i, d := range s.Pool {
		res := AssetPoolJSON{d.AssetType.ToJSON(), d.Quantity.ToJSON()}
		pool[i] = res
	}

	return AssetSchemeJSON{s.Metadata, s.Supply.ToJSON(), s.Approver.Value, s.Registrar.Value, hashes, nil, s.Seq}
}

// Removed because of low usage and causing circular import with core/transaction package.
/*func (s AssetScheme) CreateMintAssetTransaction(recipient primitives.AssetAddress) MintAsset {
	assetMintOutput, _ := NewAssetMintOutputWithRecipient(recipient, s.Supply)
	return NewMintAsset(s.NetworkID, s.ShardID, s.Metadata, assetMintOutput, &s.Approver, &s.Registrar, s.AllowedScriptHashes, []string{})
}*/
