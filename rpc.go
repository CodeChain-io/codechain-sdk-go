package rpc

// RPC object
type RPC struct {
	rpcClient rpcClient
	account   Account
	chain     Chain
}

// NewRPC is a constructor of RPC
func NewRPC(nodeURL string) RPC {
	rpcClient := newRPCClient(nodeURL, nil)
	return RPC{
		rpcClient: rpcClient,
		account:   Account{rpcClient},
		chain:     Chain{rpcClient},
	}
}

// Ping sends request to node
func (rpc *RPC) Ping() {
	const method = "ping"
	rpc.rpcClient.call(callInterface{method: method, id: ""})
}

// Version checking
func (rpc *RPC) Version() string {
	const method = "version"
	result := rpc.rpcClient.call(callInterface{method: method, id: ""}).Result.(string)
	return result
}
