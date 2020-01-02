package transaction

import (
	"github.com/CodeChain-io/codechain-sdk-go/core"
	"github.com/CodeChain-io/codechain-sdk-go/crypto"
	"github.com/CodeChain-io/codechain-sdk-go/primitives"
	"github.com/ethereum/go-ethereum/rlp"
)

type PayAction struct {
	Receiver string `json:"receiver"`
	Quantity string `json:"quantity"`
}

type Pay struct {
	core.Transaction
	Receiver primitives.PlatformAddress
	Quantity primitives.U64
}

func NewPay(receiver primitives.PlatformAddress,
	quantity primitives.U64,
	networkID string) Pay {

	t := core.NewTransaction(networkID)
	return Pay{t, receiver, quantity}
}

func (p Pay) GetType() string {
	return "pay"
}

func (p Pay) ActionToEncodeObject() []interface{} {
	return []interface{}{uint(2), p.Receiver.AccountID.ToEncodeObject(), p.Quantity.ToEncodeObject()}
}

func (p Pay) ToEncodeObject() []interface{} {
	return []interface{}{byte(p.Seq()), p.Fee().ToEncodeObject(), p.NetworkID(), p.ActionToEncodeObject()}
}

func (p Pay) ActionToJSON() interface{} {
	return PayAction{
		p.Receiver.Value,
		p.Quantity.ToJSON()}
}

func (p Pay) RlpBytes() []byte {
	x, _ := rlp.EncodeToBytes(p.ToEncodeObject())
	return x
}

func (p Pay) Hash() []byte {
	x, _ := crypto.Blake256(p.RlpBytes())
	return x
}

func (p Pay) UnsignedHash() primitives.H256 {
	hash, _ := crypto.Blake256(p.RlpBytes())

	var value [32]byte
	copy(value[:], hash[:32])
	return primitives.H256(value)
}

func (p *Pay) Sign(secret primitives.H256, seq uint, fee primitives.U64) core.SignedTransaction {
	// Handle error
	p.SetSeq(seq)
	p.SetFee(fee)
	sig := crypto.SignEcdsa(p.UnsignedHash().Bytes(), secret.Bytes())

	return core.NewSignedTransaction(p, sig, nil, nil, nil)
}
