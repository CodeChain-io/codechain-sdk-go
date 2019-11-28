package transaction

import (
	"github.com/CodeChain-io/codechain-rpc-go/primitives"
	"github.com/ethereum/go-ethereum/rlp"
)

type PayAction struct {
	Receiver string `json:"receiver"`
	Quantity string `json:"quantity"`
}

type Pay struct {
	transaction
	Receiver primitives.PlatformAddress
	Quantity primitives.U64
}

func NewPay(receiver primitives.PlatformAddress,
	quantity primitives.U64,
	networkID string) Pay {

	t := newTransaction(networkID)
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

func (p Pay) ActionToJSON() PayAction {
	return PayAction{
		p.Receiver.Value,
		p.Quantity.ToJSON()}
}

func (p Pay) RlpBytes() []byte {
	x, _ := rlp.EncodeToBytes(p.ToEncodeObject())
	return x
}
