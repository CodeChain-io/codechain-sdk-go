package rpc

import (
	"fmt"
	"testing"
)

func TestChain(t *testing.T) {
	chain := NewRPC("https://corgi-rpc.codechain.io/").chain
	res := chain.GetBlockByNumber(5)
	fmt.Println(res.Hash)
}
