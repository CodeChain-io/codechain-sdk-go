package core

import (
	"github.com/CodeChain-io/codechain-sdk-go/crypto"
	"github.com/CodeChain-io/codechain-sdk-go/primitives"
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
	Seq() uint
	Fee() primitives.U64
	SetSeq(uint)
	SetFee(primitives.U64)
	NetworkID() string
	ToEncodeObject() []interface{}
	RlpBytes() []byte
	UnsignedHash() primitives.H256
	Sign(primitives.H256, uint, primitives.U64) SignedTransaction
	ToJSON() TransactionJSON
	GetType() string
	ActionToJSON() interface{}
	ActionToEncodeObject() []interface{}
}

type Transaction struct {
	seq       uint
	fee       primitives.U64
	networkID string
}

func NewTransaction(networkID string) Transaction {
	return Transaction{0, primitives.NewU64("0"), networkID}
}

func (t Transaction) Seq() uint {
	return t.seq
}

func (t Transaction) Fee() primitives.U64 {
	return t.fee
}

func (t *Transaction) SetSeq(seq uint) {
	t.seq = seq
}

func (t *Transaction) SetFee(fee primitives.U64) {
	t.fee = fee
}

func (t Transaction) NetworkID() string {
	return t.networkID
}

func (t Transaction) UnsignedHash() primitives.H256 {
	hash, _ := crypto.Blake256([]byte{0}) //t.RlpBytes())

	var value [32]byte
	copy(value[:], hash[:32])
	return primitives.H256(value) // byte
}

func (t *Transaction) Sign(secret primitives.H256, seq uint, fee primitives.U64) SignedTransaction {
	// Handle error

	sig := crypto.SignEcdsa(t.UnsignedHash().Bytes(), secret.Bytes())
	t.SetSeq(seq)
	t.SetFee(fee)
	return NewSignedTransaction(t, sig, nil, nil, nil) // nil
}

func (t Transaction) ToJSON() TransactionJSON {
	return TransactionJSON{
		t.GetType(), // assign action type.
		t.networkID,
		t.seq,
		t.fee.ToJSON()}
}

func (t Transaction) GetType() string {
	return ""
}

func (t Transaction) ActionToEncodeObject() []interface{} {
	return nil
}

func (t Transaction) ActionToJSON() interface{} {
	return nil
}

func (t Transaction) RlpBytes() []byte {
	return nil
}

func (t Transaction) ToEncodeObject() []interface{} {
	return nil
}
