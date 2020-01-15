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

	signerAccountID, err := primitives.PlatformAddressFromString("tccq9h7vnl68frvqapzv3tujrxtxtwqdnxw6yamrrgd")
	if err != nil {
		panic(err)
	}

	signerSecret, err := primitives.StringToH256("ede1d4ccb4ec9a8bbbae9a13db3f4a7b56ea04189be86ac3a6a439d9a0a1addd")
	if err != nil {
		panic(err)
	}

	asestPrivateKey1, err1 := key.GenerateEcdsa()
	assetPrivateKey2, err2 := key.GenerateEcdsa()

	if err1 != nil {
		panic(err1)
	}
	if err2 != nil {
		panic(err2)
	}

	assetAddress1, err1 := key.CreateAssetAddress(asestPrivateKey1, networkID)
	assetAddress2, err2 := key.CreateAssetAddress(assetPrivateKey2, networkID)

	if err1 != nil {
		panic(err1)
	}
	if err2 != nil {
		panic(err2)
	}

	supply := primitives.NewU64("5000")
	output, err := core.NewAssetMintOutputWithRecipient(assetAddress1, supply)

	if err != nil {
		panic(err)
	}

	random := make([]byte, 4)
	rand.Read(random)
	assetName := "{\"name\": \"example_coin_" + hex.EncodeToString(random) + "\"}"
	shardID := uint(0)

	mintTransaction := transaction.NewMintAsset(networkID, shardID, assetName, output, nil, nil, []primitives.H160{}, []string{})

	seq, err := rpcClient.Chain.GetLatestSeq(signerAccountID.Value)
	if err != nil {
		panic(err)
	}
	fee := primitives.NewU64("200000")
	signed := mintTransaction.Sign(signerSecret, seq, fee)

	if _, err := rpcClient.Mempool.SendSignedTransaction(hex.EncodeToString(signed.RlpBytes())); err != nil {
		panic(err)
	}

	tracker := mintTransaction.Tracker()

	asset1, err := rpcClient.Chain.GetAsset(tracker.ToJSON(), 0, 0)

	if err != nil {
		panic(err)
	}

	assetType, err := primitives.StringToH160(asset1.AssetType)
	if err != nil {
		panic(err)
	}
	lockScriptHash1, err := primitives.StringToH160(asset1.LockScriptHash)
	if err != nil {
		panic(err)
	}

	inputasset1 := core.NewAsset(assetType,
		uint(0), lockScriptHash1,
		asset1.Parameters, primitives.NewU64WithHex(asset1.Quantity),
		tracker,
		uint(0))

	transfer := transaction.NewTransferAsset(
		[]core.AssetTransferInput{},
		[]core.AssetTransferInput{},
		[]core.AssetTransferOutput{}, networkID, "", []string{}, nil)

	transfer.AddInputWithAsset(inputasset1)

	q1 := primitives.NewU64("2500")
	q2 := primitives.NewU64("2500")

	output1, err1 := core.NewAssetTransferOutputWithAddress(assetAddress1, inputasset1.AssetType, uint(0), q1)
	output2, err2 := core.NewAssetTransferOutputWithAddress(assetAddress2, inputasset1.AssetType, uint(0), q2)
	if err1 != nil {
		panic(err1)
	}
	if err2 != nil {
		panic(err2)
	}
	outputs := []core.AssetTransferOutput{output1, output2}
	transfer.AddOutputs(outputs)

	if err := transfer.SignTransactionInput(0, asestPrivateKey1.GetPublicKey(), asestPrivateKey1.GetPrivateKey()); err != nil {
		panic(err)
	}

	seq, err = rpcClient.Chain.GetLatestSeq(signerAccountID.Value)
	if err != nil {
		panic(err)
	}

	fee = primitives.NewU64("2000")
	signedTransfer := transfer.Sign(signerSecret, seq, fee)

	res, err := rpcClient.Mempool.SendSignedTransaction(hex.EncodeToString(signedTransfer.RlpBytes()))

	if err != nil {
		panic(err)
	}

	fmt.Println("Success to send transferAsset transaction")
	fmt.Println("Recipient1 address: ", assetAddress1.Value)
	fmt.Println("Recipient1 secret: ", hex.EncodeToString(asestPrivateKey1.GetPrivateKey()))
	fmt.Println("Recipient2 address: ", assetAddress2.Value)
	fmt.Println("Recipient2 secret: ", hex.EncodeToString(assetPrivateKey2.GetPrivateKey()))
	fmt.Println("TransferAsset transaction hash: ", res)
	fmt.Println("TransferAsset tracker: ", hex.EncodeToString(transfer.Tracker().Bytes()))

	return
}
