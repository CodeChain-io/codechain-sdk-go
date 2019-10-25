package primitives

import (
	"encoding/hex"

	"github.com/ethereum/go-ethereum/rlp"
)

// Handles 128-bit data
type H128 [16]byte

// Handles 160-bit data
type H160 [20]byte

// Handles 256-bit data
type H256 [32]byte

// Handles 512-bit data
type H512 [64]byte

func NewH128Zero() H128 {
	var zero H128 = [16]byte{}
	return zero
}

func (h H128) Cmp(g H128) bool {
	return h == g
}

func (h H128) ToString() string {
	return string(h[:])
}

func (h H128) RlpBytes() []byte {
	x, _ := rlp.EncodeToBytes(h)
	return x
}

func (h H128) ToJSON() string {
	var test = make([]byte, 32)
	innerArrOfH := [16]byte(h)
	hex.Encode(test, innerArrOfH[:])
	return "0x" + string(test)
}

func NewH160Zero() H160 {
	var zero H160 = [20]byte{}
	return zero
}

func (h H160) Cmp(g H160) bool {
	return h == g
}

func (h H160) ToString() string {
	return string(h[:])
}

func (h H160) RlpBytes() []byte {
	x, _ := rlp.EncodeToBytes(h)
	return x
}

func (h H160) ToJSON() string {
	var test = make([]byte, 40)
	innerArrOfH := [20]byte(h)
	hex.Encode(test, innerArrOfH[:])
	return "0x" + string(test)
}

func NewH256Zero() H256 {
	var zero H256 = [32]byte{}
	return zero
}

func (h H256) Cmp(g H256) bool {
	return h == g
}

func (h H256) ToString() string {
	return string(h[:])
}

func (h H256) RlpBytes() []byte {
	x, _ := rlp.EncodeToBytes(h)
	return x
}

func (h H256) ToJSON() string {
	var test = make([]byte, 64)
	innerArrOfH := [32]byte(h)
	hex.Encode(test, innerArrOfH[:])
	return "0x" + string(test)
}

func NewH512Zero() H512 {
	var zero H512 = [64]byte{}
	return zero
}

func (h H512) Cmp(g H512) bool {
	return h == g
}

func (h H512) ToString() string {
	return string(h[:])
}

func (h H512) RlpBytes() []byte {
	x, _ := rlp.EncodeToBytes(h)
	return x
}

func (h H512) ToJSON() string {
	var test = make([]byte, 128)
	innerArrOfH := [64]byte(h)
	hex.Encode(test, innerArrOfH[:])
	return "0x" + string(test)
}
