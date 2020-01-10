package main

import (
	"encoding/hex"
	"fmt"

	"github.com/CodeChain-io/codechain-sdk-go/core/transaction"
	"github.com/CodeChain-io/codechain-sdk-go/key"
	"github.com/CodeChain-io/codechain-sdk-go/primitives"
	"github.com/CodeChain-io/codechain-sdk-go/rpc"
)

func main() {

	rpcClient := rpc.NewRPC("http://localhost:8080")
	networkID := "tc"

	signerID, _ := primitives.PlatformAddressFromString("tccq9h7vnl68frvqapzv3tujrxtxtwqdnxw6yamrrgd")
	signerSecret, _ := primitives.StringToH256("ede1d4ccb4ec9a8bbbae9a13db3f4a7b56ea04189be86ac3a6a439d9a0a1addd")

	ecdsa, _ := key.GenerateEcdsa()
	recipient, _ := key.CreatePlatformAddress(ecdsa, networkID)

	quantity := primitives.NewU64("100")
	fee := primitives.NewU64("100")
	seq, _ := rpcClient.Chain.GetLatestSeq(signerID.Value)

	pay := transaction.NewPay(recipient, quantity, networkID)
	signed := pay.Sign(signerSecret, seq, fee)

	res, err := rpcClient.Mempool.SendSignedTransaction(hex.EncodeToString(signed.RlpBytes()))

	if err != nil {
		fmt.Println("Failed to send pay transaction: ", err)
		return
	}

	fmt.Println("Success to send pay transaction")
	fmt.Println("Recipient address: ", recipient.Value)
	fmt.Println("Recipient secret: ", hex.EncodeToString(ecdsa.GetPrivateKey()))
	fmt.Println("Pay transaction hash: ", res)

	return
}
