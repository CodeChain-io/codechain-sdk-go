package rpc

import (
	"testing"
)

func TestChain(t *testing.T) {
	chain := NewRPC("https://corgi-rpc.codechain.io/").Chain
	_, err := chain.GetBlockByNumber(5)
	if err != nil {
		t.Fatal("Chain test failed")
	}

	_, err = chain.GetLatestSeq("wccq9dddym9sc6rn3jgsnmp8qel57s5mwjq6v5ye68e")
	if err != nil {
		t.Fatal("GetLatestSeq() Failed")
	}
}
