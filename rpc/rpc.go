package rpc

// RPC object
type RPC struct {
	rpcClient rpcClient
	Account   Account
	Chain     Chain
	Devel     Devel
	Engine    Engine
	Mempool   Mempool
	Net       Net
}

// NewRPC is a constructor of RPC
func NewRPC(nodeURL string) RPC {
	rpcClient := newRPCClient(nodeURL, nil)
	return RPC{
		rpcClient: rpcClient,
		Account:   Account{rpcClient},
		Chain:     Chain{rpcClient},
		Devel:     Devel{rpcClient},
		Engine:    Engine{rpcClient},
		Mempool:   Mempool{rpcClient}}
}

// Ping sends request to node
func (rpc *RPC) Ping() error {
	const method = "ping"
	err := rpc.rpcClient.call(callInterface{method: method, id: ""}, nil)
	return err
}

// Version checking
func (rpc *RPC) Version() (string, error) {
	const method = "version"
	var version string
	err := rpc.rpcClient.call(callInterface{method: method, id: ""}, &version)
	return version, err
}
