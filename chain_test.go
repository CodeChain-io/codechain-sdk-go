package rpc

import (
	"fmt"
	"testing"
)

func TestChain(t *testing.T) {
	chain := NewRPC("https://corgi-rpc.codechain.io/").chain
	res, err := chain.GetBlockByNumber(5)
	if err != nil {
		t.Fatal("TestChain failed")
	}
	fmt.Println(res.Hash)
}
