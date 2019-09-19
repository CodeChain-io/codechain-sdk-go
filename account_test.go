package rpc

import (
	"fmt"
	"testing"
)

func TestHaHa(t *testing.T) {
	Init("https://corgi-rpc.codechain.io/")
	getList()
	fmt.Println(create("asdfasdf"))
}
