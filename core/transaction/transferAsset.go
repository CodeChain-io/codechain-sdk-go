package transaction

import (
	"encoding/hex"
	"errors"

	"github.com/CodeChain-io/codechain-sdk-go/core"
	"github.com/CodeChain-io/codechain-sdk-go/crypto"
	"github.com/CodeChain-io/codechain-sdk-go/key"
	"github.com/CodeChain-io/codechain-sdk-go/primitives"
	"github.com/ethereum/go-ethereum/rlp"
)

type AssetTransferTransactionJSON struct {
	NetworkID string                         `json:"networkId"`
	Burns     []core.AssetTransferInputJSON  `json:"burns"`
	Inputs    []core.AssetTransferInputJSON  `json:"inputs"`
	Outputs   []core.AssetTransferOutputJSON `json:"outputs"`
}

type TransferAssetActionJSON struct {
	AssetTransferTransactionJSON
	Metadata   string   `json:"metadata"`
	Approvals  []string `json:"aprovals"`
	Expiration int      `json:"expiration:omitempty"`
}

type TransferAsset struct {
	core.Transaction
	Burns      []core.AssetTransferInput
	Inputs     []core.AssetTransferInput
	Outputs    []core.AssetTransferOutput
	Approvals  []string
	Metadata   string
	Expiration *int
}

func NewTransferAsset(
	burns []core.AssetTransferInput,
	inputs []core.AssetTransferInput,
	outputs []core.AssetTransferOutput,
	networkID string,
	metadata string,
	approvals []string,
	expiration *int) TransferAsset {

	tx := core.NewTransaction(networkID)
	return TransferAsset{tx, burns, inputs, outputs, approvals, metadata, expiration}
}

func (t TransferAsset) Tracker() primitives.H256 {
	hash, _ := crypto.Blake256(t.ActionRlpBytes())
	return primitives.NewH256(hash)
}

func (t *TransferAsset) AddApproval(approval string) {
	t.Approvals = append(t.Approvals, approval)
}

func (t *TransferAsset) AddBurn(burn core.AssetTransferInput) {
	t.Burns = append(t.Burns, burn)
}

func (t *TransferAsset) AddBurns(burns []core.AssetTransferInput) {
	t.Burns = append(t.Burns, burns...)
}

func (t *TransferAsset) AddBurnWithAsset(asset core.Asset) {
	t.Burns = append(t.Burns, asset.CreateTransferInput())
}

func (t *TransferAsset) AddBurnsWithAsset(assets []core.Asset) {

	burns := make([]core.AssetTransferInput, len(assets))
	for i, d := range assets {
		burns[i] = d.CreateTransferInput()
	}
	t.Burns = append(t.Burns, burns...)
}

func (t *TransferAsset) AddInput(input core.AssetTransferInput) {
	t.Inputs = append(t.Inputs, input)
}

func (t *TransferAsset) AddInputs(inputs []core.AssetTransferInput) {
	t.Inputs = append(t.Inputs, inputs...)
}

func (t *TransferAsset) AddInputWithAsset(asset core.Asset) {
	t.Inputs = append(t.Inputs, asset.CreateTransferInput())
}

func (t *TransferAsset) AddInputsWithAsset(assets []core.Asset) {

	inputs := make([]core.AssetTransferInput, len(assets))
	for i, d := range assets {
		inputs[i] = d.CreateTransferInput()
	}
	t.Inputs = append(t.Inputs, inputs...)
}

func (t *TransferAsset) AddOutputs(outputs []core.AssetTransferOutput) TransferAsset {
	t.Outputs = append(t.Outputs, outputs...)
	return *t
}

func (t TransferAsset) ToEncodeObject() []interface{} {
	actionEncode := t.ActionToEncodeObject()

	actionEncode = append(actionEncode, t.Metadata)
	actionEncode = append(actionEncode, t.Approvals)
	if t.Expiration != nil {
		actionEncode = append(actionEncode, []interface{}{t.Expiration})
	} else {
		actionEncode = append(actionEncode, []interface{}{})
	}

	return []interface{}{t.Seq(), t.Fee().ToEncodeObject(), t.NetworkID(), actionEncode}

}

