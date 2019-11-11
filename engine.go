package rpc

type Engine struct {
	rpcClient rpcClient
}

func (e *Engine) GetCoinbase() (string, error) {
	const method = "engine_getCoinbase"
	var result string
	err := e.rpcClient.call(callInterface{method: method, id: ""}, &result)
	return result, err
}

func (e *Engine) GetBlockReward(blockNumber int) (interface{}, error) {
	const method = "engine_getBlockReward"
	var result interface{}
	err := e.rpcClient.call(callInterface{method: method, id: ""}, &result, blockNumber)
	return result, err
}

func (e *Engine) GetRecommendedConfirmation() (int, error) {
	const method = "engine_getRecommendedConfirmation"
	var result int
	err := e.rpcClient.call(callInterface{method: method, id: ""}, &result)
	return result, err
}

func (e *Engine) GetCustomActionData(handerID int, bytes string, blockNumber int) (string, error) {
	const method = "engine_getCustomActionData"
	var result string
	err := e.rpcClient.call(callInterface{method: method, id: ""}, &result, handerID, bytes, blockNumber)
	return result, err
}
