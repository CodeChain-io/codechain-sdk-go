package main

import (
	"encoding/hex"
	"fmt"

	"github.com/CodeChain-io/codechain-sdk-go/key"
)

func main() {
	ecdsa, err := key.GenerateEcdsa()
	if err != nil {
		fmt.Println("ECDSA key pair generation error")
		return
	}

	assetAddress, err := key.CreateAssetAddress(ecdsa, "wc")

	if err != nil {
		fmt.Println("AssetAddress creation error: ", err)
		return
	}
	fmt.Println("assetAddress: ", assetAddress.Value)
	fmt.Println("accountID: ", hex.EncodeToString(assetAddress.Payload.Bytes()))
	fmt.Println("private key: ", hex.EncodeToString(ecdsa.GetPrivateKey()))
	fmt.Println("public key: ", hex.EncodeToString(ecdsa.GetPublicKey()))
}
