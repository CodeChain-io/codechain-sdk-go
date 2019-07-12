package rpc

func Ping(params interface{}, id ...interface{}) string {
	const method = "ping"
	return call(callInterface{node: "https://corgi-rpc.codechain.io/", method: method, id: id})
	// TODO: Use node uri from environment
}
