package rpc

type Mempool struct {
	rpcClient rpcClient
}

func (m *Mempool) SendSignedTransaction(tx string) (string, error) {
	const method = "mempool_sendSignedTransaction"
	var result string
	err := m.rpcClient.call(callInterface{method: method, id: ""}, &result, tx)
	return result, err
}

func (m *Mempool) GetErrorHint(transactionHash string) (string, error) {
	const method = "mempool_getErrorHint"
	var result string
	err := m.rpcClient.call(callInterface{method: method, id: ""}, &result, transactionHash)
	return result, err

}

func (m *Mempool) GetTransactionResultsByTracker(tracker string) ([]bool, error) {
	const method = "mempool_getTransactionResultsByTracker"
	var result []bool
	err := m.rpcClient.call(callInterface{method: method, id: ""}, &result, tracker)
	return result, err
}

func (m *Mempool) GetPendingTransactions(from int, to int) (interface{}, error) {
	const method = "mempool_getPendingTransactions"
	var result interface{}
	err := m.rpcClient.call(callInterface{method: method, id: ""}, from, to)
	return result, err
}