func (t TransferAsset) ActionToEncodeObject() []interface{} {
	burns := make([][]interface{}, len(t.Burns))
	for i, d := range t.Burns {
		res := d.ToEncodeObject()
		burns[i] = res
	}

	inputs := make([][]interface{}, len(t.Inputs))
	for i, d := range t.Inputs {
		res := d.ToEncodeObject()
		inputs[i] = res
	}

	outputs := make([][]interface{}, len(t.Outputs))
	for i, d := range t.Outputs {
		res := d.ToEncodeObject()
		outputs[i] = res
	}

	return []interface{}{
		uint(0x14),
		t.Transaction.NetworkID(),
		burns,
		inputs,
		outputs,
		[]interface{}{}}
}

func (t TransferAsset) GetTransferredAsset(index int) core.Asset {
	output := t.Outputs[index]

	parameters := make([]string, len(output.Parameters))
	for i, d := range output.Parameters {
		res := hex.EncodeToString(d)
		parameters[i] = res
	}

	return core.NewAsset(output.AssetType,
		output.ShardID,
		output.LockScriptHash,
		parameters,
		output.Quantity,
		t.Tracker(),
		uint(index))
}

func (t TransferAsset) GetTransferredAssets() []core.Asset {
	assets := make([]core.Asset, len(t.Outputs))
	for i := range t.Outputs {
		assets[i] = t.GetTransferredAsset(i)
	}
	return assets
}

func (t TransferAsset) HashWithoutScript() primitives.H256 {
	inputs := make([]core.AssetTransferInput, len(t.Inputs))
	for i, d := range t.Inputs {
		inputs[i] = d.WithOutScript()
	}

	burns := make([]core.AssetTransferInput, len(t.Burns))
	for i, d := range t.Burns {
		burns[i] = d.WithOutScript()
	}

	outputs := t.Outputs

	t = NewTransferAsset(burns, inputs, outputs, t.NetworkID(), "", []string{}, nil)
	hashKey, _ := crypto.Blake128([]byte{3})

	blake, _ := crypto.Blake256WithKey(t.ActionRlpBytes(), hashKey)

	return primitives.NewH256(blake)
}

func (t TransferAsset) ActionRlpBytes() []byte {
	res, _ := rlp.EncodeToBytes(t.ActionToEncodeObject())
	return res
}

func (t TransferAsset) ActionToJSON() interface{} {
	burns := make([]core.AssetTransferInputJSON, len(t.Burns))
	for i, d := range t.Burns {
		res := d.ToJSON()
		burns[i] = res
	}

	inputs := make([]core.AssetTransferInputJSON, len(t.Inputs))
	for i, d := range t.Inputs {
		res := d.ToJSON()
		inputs[i] = res
	}

	outputs := make([]core.AssetTransferOutputJSON, len(t.Outputs))
	for i, d := range t.Outputs {
		res := d.ToJSON()
		outputs[i] = res
	}

	return AssetTransferTransactionJSON{t.NetworkID(), burns, inputs, outputs}
}

func (t TransferAsset) UnsignedHash() primitives.H256 {
	hash, _ := crypto.Blake256(t.RlpBytes())

	var value [32]byte
	copy(value[:], hash[:32])
	return primitives.H256(value)
}

func (t TransferAsset) RlpBytes() []byte {
	x, _ := rlp.EncodeToBytes(t.ToEncodeObject())
	return x
}

func (t *TransferAsset) Sign(secret primitives.H256, seq uint, fee primitives.U64) core.SignedTransaction {
	t.SetSeq(seq)
	t.SetFee(fee)
	sig := crypto.SignEcdsa(t.UnsignedHash().Bytes(), secret.Bytes())

	return core.NewSignedTransaction(t, sig, nil, nil, nil)
}

func (t TransferAsset) GetType() string {
	return "transferAsset"
}

func (t *TransferAsset) SignTransactionInput(index int, pubkey []byte, privkey []byte) error {

	parameters := t.Inputs[index].PrevOut.Parameters
	parameter := hex.EncodeToString(parameters[0])
	pubkeyhsh, _ := crypto.Blake160(pubkey)
	if parameter != hex.EncodeToString(pubkeyhsh) {
		return errors.New("Wrong Public Key")
	}

	scriptOpcode := key.GetP2PKHLockScript()
	opcode := make([]byte, len(scriptOpcode))
	for i, d := range scriptOpcode {
		opcode[i] = byte(d)
	}

	t.Inputs[index].SetLockScript(opcode)

	message := t.HashWithoutScript()

	script := key.CreateP2PKHUnlockScript(pubkey, privkey, message)

	bytescript := make([]byte, len(script))
	for i, d := range script {
		bytescript[i] = byte(d)
	}
	t.Inputs[index].SetUnlockScript(bytescript)

	return nil
}
