package core

import (
	"encoding/hex"

	"github.com/CodeChain-io/codechain-sdk-go/crypto"
	"github.com/CodeChain-io/codechain-sdk-go/primitives"
	"github.com/ethereum/go-ethereum/rlp"
)

type SignedTransactionJSON struct {
	BlockNumber      int    `json:"blockNumber,omitempty"`
	BlockHash        string `json:"blockHash,omitempty"`
	TransactionIndex int    `json:"transactionIndex,omitempty"`
	Sig              string `json:"sig"`
	Hash             string `json:"hash"`
}

type SignedTransaction struct {
	Unsigned         TransactionInterface
	BlockNumber      *int
	BlockHash        *primitives.H256
	TransactionIndex *int
	signature        []byte
}

func NewSignedTransaction(
	unsigned TransactionInterface,
	signature []byte,
	blockNumber *int,
	blockHash *primitives.H256,
	transactionIndex *int) SignedTransaction {
	return SignedTransaction{unsigned, blockNumber, blockHash, transactionIndex, signature}
}

func (t SignedTransaction) Signature() []byte {
	return t.signature
}

func (t SignedTransaction) ToEncodeObject() []interface{} {
	// FIXME return error when seq, fee are not set.
	return append(t.Unsigned.ToEncodeObject(), t.Signature())
}

func (t SignedTransaction) RlpBytes() []byte {
	x, _ := rlp.EncodeToBytes(t.ToEncodeObject())
	return x
}

func (t SignedTransaction) Hash() primitives.H256 {
	hash, _ := crypto.Blake256([]byte{0}) //t.RlpBytes())

	var value [32]byte
	copy(value[:], hash[:32])
	return primitives.H256(value) // byte
}

func (t SignedTransaction) GetSignerAccountID() primitives.H160 {
	pubKey := crypto.RecoverEcdsa(t.Unsigned.UnsignedHash().Bytes(), []byte(t.signature))
	id, _ := crypto.Blake160(pubKey)
	return primitives.NewH160(id)
}

func (t SignedTransaction) GetSignerAddress(networkID string) primitives.PlatformAddress {
	a, _ := primitives.PlatformAddressFromAccountID(t.GetSignerAccountID(), networkID)
	return a
}

func (t SignedTransaction) GetSignerPublic() primitives.H512 {
	a := crypto.RecoverEcdsa(t.Unsigned.UnsignedHash().Bytes(), []byte(t.signature))
	return primitives.NewH512(a)
}

func (t SignedTransaction) ToJSON() SignedTransactionJSON {
	return SignedTransactionJSON{
		*t.BlockNumber,
		t.BlockHash.ToString(),
		*t.TransactionIndex,
		hex.EncodeToString(t.signature),
		t.Hash().ToJSON()}
}
