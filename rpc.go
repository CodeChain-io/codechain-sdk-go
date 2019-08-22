package rpc

// NewRPC is a constructor of RPC
func Init(node string) {
	newRPCClient(node, nil)
}

// Ping sends request to node
func Ping() {
	const method = "ping"
	call(callInterface{method: method, id: ""})
}

// Version checking
func Version() string {
	const method = "version"
	result := call(callInterface{method: method, id: ""}).Result.(string)
	return result
}
