package primitives

import (
	"encoding/hex"
	"math/big"

	"github.com/ethereum/go-ethereum/rlp"
)

type U64 big.Int
type U128 big.Int
type U256 big.Int

func (u U64) String() string {
	var a = big.Int(u)
	return a.String()
}

func NewU64(v string) U64 {
	var x big.Int
	x.SetString(v, 10)
	return U64(x)
}

func (u *U64) set(y big.Int) {
	*u = U64(y)
}

func (u *U64) Add(x *U64, y *U64) *U64 {
	// Handle larger than max
	x1 := big.Int(*x)
	y1 := big.Int(*y)
	z1 := big.Int(*u)
	z1.Add(&x1, &y1)
	u.set(z1)
	return u
}

func (u *U64) Sub(x *U64, y *U64) *U64 {
	// Handle less than zero
	x1 := big.Int(*x)
	y1 := big.Int(*y)
	z1 := big.Int(*u)
	z1.Sub(&x1, &y1)
	u.set(z1)
	return u
}

func (u *U64) Mul(x *U64, y *U64) *U64 {
	x1 := big.Int(*x)
	y1 := big.Int(*y)
	z1 := big.Int(*u)
	z1.Sub(&x1, &y1)
	u.set(z1)
	return u
}

func (u *U64) Idiv(x *U64, y *U64) *U64 {
	x1 := big.Int(*x)
	y1 := big.Int(*y)
	z1 := big.Int(*u)
	z1.Quo(&x1, &y1)
	u.set(z1)
	return u
}

func (u *U64) Mod(x *U64, y *U64) *U64 {
	x1 := big.Int(*x)
	y1 := big.Int(*y)
	z1 := big.Int(*u)
	z1.Mod(&x1, &y1)
	u.set(z1)
	return u
}

func (u U64) RlpBytes() []byte {
	z1 := big.Int(u)
	x, _ := rlp.EncodeToBytes(z1.Bytes())
	return x
}

func (u *U64) cmp(v *U64) (r int) {
	u1 := big.Int(*u)
	v1 := big.Int(*v)
	r = u1.Cmp(&v1)
	return r
}

func (u *U64) IsEqualTo(v *U64) (r int) {
	x := u.cmp(v)
	if x == 0 {
		r = 1
	} else {
		r = 0
	}
	return
}

func (u *U64) EQ(v *U64) (r int) {
	x := u.cmp(v)
	if x == 0 {
		r = 1
	} else {
		r = 0
	}
	return
}

func (u *U64) IsGreaterThan(v *U64) (r int) {
	x := u.cmp(v)
	if x == 1 {
		r = 1
	} else {
		r = 0
	}
	return
}

func (u *U64) GT(v *U64) (r int) {
	x := u.cmp(v)
	if x == 1 {
		r = 1
	} else {
		r = 0
	}
	return
}

func (u *U64) IsGreaterThanOrEqualTo(v *U64) (r int) {
	x := u.cmp(v)
	if x == -1 {
		r = 0
	} else {
		r = 1
	}
	return
}

func (u *U64) GTE(v *U64) (r int) {
	x := u.cmp(v)
	if x == -1 {
		r = 0
	} else {
		r = 1
	}
	return
}

func (u *U64) IsLessThan(v *U64) (r int) {
	x := u.cmp(v)
	if x == -1 {
		r = 1
	} else {
		r = 0
	}
	return
}

func (u *U64) LT(v *U64) (r int) {
	x := u.cmp(v)
	if x == -1 {
		r = 1
	} else {
		r = 0
	}
	return
}

func (u *U64) IsLessThanOrEqualTo(v *U64) (r int) {
	x := u.cmp(v)
	if x == 1 {
		r = 0
	} else {
		r = 1
	}
	return
}

func (u *U64) LTE(v *U64) (r int) {
	x := u.cmp(v)
	if x == 1 {
		r = 0
	} else {
		r = 1
	}
	return
}

