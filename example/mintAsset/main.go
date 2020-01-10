package main

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"

	"github.com/CodeChain-io/codechain-sdk-go/core"
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
	assetAddress, err := key.CreateAssetAddress(ecdsa, networkID)

	if err != nil {
		fmt.Println("AssetAdress creation error: ", err)
		return
	}

	supply := primitives.NewU64("5000")
	output, _ := core.NewAssetMintOutputWithRecipient(assetAddress, supply)

	random := make([]byte, 4)
	rand.Read(random)
	assetName := "{\"name\": \"example_coin_" + hex.EncodeToString(random) + "\"}"
	shardID := uint(0)

	mintTransaction := transaction.NewMintAsset(networkID, shardID, assetName, output, nil, nil, []primitives.H160{}, []string{})

	seq, _ := rpcClient.Chain.GetLatestSeq(signerID.Value)
	fee := primitives.NewU64("200000")
	signed := mintTransaction.Sign(signerSecret, seq, fee)

	res, err := rpcClient.Mempool.SendSignedTransaction(hex.EncodeToString(signed.RlpBytes()))

	if err != nil {
		fmt.Println("Failed to send mint transaction: ", err)
		return
	}

	fmt.Println("Success to send mint transaction")
	fmt.Println("Recipient address: ", assetAddress.Value)
	fmt.Println("Recipient secret: ", hex.EncodeToString(ecdsa.GetPrivateKey()))
	fmt.Println("MintAsset transaction hash: ", res)
	fmt.Println("MintAsset tracker", hex.EncodeToString(mintTransaction.Tracker().Bytes()))

	return
}
