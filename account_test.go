package rpc

import (
	"fmt"
	"testing"
)

func TestAccount(t *testing.T) {
	account := NewRPC("https://corgi-rpc.codechain.io/").Account
	fmt.Println(account.Create("asdfasdf"))
}
