package transaction

import (
	"encoding/hex"

	"github.com/CodeChain-io/codechain-sdk-go/core"
	"github.com/CodeChain-io/codechain-sdk-go/crypto"
	"github.com/CodeChain-io/codechain-sdk-go/primitives"
	"github.com/ethereum/go-ethereum/rlp"
)

type AssetMintTransactionJSON struct {
	NetworkID           string                   `json:"networkId"`
	ShardID             uint                     `json:"shardId"`
	Metadata            string                   `json:"metadata"`
	Output              core.AssetMintOutputJSON `json:"output"`
	Approver            string                   `json:"approver,omitempty"`
	Registrar           string                   `json:"regitrar,omitempty"`
	AllowedScriptHashes []string                 `json:"allowedScriptHashes"`
}

type MintAssetActionJSON struct {
	Approvals []string `json:"networkId"`
}

type MintAsset struct {
	core.Transaction
	// NetworkID           string
	ShardID             uint
	Metadata            string
	Output              core.AssetMintOutput
	Approver            *primitives.PlatformAddress
	Registrar           *primitives.PlatformAddress
	AllowedScriptHashes []primitives.H160
	Approvals           []string
}

func NewMintAsset(networkID string,
	shardID uint, metadata string,
	output core.AssetMintOutput,
	approver *primitives.PlatformAddress,
	registrar *primitives.PlatformAddress,
	allowedScriptHashes []primitives.H160,
	approvals []string) MintAsset {

	// tx := NewAssetMintTransaction(networkID, shardID, metadata, output, approver, registrar, allowedScriptHashes)
	transaction := core.NewTransaction(networkID)

	return MintAsset{transaction, shardID, metadata, output, approver, registrar, allowedScriptHashes, []string{}}

}

func (t MintAsset) Tracker() primitives.H256 {
	hash, _ := crypto.Blake256(t.ActionRlpBytes())
	return primitives.NewH256(hash)
}

func (t *MintAsset) AddApproval(approval string) {
	t.Approvals = append(t.Approvals, approval)
}

func (t MintAsset) GetOutput() core.AssetMintOutput {
	return t.Output
}

func (t MintAsset) GetMintedAsset() core.Asset {

	parameters := make([]string, len(t.Output.Parameters))
	for i, d := range t.Output.Parameters {
		parameters[i] = hex.EncodeToString(d)
	}
	return core.NewAsset(t.GetAssetType(),
		t.ShardID,
		t.Output.LockScriptHash,
		parameters,
		t.Output.Supply,
		t.Tracker(),
		uint(0))
}

func (t MintAsset) GetAssetScheme() core.AssetScheme {
	return core.NewAssetScheme(t.Transaction.NetworkID(), t.ShardID, t.Metadata, t.Output.Supply, *t.Approver, *t.Registrar, t.AllowedScriptHashes, []core.AssetPool{}, uint(0))
}

func (t MintAsset) GetAssetType() primitives.H160 {
	blake, _ := crypto.Blake160(t.Tracker().Bytes())
	return primitives.NewH160(blake)
}

func (t MintAsset) GetType() string {
	return "mintAsset"
}

func (t MintAsset) ActionToEncodeObject() []interface{} {
	allowedScriptHashes := make([][]byte, len(t.AllowedScriptHashes))
	for i, d := range t.AllowedScriptHashes {
		res := d.ToEncodeObject()
		allowedScriptHashes[i] = res
	}

	var approver []interface{}
	if t.Approver != nil {
		approver = []interface{}{t.Approver.AccountID.ToEncodeObject()}
	}

	var registrar []interface{}
	if t.Approver != nil {
		registrar = []interface{}{t.Registrar.AccountID.ToEncodeObject()}
	}

	return []interface{}{
		uint(0x13),
		t.NetworkID(),
		t.ShardID,
		t.Metadata,
		t.Output.LockScriptHash.ToEncodeObject(),
		t.Output.Parameters,
		t.Output.Supply.ToEncodeObject(),
		approver,
		registrar,
		allowedScriptHashes}
}

func (t MintAsset) ActionToJSON() interface{} {
	return MintAssetActionJSON{t.Approvals}
}

func (t *MintAsset) Sign(secret primitives.H256, seq uint, fee primitives.U64) core.SignedTransaction {
	t.SetSeq(seq)
	t.SetFee(fee)
	sig := crypto.SignEcdsa(t.UnsignedHash().Bytes(), secret.Bytes())

	return core.NewSignedTransaction(t, sig, nil, nil, nil)
}

func (t MintAsset) ToEncodeObject() []interface{} {
	return []interface{}{byte(t.Seq()), t.Fee().ToEncodeObject(), t.NetworkID(), append(t.ActionToEncodeObject(), t.Approvals)}
}

func (t MintAsset) UnsignedHash() primitives.H256 {
	blake, _ := crypto.Blake256(t.RlpBytes())
	return primitives.NewH256(blake)
}

func (t MintAsset) RlpBytes() []byte {
	x, _ := rlp.EncodeToBytes(t.ToEncodeObject())
	return x
}

func (t MintAsset) ActionRlpBytes() []byte {
	x, _ := rlp.EncodeToBytes(t.ActionToEncodeObject())
	return x
}
