package rpc

type Devel struct {
	rpcClient rpcClient
}

func (d *Devel) GetStateTrieKeys(offset int, limit int) ([]string, error) {
	const method = "devel_getStateTrieKeys"
	var keys []string
	err := d.rpcClient.call(callInterface{method: method, id: ""}, &keys, offset, limit)
	return keys, err
}

func (d *Devel) GetStateTrieValue(key string) ([]string, error) {
	const method = "devel_getStateTrieValue"
	var keys []string
	err := d.rpcClient.call(callInterface{method: method, id: ""}, &keys, key)
	return keys, err
}

func (d *Devel) StartSealing() error {
	const method = "devel_startSealing"
	return d.rpcClient.call(callInterface{method: method, id: ""}, nil)
}

func (d *Devel) stopSealing() error {
	const method = "devel_stopSealing"
	err := d.rpcClient.call(callInterface{method: method, id: ""}, nil)
	return err
}

func (d *Devel) getBlockSyncPeers() ([]string, error) {
	const method = "devel_getBlockSyncPeers"
	var result []string
	err := d.rpcClient.call(callInterface{method: method, id: ""}, &result)
	return result, err
}

func (d *Devel) testTPS(count int, seed int, option string) (int, error) {
	const method = "devel_testTPS"
	var result int
	err := d.rpcClient.call(callInterface{method: method, id: ""}, &result, count, seed, option)
	return result, err
}
