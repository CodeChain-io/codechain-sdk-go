package rpc

import (
	"testing"
)

func TestCall(t *testing.T) {
	r1 := call(callInterface{node: "https://corgi-rpc.codechain.io/", method: "ping", id: nil})
	if r1 != "{\"jsonrpc\":\"2.0\",\"result\":\"pong\",\"id\":null}\n" {
		t.Error("Error")
	}
	r2 := call(callInterface{node: "https://corgi-rpc.codechain.i/", method: "ping", id: nil})
	if r2 != "" {
		t.Error("Error2")
	}
}
