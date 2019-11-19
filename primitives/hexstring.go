package primitives

import (
	"encoding/hex"
	"errors"

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

func StringToH128(s string) (H128, error) {

	var res H128

	if len(s) != 32 && (len(s) != 34 || s[0:2] != "0x") {
		return res, errors.New("Not a 128bit data")
	}

	if len(s) == 34 {
		s = s[2:]
	}

	b, err := hex.DecodeString(s)
	if err != nil {
		return res, err
	}

	copy(res[:], b[:16])

	return res, err
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

func NewH160FromSlice(raw []byte) (h H160, err error) {
	if len(raw) != 20 {
		err = errors.New("Unexpected type casting to H160")
		return
	}
	copy(h[:], raw)
	return

}
func StringToH160(s string) (H160, error) {

	var res H160

	if len(s) != 40 && (len(s) != 42 || s[0:2] != "0x") {
		return res, errors.New("Not a 160bit data")
	}

	if len(s) == 42 {
		s = s[2:]
	}

	b, err := hex.DecodeString(s)
	if err != nil {
		return res, err
	}

	copy(res[:], b[:20])

	return res, err
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

func StringToH256(s string) (H256, error) {

	var res H256

	if len(s) != 64 && (len(s) != 66 || s[0:2] != "0x") {
		return res, errors.New("Not a 256bit data")
	}

	if len(s) == 66 {
		s = s[2:]
	}

	b, err := hex.DecodeString(s)
	if err != nil {
		return res, err
	}

	copy(res[:], b[:32])

	return res, err
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

func StringToH512(s string) (H512, error) {

	var res H512

	if len(s) != 128 && (len(s) != 130 || s[0:2] != "0x") {
		return res, errors.New("Not a 512bit data")
	}

	if len(s) == 130 {
		s = s[2:]
	}

	b, err := hex.DecodeString(s)
	if err != nil {
		return res, err
	}

	copy(res[:], b[:64])

	return res, err
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