func (u U64) toJSON() string {
	u1 := big.Int(u)
	innerArrOfU := u1.Bytes()
	var test = make([]byte, 2*len(innerArrOfU))
	hex.Encode(test, innerArrOfU[:])
	return "0x" + string(test)
}

func (u U128) String() string {
	var a = big.Int(u)
	return a.String()
}

func NewU128(v string) U128 {
	var x big.Int
	x.SetString(v, 10)
	return U128(x)
}

func (u *U128) set(y big.Int) {
	*u = U128(y)
}

func (u *U128) Add(x *U128, y *U128) *U128 {
	// Error larger than max
	x1 := big.Int(*x)
	y1 := big.Int(*y)
	z1 := big.Int(*u)
	z1.Add(&x1, &y1)
	u.set(z1)
	return u
}

func (u *U128) Sub(x *U128, y *U128) *U128 {
	// Error less than zero
	x1 := big.Int(*x)
	y1 := big.Int(*y)
	z1 := big.Int(*u)
	z1.Sub(&x1, &y1)
	u.set(z1)
	return u
}

func (u *U128) Mul(x *U128, y *U128) *U128 {
	x1 := big.Int(*x)
	y1 := big.Int(*y)
	z1 := big.Int(*u)
	z1.Sub(&x1, &y1)
	u.set(z1)
	return u
}

func (u *U128) Idiv(x *U128, y *U128) *U128 {
	x1 := big.Int(*x)
	y1 := big.Int(*y)
	z1 := big.Int(*u)
	z1.Quo(&x1, &y1)
	u.set(z1)
	return u
}

func (u *U128) Mod(x *U128, y *U128) *U128 {
	x1 := big.Int(*x)
	y1 := big.Int(*y)
	z1 := big.Int(*u)
	z1.Mod(&x1, &y1)
	u.set(z1)
	return u
}

func (u U128) RlpBytes() []byte {
	z1 := big.Int(u)
	x, _ := rlp.EncodeToBytes(z1.Bytes())
	return x
}

func (u *U128) cmp(v *U128) (r int) {
	u1 := big.Int(*u)
	v1 := big.Int(*v)
	r = u1.Cmp(&v1)
	return r
}

func (u *U128) IsEqualTo(v *U128) (r int) {
	x := u.cmp(v)
	if x == 0 {
		r = 1
	} else {
		r = 0
	}
	return
}

func (u *U128) EQ(v *U128) (r int) {
	x := u.cmp(v)
	if x == 0 {
		r = 1
	} else {
		r = 0
	}
	return
}

func (u *U128) IsGreaterThan(v *U128) (r int) {
	x := u.cmp(v)
	if x == 1 {
		r = 1
	} else {
		r = 0
	}
	return
}

func (u *U128) GT(v *U128) (r int) {
	x := u.cmp(v)
	if x == 1 {
		r = 1
	} else {
		r = 0
	}
	return
}

func (u *U128) IsGreaterThanOrEqualTo(v *U128) (r int) {
	x := u.cmp(v)
	if x == -1 {
		r = 0
	} else {
		r = 1
	}
	return
}

func (u *U128) GTE(v *U128) (r int) {
	x := u.cmp(v)
	if x == -1 {
		r = 0
	} else {
		r = 1
	}
	return
}

func (u *U128) IsLessThan(v *U128) (r int) {
	x := u.cmp(v)
	if x == -1 {
		r = 1
	} else {
		r = 0
	}
	return
}

func (u *U128) LT(v *U128) (r int) {
	x := u.cmp(v)
	if x == -1 {
		r = 1
	} else {
		r = 0
	}
	return
}

func (u *U128) IsLessThanOrEqualTo(v *U128) (r int) {
	x := u.cmp(v)
	if x == 1 {
		r = 0
	} else {
		r = 1
	}
	return
}

func (u *U128) LTE(v *U128) (r int) {
	x := u.cmp(v)
	if x == 1 {
		r = 0
	} else {
		r = 1
	}
	return
}

func (u U128) toJSON() string {
	u1 := big.Int(u)
	innerArrOfU := u1.Bytes()
	var test = make([]byte, 2*len(innerArrOfU))
	hex.Encode(test, innerArrOfU[:])
	return "0x" + string(test)
}

