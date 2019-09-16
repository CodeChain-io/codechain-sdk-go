package rpc

import (
	"fmt"
	"testing"
)

func TestHaHa(t *testing.T) {
	account := NewRPC("https://corgi-rpc.codechain.io/").account
	account.GetList()
	fmt.Println(account.Create("asdfasdf"))
}
