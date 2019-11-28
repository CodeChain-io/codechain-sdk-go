package transaction

import (
	"bytes"
	"encoding/hex"
	"github.com/CodeChain-io/codechain-rpc-go/primitives"
	"testing"
)

func TestPay(t *testing.T) {

	id, _ := hex.DecodeString("0000000000000000000000000000000000000000")
	accountID := primitives.NewH160(id)

	p, _ := primitives.PlatformAddressFromAccountID(accountID, "tc")

	q := primitives.NewU64("11")
	pay := NewPay(p, q, "tc")
	fee := primitives.NewU64("0")
	pay.SetFee(fee)
	pay.SetSeq(0)

	if bytes.Compare(
		pay.RlpBytes(),
		[]byte{221, 128, 128, 130, 116, 99,
			215, 2, 148, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 11}) != 0 {
		t.Fatal("Pay transaction creation error")
	}

}
