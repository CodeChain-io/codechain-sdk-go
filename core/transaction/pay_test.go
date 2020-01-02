package transaction

import (
	"bytes"
	"encoding/hex"
	"testing"

	"github.com/CodeChain-io/codechain-sdk-go/primitives"
	"github.com/CodeChain-io/codechain-sdk-go/rpc"
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

func TestSigning(t *testing.T) {
	id, _ := hex.DecodeString("0000000000000000000000000000000000000000")
	accountID := primitives.NewH160(id)

	p, _ := primitives.PlatformAddressFromAccountID(accountID, "tc")
	q := primitives.NewU64("11")

	pay := NewPay(p, q, "tc")

	secret, _ := primitives.StringToH256("ede1d4ccb4ec9a8bbbae9a13db3f4a7b56ea04189be86ac3a6a439d9a0a1addd")
	fee := primitives.NewU64("0")

	signed := pay.Sign(secret, uint(0), fee)

	if bytes.Compare(
		signed.RlpBytes(),
		[]byte{248, 96, 128, 128, 130, 116, 99, 215, 2, 148, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 11, 184, 65, 63, 155, 207, 244, 132, 189, 95, 29, 85, 73, 249, 18, 249, 238, 175, 140, 47, 227, 73, 178, 87, 189, 226, 182, 31, 177, 3, 96, 19, 212, 228, 76, 32, 74, 66, 21, 210, 108, 184, 121, 234, 173, 32, 40, 254, 26, 120, 152, 228, 207, 154, 93, 151, 158, 179, 131, 224, 163, 132, 20, 13, 110, 4, 193, 1}) != 0 {
		t.Fatal("Signing Pay transaction Error")
	}
}

func TestSendSignedTransaction(t *testing.T) {
	rpcClient := rpc.NewRPC("https://corgi-rpc.codechain.io/")

	recipient, err1 := primitives.PlatformAddressFromString("wccq95dss644nd3xgddje38k2kdldzsy6ns5gmscxu2")
	if err1 != nil {
		t.Fatal(err1)
	}

	quantity := primitives.NewU64("12345")

	pay := NewPay(recipient, quantity, "wc")

	secret, err2 := primitives.StringToH256("6692200de85d2ae6fbb817c656cc9f5a0b8ee1ed82795289aaf4400f1d817d62")
	if err2 != nil {
		t.Fatal(err2)
	}

	fee := primitives.NewU64("100")
	accountID, err3 := primitives.PlatformAddressFromString("wccq9dddym9sc6rn3jgsnmp8qel57s5mwjq6v5ye68e")

	if err3 != nil {
		t.Fatal(err3)
	}

	seq, err4 := rpcClient.Chain.GetLatestSeq(accountID.Value)

	if err4 != nil {
		t.Fatal(err4)
	}

	signed := pay.Sign(secret, seq, fee)

	_, err5 := rpcClient.Mempool.SendSignedTransaction(hex.EncodeToString(signed.RlpBytes()))

	if err5 != nil {
		t.Fatal(err5)
	}

}
