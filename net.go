package rpc

type Net struct {
	rpcClient rpcClient
}

func (n *Net) LocalKeyFor(address string, port int) (string, error) {
	const method = "net_localKeyFor"
	var result string
	err := n.rpcClient.call(callInterface{method: method, id: ""}, &result, address, port)
	return result, err
}

func (n *Net) RegisterRemoteKeyFor(address string, port int, remotePublicKey string) (string, error) {
	const method = "net_registerRemoteKeyFor"
	var result string
	err := n.rpcClient.call(callInterface{method: method, id: ""}, &result, address, port, remotePublicKey)
	return result, err
}

func (n *Net) Connect(address string, port int) error {
	const method = "net_connect"
	return n.rpcClient.call(callInterface{method: method, id: ""}, nil, address, port)
}

func (n *Net) IsConnected(address string, port int) (bool, error) {
	const method = "net_isConnected"
	var result bool
	err := n.rpcClient.call(callInterface{method: method, id: ""}, &result, address, port)
	return result, err
}

func (n *Net) Disconnect(address string, port int) error {
	const method = "net_disconnect"
	return n.rpcClient.call(callInterface{method: method, id: ""}, nil, address, port)
}

func (n *Net) GetPeerCount() (int, error) {
	const method = "net_getPeerCount"
	var result int
	err := n.rpcClient.call(callInterface{method: method, id: ""}, &result)
	return result, err
}

func (n *Net) GetEstablishedPeers() ([]string, error) {
	const method = "net_getEstablishedPeers"
	var result []string
	err := n.rpcClient.call(callInterface{method: method, id: ""}, &result)
	return result, err
}

func (n *Net) GetPort() (int, error) {
	const method = "net_getPort"
	var port int
	err := n.rpcClient.call(callInterface{method: method, id: ""}, &port)
	return port, err
}

func (n *Net) AddToWhitelist(address string, tag string) error {
	const method = "net_addToWhitelist"
	return n.rpcClient.call(callInterface{method: method, id: ""}, nil, address, tag)
}

func (n *Net) RemoveFromWhitelist(address string) error {
	const method = "net_removeFromWhitelist"
	return n.rpcClient.call(callInterface{method: method, id: ""}, nil, address)
}

func (n *Net) AddToBlacklist(address string, tag string) error {
	const method = "net_addToBlacklist"
	return n.rpcClient.call(callInterface{method: method, id: ""}, nil, address, tag)
}

func (n *Net) RemoveFromBlacklist(address string) error {
	const method = "net_removeFromBlacklist"
	return n.rpcClient.call(callInterface{method: method, id: ""}, nil, address)
}

func (n *Net) EnableWhitelist() error {
	const method = "net_enableWhitelist"
	return n.rpcClient.call(callInterface{method: method, id: ""}, nil)
}

func (n *Net) DisableWhitelist() error {
	const method = "net_disableWhitelist"
	return n.rpcClient.call(callInterface{method: method, id: ""}, nil)
}

func (n *Net) EnableBlacklist() error {
	const method = "net_enableBlacklist"
	return n.rpcClient.call(callInterface{method: method, id: ""}, nil)
}

func (n *Net) DisableBlacklist() error {
	const method = "net_disableBlacklist"
	return n.rpcClient.call(callInterface{method: method, id: ""}, nil)
}

func (n *Net) GetWhitelist() (interface{}, error) {
	const method = "net_getWhitelist"
	var reusult interface{}
	err := n.rpcClient.call(callInterface{method: method, id: ""}, &reusult)
	return reusult, err
}

func (n *Net) GetBlacklist() (interface{}, error) {
	const method = "net_getBlacklist"
	var reusult interface{}
	err := n.rpcClient.call(callInterface{method: method, id: ""}, &reusult)
	return reusult, err
}

func (n *Net) RecentNetworkUsage() (interface{}, error) {
	const method = "net_recentNetworkUsage"
	var reusult interface{}
	err := n.rpcClient.call(callInterface{method: method, id: ""}, &reusult)
	return reusult, err
}
