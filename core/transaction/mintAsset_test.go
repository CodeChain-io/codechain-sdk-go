package transaction

import (
	"encoding/hex"
	"testing"

	"github.com/CodeChain-io/codechain-sdk-go/core"
	"github.com/CodeChain-io/codechain-sdk-go/key"
	"github.com/CodeChain-io/codechain-sdk-go/primitives"
	"github.com/CodeChain-io/codechain-sdk-go/rpc"
)

func TestMint(t *testing.T) {

	rpcClient := rpc.NewRPC("https://corgi-rpc.codechain.io/")

	accountID, _ := primitives.PlatformAddressFromString("wccqxhpcpeslxhrwn9mq4cdrfaug2079vsqwvaqxudn")

	secret, _ := primitives.StringToH256("94dc7a29037d623acdd9eb2f1c5d8e88e019bf83c3ee930757ce1e40b0b50634")

	seq, _ := rpcClient.Chain.GetLatestSeq(accountID.Value)

	asset, _ := key.CreateAssetAddress("wc")
	output, _ := core.NewAssetMintOutputWithRecipient(asset, primitives.NewU64("4422333"))
	mint := NewMintAsset("wc", uint(0), "asdfasdf", output, nil, nil, []primitives.H160{}, []string{})

	fee := primitives.NewU64("200000")

	signed := mint.Sign(secret, seq, fee)
	_, err5 := rpcClient.Mempool.SendSignedTransaction(hex.EncodeToString(signed.RlpBytes()))

	if err5 != nil {
		t.Fatal(err5)
	}
}
