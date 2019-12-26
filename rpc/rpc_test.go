package rpc

import (
	"fmt"
	"testing"
)

func TestRPC(t *testing.T) {
	rpc := NewRPC("https://corgi-rpc.codechain.io/")
	rpc.Ping()
	fmt.Println(rpc.Version())
}
