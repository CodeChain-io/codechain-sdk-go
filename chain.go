package rpc

type Chain struct {
	rpcClient rpcClient
}

func (c *Chain) GetBestBlockNumber() int {
	const method = "chain_getBestBlockNumber"
	response := c.rpcClient.call(callInterface{method: method, id: ""})
	return response.Result.(int)
}

func (c *Chain) GetBestBlockId() string {
	const method = "chain_getBestBlockId"
	response := c.rpcClient.call(callInterface{method: method, id: ""})
	return response.Result.(string)
}

func (c *Chain) GetBlockHash(blockNumber int) string {
	const method = "chain_getBlockHash"
	response := c.rpcClient.call(callInterface{method: method, id: ""}, blockNumber)
	return response.Result.(string)
}

func (c *Chain) GetBlockByNumber(blockNumber int) interface{} {
	const method = "chain_getBlockByNumber"
	response := c.rpcClient.call(callInterface{method: method, id: ""}, blockNumber)
	return response.Result
}

func (c *Chain) GetBlockByHash(blockHash string) interface{} {
	const method = "chain_getBlockByHash"
	response := c.rpcClient.call(callInterface{method: method, id: ""}, blockHash)
	return response.Result
}

func (c *Chain) GetBlockTransactionCountByHash(blockHash string) int {
	const method = "chain_getBlockTransactionCountByHash"
	response := c.rpcClient.call(callInterface{method: method, id: ""}, blockHash)
	return response.Result.(int)
}

func (c *Chain) GetTransaction(transactionHash string) interface{} {
	const method = "chain_getTransaction"
	response := c.rpcClient.call(callInterface{method: method, id: ""}, transactionHash)
	return response.Result
}

func (c *Chain) GetTransactionSigner(transactionHash string) string {
	const method = "chain_getTransactionSigner"
	response := c.rpcClient.call(callInterface{method: method, id: ""}, transactionHash)
	return response.Result.(string)
}

func (c *Chain) ContainsTransaction(transactionHash string) bool {
	const method = "chain_containsTransaction"
	response := c.rpcClient.call(callInterface{method: method, id: ""}, transactionHash)
	return response.Result.(bool)
}

func (c *Chain) GetTransactionByTracker(tracker string) interface{} {
	const method = "chain_getTransactionByTracker"
	response := c.rpcClient.call(callInterface{method: method, id: ""}, tracker)
	return response.Result
}

func (c *Chain) GetAssetSchemeByTracker(tracker string) interface{} {
	const method = "chain_getAssetSchemeByTracker"
	response := c.rpcClient.call(callInterface{method: method, id: ""}, tracker)
	return response.Result
}

func (c *Chain) GetAssetSchemeByType(assetType string) interface{} {
	const method = "chain_getAssetSchemeByType"
	response := c.rpcClient.call(callInterface{method: method, id: ""}, assetType)
	return response.Result
}

func (c *Chain) GetAsset(tracker string, transactionIndex int, shardId int, blockNumber int) interface{} {
	const method = "chain_getAsset"
	response := c.rpcClient.call(callInterface{method: method, id: ""}, tracker, transactionIndex, shardId, blockNumber)
	return response.Result
}

func (c *Chain) GetText(transactionHash string, blockNumber int) interface{} {
	const method = "chain_getText"
	response := c.rpcClient.call(callInterface{method: method, id: ""}, transactionHash, blockNumber)
	return response.Result
}

func (c *Chain) IsAssetSpent(tracker string, transactionIndex int, shardId int, blockNumber int) interface{} {
	const method = "chain_isAssetSpent"
	response := c.rpcClient.call(callInterface{method: method, id: ""}, tracker, transactionIndex, shardId, blockNumber)
	return response.Result
}

func (c *Chain) GetSeq(address string, blockNumber int) int {
	const method = "chain_getSeq"
	response := c.rpcClient.call(callInterface{method: method, id: ""}, address, blockNumber)
	return response.Result.(int)
}