func (u U256) String() string {
	var a = big.Int(u)
	return a.String()
}

func NewU256(v string) U256 {
	var x big.Int
	x.SetString(v, 10)
	return U256(x)
}

func (u *U256) set(y big.Int) {
	*u = U256(y)
}

func (u *U256) Add(x *U256, y *U256) *U256 {
	// Error larger than max
	x1 := big.Int(*x)
	y1 := big.Int(*y)
	z1 := big.Int(*u)
	z1.Add(&x1, &y1)
	u.set(z1)
	return u
}

func (u *U256) Sub(x *U256, y *U256) *U256 {
	// Error less than zero
	x1 := big.Int(*x)
	y1 := big.Int(*y)
	z1 := big.Int(*u)
	z1.Sub(&x1, &y1)
	u.set(z1)
	return u
}

func (u *U256) Mul(x *U256, y *U256) *U256 {
	x1 := big.Int(*x)
	y1 := big.Int(*y)
	z1 := big.Int(*u)
	z1.Sub(&x1, &y1)
	u.set(z1)
	return u
}

func (u *U256) Idiv(x *U256, y *U256) *U256 {
	x1 := big.Int(*x)
	y1 := big.Int(*y)
	z1 := big.Int(*u)
	z1.Quo(&x1, &y1)
	u.set(z1)
	return u
}

func (u *U256) Mod(x *U256, y *U256) *U256 {
	x1 := big.Int(*x)
	y1 := big.Int(*y)
	z1 := big.Int(*u)
	z1.Mod(&x1, &y1)
	u.set(z1)
	return u
}

func (u U256) RlpBytes() []byte {
	z1 := big.Int(u)
	x, _ := rlp.EncodeToBytes(z1.Bytes())
	return x
}

func (u *U256) cmp(v *U256) (r int) {
	u1 := big.Int(*u)
	v1 := big.Int(*v)
	r = u1.Cmp(&v1)
	return r
}

func (u *U256) IsEqualTo(v *U256) (r int) {
	x := u.cmp(v)
	if x == 0 {
		r = 1
	} else {
		r = 0
	}
	return
}

func (u *U256) EQ(v *U256) (r int) {
	x := u.cmp(v)
	if x == 0 {
		r = 1
	} else {
		r = 0
	}
	return
}

func (u *U256) IsGreaterThan(v *U256) (r int) {
	x := u.cmp(v)
	if x == 1 {
		r = 1
	} else {
		r = 0
	}
	return
}

func (u *U256) GT(v *U256) (r int) {
	x := u.cmp(v)
	if x == 1 {
		r = 1
	} else {
		r = 0
	}
	return
}

func (u *U256) IsGreaterThanOrEqualTo(v *U256) (r int) {
	x := u.cmp(v)
	if x == -1 {
		r = 0
	} else {
		r = 1
	}
	return
}

func (u *U256) GTE(v *U256) (r int) {
	x := u.cmp(v)
	if x == -1 {
		r = 0
	} else {
		r = 1
	}
	return
}

func (u *U256) IsLessThan(v *U256) (r int) {
	x := u.cmp(v)
	if x == -1 {
		r = 1
	} else {
		r = 0
	}
	return
}

func (u *U256) LT(v *U256) (r int) {
	x := u.cmp(v)
	if x == -1 {
		r = 1
	} else {
		r = 0
	}
	return
}

func (u *U256) IsLessThanOrEqualTo(v *U256) (r int) {
	x := u.cmp(v)
	if x == 1 {
		r = 0
	} else {
		r = 1
	}
	return
}

func (u *U256) LTE(v *U256) (r int) {
	x := u.cmp(v)
	if x == 1 {
		r = 0
	} else {
		r = 1
	}
	return
}

func (u U256) toJSON() string {
	u1 := big.Int(u)
	innerArrOfU := u1.Bytes()
	var test = make([]byte, 2*len(innerArrOfU))
	hex.Encode(test, innerArrOfU[:])
	return "0x" + string(test)
}
