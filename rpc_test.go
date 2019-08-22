package rpc

import (
	"fmt"
	"testing"
)

func TestRPC(t *testing.T) {
	Init("https://corgi-rpc.codechain.io/")
	Ping()
	fmt.Println(Version())
}
