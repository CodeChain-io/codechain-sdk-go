package rpc

import (
	"testing"
)

func TestAccount(t *testing.T) {
	account := NewRPC("https://corgi-rpc.codechain.io/").Account
	account.Create("asdfasdf")
	account.GetList()
}