func (c *Chain) GetBalance(address string, blockNumber int) int {
	const method = "chain_getBalance"
	response := c.rpcClient.call(callInterface{method: method, id: ""}, address, blockNumber)
	return response.Result.(int)
}

func (c *Chain) GetRegularKey(address string, blockNumber int) string {
	const method = "chain_getRegularKey"
	response := c.rpcClient.call(callInterface{method: method, id: ""}, address, blockNumber)
	return response.Result.(string)
}

func (c *Chain) GetRegularKeyOwner(publicKey string, blockNumber int) string {
	const method = "chain_getRegularKey"
	response := c.rpcClient.call(callInterface{method: method, id: ""}, publicKey, blockNumber)
	return response.Result.(string)
}

func (c *Chain) GetGenesisAccounts() []interface{} {
	const method = "chain_getGenesisAccounts"
	response := c.rpcClient.call(callInterface{method: method, id: ""})
	return response.Result.([]interface{})
}

func (c *Chain) GetNumberOfShards(blockNumber int) float64 {
	const method = "chain_getNumberOfShards"
	response := c.rpcClient.call(callInterface{method: method, id: ""}, blockNumber)
	return response.Result.(float64)
}

func (c *Chain) GetShardIdByHash(transactionHash string, blockNumber int) int {
	const method = "chain_getShardIdByHash"
	response := c.rpcClient.call(callInterface{method: method, id: ""}, transactionHash, blockNumber)
	return response.Result.(int)
}

func (c *Chain) GetShardRoot(shardId int, blockNumber int) string {
	const method = "chain_getShardRoot"
	response := c.rpcClient.call(callInterface{method: method, id: ""}, shardId, blockNumber)
	return response.Result.(string)
}

func (c *Chain) GetShardOwners(shardId int, blockNumber int) []string {
	const method = "chain_getShardOwners"
	response := c.rpcClient.call(callInterface{method: method, id: ""}, shardId, blockNumber)
	return response.Result.([]string)
}

func (c *Chain) GetShardUsers(shardId int, blockNumber int) []string {
	const method = "chain_getShardUsers"
	response := c.rpcClient.call(callInterface{method: method, id: ""}, shardId, blockNumber)
	return response.Result.([]string)

}

func (c *Chain) GetMiningReward(blockNumber int) int {
	const method = "chain_getMiningReward"
	response := c.rpcClient.call(callInterface{method: method, id: ""}, blockNumber)
	return response.Result.(int)
}

func (c *Chain) GetMinTransactionFee(transactionType string, blockNumber int) int {
	const method = "chain_getMinTransactionFee"
	response := c.rpcClient.call(callInterface{method: method, id: ""}, transactionType, blockNumber)
	return response.Result.(int)
}

func (c *Chain) GetCommonParams(blockNumber int) interface{} {
	const method = "chain_getCommonParamss"
	response := c.rpcClient.call(callInterface{method: method, id: ""}, blockNumber)
	return response.Result
}

func (c *Chain) GetTermMetadata(blockNumber int) interface{} {
	const method = "chain_getTermMetadata"
	response := c.rpcClient.call(callInterface{method: method, id: ""}, blockNumber)
	return response.Result
}

func (c *Chain) ExecuteTransaction(transaction interface{}, sender string) string {
	const method = "chain_getTermMetadata"
	response := c.rpcClient.call(callInterface{method: method, id: ""}, transaction, sender)
	return response.Result.(string)
}

func (c *Chain) ExecuteVM(transaction interface{}, parameters [][][]int, indices int) []string {
	const method = "chain_executeVM"
	response := c.rpcClient.call(callInterface{method: method, id: ""}, transaction, parameters, indices)
	return response.Result.([]string)
}

func (c *Chain) GetNetworkId() string {
	const method = "chain_getNetworkId"
	response := c.rpcClient.call(callInterface{method: method, id: ""})
	return response.Result.(string)
}
