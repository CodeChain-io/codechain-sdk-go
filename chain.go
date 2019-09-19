package rpc

func GetBestBlockNumber() int {
	const method = "chain_getBestBlockNumber"
	response := call(callInterface{method: method, id: ""})
	return response.Result.(int)
}

func GetBestBlockId() string {
	const method = "chain_getBestBlockId"
	response := call(callInterface{method: method, id: ""})
	return response.Result.(string)
}

func GetBlockHash(blockNumber int) string {
	const method = "chain_getBlockHash"
	response := call(callInterface{method: method, id: ""}, blockNumber)
	return response.Result.(string)
}

func GetBlockByNumber(blockNumber int) interface{} {
	const method = "chain_getBlockByNumber"
	response := call(callInterface{method: method, id: ""}, blockNumber)
	return response.Result
}

func GetBlockByHash(blockHash string) interface{} {
	const method = "chain_getBlockByHash"
	response := call(callInterface{method: method, id: ""}, blockHash)
	return response.Result
}

func GetBlockTransactionCountByHash(blockHash string) int {
	const method = "chain_getBlockTransactionCountByHash"
	response := call(callInterface{method: method, id: ""}, blockHash)
	return response.Result.(int)
}

func GetTransaction(transactionHash string) interface{} {
	const method = "chain_getTransaction"
	response := call(callInterface{method: method, id: ""}, transactionHash)
	return response.Result
}

func GetTransactionSigner(transactionHash string) string {
	const method = "chain_getTransactionSigner"
	response := call(callInterface{method: method, id: ""}, transactionHash)
	return response.Result.(string)
}

func ContainsTransaction(transactionHash string) bool {
	const method = "chain_containsTransaction"
	response := call(callInterface{method: method, id: ""}, transactionHash)
	return response.Result.(bool)
}

func GetTransactionByTracker(tracker string) interface{} {
	const method = "chain_getTransactionByTracker"
	response := call(callInterface{method: method, id: ""}, tracker)
	return response.Result
}

func GetAssetSchemeByTracker(tracker string) interface{} {
	const method = "chain_getAssetSchemeByTracker"
	response := call(callInterface{method: method, id: ""}, tracker)
	return response.Result
}

func GetAssetSchemeByType(assetType string) interface{} {
	const method = "chain_getAssetSchemeByType"
	response := call(callInterface{method: method, id: ""}, assetType)
	return response.Result
}

func GetAsset(tracker string, transactionIndex int, shardId int, blockNumber int) interface{} {
	const method = "chain_getAsset"
	response := call(callInterface{method: method, id: ""}, tracker, transactionIndex, shardId, blockNumber)
	return response.Result
}

func GetText(transactionHash string, blockNumber int) interface{} {
	const method = "chain_getText"
	response := call(callInterface{method: method, id: ""}, transactionHash, blockNumber)
	return response.Result
}

func IsAssetSpent(tracker string, transactionIndex int, shardId int, blockNumber int) interface{} {
	const method = "chain_isAssetSpent"
	response := call(callInterface{method: method, id: ""}, tracker, transactionIndex, shardId, blockNumber)
	return response.Result
}

func GetSeq(address string, blockNumber int) int {
	const method = "chain_getSeq"
	response := call(callInterface{method: method, id: ""}, address, blockNumber)
	return response.Result.(int)
}

func GetBalance(address string, blockNumber int) int {
	const method = "chain_getBalance"
	response := call(callInterface{method: method, id: ""}, address, blockNumber)
	return response.Result.(int)
}

func GetRegularKey(address string, blockNumber int) string {
	const method = "chain_getRegularKey"
	response := call(callInterface{method: method, id: ""}, address, blockNumber)
	return response.Result.(string)
}

func GetRegularKeyOwner(publicKey string, blockNumber int) string {
	const method = "chain_getRegularKey"
	response := call(callInterface{method: method, id: ""}, publicKey, blockNumber)
	return response.Result.(string)
}

func GetGenesisAccounts() []interface{} {
	const method = "chain_getGenesisAccounts"
	response := call(callInterface{method: method, id: ""})
	return response.Result.([]interface{})
}

func GetNumberOfShards(blockNumber int) float64 {
	const method = "chain_getNumberOfShards"
	response := call(callInterface{method: method, id: ""}, blockNumber)
	return response.Result.(float64)
}

func GetShardIdByHash(transactionHash string, blockNumber int) int {
	const method = "chain_getShardIdByHash"
	response := call(callInterface{method: method, id: ""}, transactionHash, blockNumber)
	return response.Result.(int)
}

func GetShardRoot(shardId int, blockNumber int) string {
	const method = "chain_getShardRoot"
	response := call(callInterface{method: method, id: ""}, shardId, blockNumber)
	return response.Result.(string)
}

func GetShardOwners(shardId int, blockNumber int) []string {
	const method = "chain_getShardOwners"
	response := call(callInterface{method: method, id: ""}, shardId, blockNumber)
	return response.Result.([]string)
}

func GetShardUsers(shardId int, blockNumber int) []string {
	const method = "chain_getShardUsers"
	response := call(callInterface{method: method, id: ""}, shardId, blockNumber)
	return response.Result.([]string)

}

func GetMiningReward(blockNumber int) int {
	const method = "chain_getMiningReward"
	response := call(callInterface{method: method, id: ""}, blockNumber)
	return response.Result.(int)
}

func GetMinTransactionFee(transactionType string, blockNumber int) int {
	const method = "chain_getMinTransactionFee"
	response := call(callInterface{method: method, id: ""}, transactionType, blockNumber)
	return response.Result.(int)
}

func GetCommonParams(blockNumber int) interface{} {
	const method = "chain_getCommonParamss"
	response := call(callInterface{method: method, id: ""}, blockNumber)
	return response.Result
}

func GetTermMetadata(blockNumber int) interface{} {
	const method = "chain_getTermMetadata"
	response := call(callInterface{method: method, id: ""}, blockNumber)
	return response.Result
}

func ExecuteTransaction(transaction interface{}, sender string) string {
	const method = "chain_getTermMetadata"
	response := call(callInterface{method: method, id: ""}, transaction, sender)
	return response.Result.(string)
}

func ExecuteVM(transaction interface{}, parameters [][][]int, indices int) []string {
	const method = "chain_executeVM"
	response := call(callInterface{method: method, id: ""}, transaction, parameters, indices)
	return response.Result.([]string)
}

func GetNetworkId() string {
	const method = "chain_getNetworkId"
	response := call(callInterface{method: method, id: ""})
	return response.Result.(string)
}
