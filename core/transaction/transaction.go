package transaction

import (
	"encoding/hex"

	"github.com/CodeChain-io/codechain-rpc-go/crypto"
	"github.com/CodeChain-io/codechain-rpc-go/primitives"
)

type ActionJSON struct {
	Action string `json:"action"`
	Type   string `json:"type"`
}

type TransactionJSON struct {
	Action    string `json:"action"` // Assign action
	NetworkID string `json:"networkId"`
	Seq       uint   `json:"seq,omitempty"`
	Fee       string `json:"string,omitempty"`
}

type TransactionInterface interface {
	Seq() int
	Fee() primitives.U64
	SetSeq(int)
	SetFee(primitives.U64)
	NetworkID() string
	ToEncodeObject() string
	RlpBytes() []byte
	UnsignedHash() primitives.H256
	Sign(primitives.H256, int, primitives.U64) interface{} // FIXME
	ToJSON() string
	GetType() string
	ActionToJSON() interface{}
	ActionToEncodeObject() []interface{}
}

type transaction struct {
	seq       uint
	fee       primitives.U64
	networkID string
}

func newTransaction(networkID string) transaction {
	return transaction{0, primitives.NewU64("0"), networkID}
}

func (t transaction) Seq() uint {
	return t.seq
}

func (t transaction) Fee() primitives.U64 {
	return t.fee
}

func (t *transaction) SetSeq(seq uint) {
	t.seq = seq
}

func (t *transaction) SetFee(fee primitives.U64) {
	t.fee = fee
}

func (t transaction) NetworkID() string {
	return t.networkID
}

func (t transaction) UnsignedHash() primitives.H256 {
	hash, _ := crypto.Blake256([]byte{0}) //t.RlpBytes())

	var value [32]byte
	copy(value[:], hash[:32])
	return primitives.H256(value) // byte
}

func (t transaction) Sign(secret primitives.H256, seq *int, fee *primitives.U64) SignedTransaction {
	// Handle error
	sig := crypto.SignEcdsa([]byte(t.UnsignedHash().ToString()), []byte(secret.ToString()))

	return NewSignedTransaction(t, hex.EncodeToString(sig), nil, nil, nil) // nil
}

func (t transaction) ToJSON() TransactionJSON {
	return TransactionJSON{
		t.GetType(), // assign action type.
		t.networkID,
		t.seq,
		t.fee.ToJSON()}
}

func (t transaction) GetType() string {
	return "Transaction"
}
