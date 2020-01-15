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

	signerAccountID, _ := primitives.PlatformAddressFromString("tccq9h7vnl68frvqapzv3tujrxtxtwqdnxw6yamrrgd")
	signerSecret, _ := primitives.StringToH256("ede1d4ccb4ec9a8bbbae9a13db3f4a7b56ea04189be86ac3a6a439d9a0a1addd")

	ecdsa1, _ := key.GenerateEcdsa()
	ecdsa2, _ := key.GenerateEcdsa()
	assetAddress1, err1 := key.CreateAssetAddress(ecdsa1, networkID)
	assetAddress2, err2 := key.CreateAssetAddress(ecdsa2, networkID)

	if err1 != nil || err2 != nil {
		fmt.Println("AssetAdress creation error: ", err1, err2)
		return
	}

	supply := primitives.NewU64("5000")
	output, _ := core.NewAssetMintOutputWithRecipient(assetAddress1, supply)

	random := make([]byte, 4)
	rand.Read(random)
	assetName := "{\"name\": \"example_coin_" + hex.EncodeToString(random) + "\"}"
	shardID := uint(0)

	mintTransaction := transaction.NewMintAsset(networkID, shardID, assetName, output, nil, nil, []primitives.H160{}, []string{})

	seq, _ := rpcClient.Chain.GetLatestSeq(signerAccountID.Value)
	fee := primitives.NewU64("200000")
	signed := mintTransaction.Sign(signerSecret, seq, fee)

	_, err := rpcClient.Mempool.SendSignedTransaction(hex.EncodeToString(signed.RlpBytes()))

	if err != nil {
		fmt.Println("Failed to send mint transaction: ", err)
		return
	}

	tracker := mintTransaction.Tracker()

	asset1, _ := rpcClient.Chain.GetAsset(tracker.ToJSON(), 0, 0)

	assetType, _ := primitives.StringToH160(asset1.AssetType)
	lockScriptHash1, _ := primitives.StringToH160(asset1.LockScriptHash)

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

	output1, _ := core.NewAssetTransferOutputWithAddress(assetAddress1, inputasset1.AssetType, uint(0), q1)
	output2, _ := core.NewAssetTransferOutputWithAddress(assetAddress2, inputasset1.AssetType, uint(0), q2)
	outputs := []core.AssetTransferOutput{output1, output2}
	transfer.AddOutputs(outputs)

	err3 := transfer.SignTransactionInput(0, ecdsa1.GetPublicKey(), ecdsa1.GetPrivateKey())
	if err3 != nil {
		fmt.Println(err3)
		return
	}

	seq, _ = rpcClient.Chain.GetLatestSeq(signerAccountID.Value)

	fee = primitives.NewU64("2000")
	signedTransfer := transfer.Sign(signerSecret, seq, fee)

	fmt.Println("Tracker:", hex.EncodeToString(signedTransfer.Unsigned.(*transaction.TransferAsset).Tracker().Bytes()))
	res, err5 := rpcClient.Mempool.SendSignedTransaction(hex.EncodeToString(signedTransfer.RlpBytes()))

	if err5 != nil {
		fmt.Println("Failed to send transferAsset transaction")
		return
	}

	fmt.Println("Success to send transferAsset transaction")
	fmt.Println("Recipient1 address: ", assetAddress1.Value)
	fmt.Println("Recipient1 secret: ", hex.EncodeToString(ecdsa1.GetPrivateKey()))
	fmt.Println("Recipient2 address: ", assetAddress2.Value)
	fmt.Println("Recipient2 secret: ", hex.EncodeToString(ecdsa2.GetPrivateKey()))
	fmt.Println("TransferAsset transaction hash: ", res)
	fmt.Println("TransferAsset tracker: ", hex.EncodeToString(transfer.Tracker().Bytes()))

	return
}
