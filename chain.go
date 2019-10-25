package rpc

type Chain struct {
	rpcClient rpcClient
}

func (c *Chain) GetBestBlockNumber() int {
	const method = "chain_getBestBlockNumber"
	var blockNumber int
	c.rpcClient.call(callInterface{method: method, id: ""}, &blockNumber)
	return blockNumber
}

func (c *Chain) GetBestBlockID() interface{} {
	const method = "chain_getBestBlockId"
	var blockID interface{}
	c.rpcClient.call(callInterface{method: method, id: ""}, &blockID)
	return blockID
}

func (c *Chain) GetBlockHash(blockNumber int) string {
	const method = "chain_getBlockHash"
	var blockHash string
	c.rpcClient.call(callInterface{method: method, id: ""}, &blockHash, blockNumber)
	return blockHash
}

func (c *Chain) GetBlockByNumber(blockNumber int) Block {
	const method = "chain_getBlockByNumber"
	var block Block
	c.rpcClient.call(callInterface{method: method, id: ""}, &block, blockNumber)
	return block
}

func (c *Chain) GetBlockByHash(blockHash string) interface{} {
	const method = "chain_getBlockByHash"
	var block Block
	c.rpcClient.call(callInterface{method: method, id: ""}, &block, blockHash)
	return block
}

func (c *Chain) GetBlockTransactionCountByHash(blockHash string) int {
	const method = "chain_getBlockTransactionCountByHash"
	var count int
	c.rpcClient.call(callInterface{method: method, id: ""}, &count, blockHash)
	return count
}

func (c *Chain) GetTransaction(transactionHash string) interface{} {
	const method = "chain_getTransaction"
	var transaction interface{}
	c.rpcClient.call(callInterface{method: method, id: ""}, &transaction, transactionHash)
	return transaction
}

func (c *Chain) GetTransactionSigner(transactionHash string) string {
	const method = "chain_getTransactionSigner"
	var platformAddress string
	c.rpcClient.call(callInterface{method: method, id: ""}, &platformAddress, transactionHash)
	return platformAddress
}

func (c *Chain) ContainsTransaction(transactionHash string) bool {
	const method = "chain_containsTransaction"
	var contains bool
	c.rpcClient.call(callInterface{method: method, id: ""}, &contains, transactionHash)
	return contains
}

func (c *Chain) GetTransactionByTracker(tracker string) interface{} {
	const method = "chain_getTransactionByTracker"
	var transaction interface{}
	c.rpcClient.call(callInterface{method: method, id: ""}, &transaction, tracker)
	return transaction
}

func (c *Chain) GetAssetSchemeByTracker(tracker string) interface{} {
	const method = "chain_getAssetSchemeByTracker"
	var asset interface{}
	c.rpcClient.call(callInterface{method: method, id: ""}, &asset, tracker)
	return asset
}

func (c *Chain) GetAssetSchemeByType(assetType string) interface{} {
	const method = "chain_getAssetSchemeByType"
	var assetScheme interface{}
	c.rpcClient.call(callInterface{method: method, id: ""}, &assetScheme, assetType)
	return assetScheme
}

func (c *Chain) GetAsset(tracker string, transactionIndex int, shardID int, blockNumber int) interface{} {
	const method = "chain_getAsset"
	var asset interface{}
	c.rpcClient.call(callInterface{method: method, id: ""}, &asset, tracker, transactionIndex, shardID, blockNumber)
	return asset
}

func (c *Chain) GetText(transactionHash string, blockNumber int) interface{} {
	const method = "chain_getText"
	var text string
	c.rpcClient.call(callInterface{method: method, id: ""}, &text, transactionHash, blockNumber)
	return text
}

func (c *Chain) IsAssetSpent(tracker string, transactionIndex int, shardID int, blockNumber int) interface{} {
	const method = "chain_isAssetSpent"
	var spent bool
	c.rpcClient.call(callInterface{method: method, id: ""}, &spent, tracker, transactionIndex, shardID, blockNumber)
	return spent
}

func (c *Chain) GetSeq(address string, blockNumber int) int {
	const method = "chain_getSeq"
	var seq int
	c.rpcClient.call(callInterface{method: method, id: ""}, &seq, address, blockNumber)
	return seq
}

func (c *Chain) GetBalance(address string, blockNumber int) int {
	const method = "chain_getBalance"
	var balance int
	c.rpcClient.call(callInterface{method: method, id: ""}, &balance, address, blockNumber)
	return balance
}

func (c *Chain) GetRegularKey(address string, blockNumber int) string {
	const method = "chain_getRegularKey"
	var key string
	c.rpcClient.call(callInterface{method: method, id: ""}, &key, address, blockNumber)
	return key
}

func (c *Chain) GetRegularKeyOwner(publicKey string, blockNumber int) string {
	const method = "chain_getRegularKeyOwner"
	var platformAddress string
	c.rpcClient.call(callInterface{method: method, id: ""}, &platformAddress, publicKey, blockNumber)
	return platformAddress
}

func (c *Chain) GetGenesisAccounts() []interface{} {
	const method = "chain_getGenesisAccounts"
	var accounts []interface{}
	c.rpcClient.call(callInterface{method: method, id: ""}, &accounts)
	return accounts
}

func (c *Chain) GetNumberOfShards(blockNumber int) float64 {
	const method = "chain_getNumberOfShards"
	var num float64
	c.rpcClient.call(callInterface{method: method, id: ""}, &num, blockNumber)
	return num
}

func (c *Chain) GetShardIDByHash(transactionHash string, blockNumber int) int {
	const method = "chain_getShardIdByHash"
	var shardID int
	c.rpcClient.call(callInterface{method: method, id: ""}, &shardID, transactionHash, blockNumber)
	return shardID
}

func (c *Chain) GetShardRoot(shardID int, blockNumber int) string {
	const method = "chain_getShardRoot"
	var shardRoot string
	c.rpcClient.call(callInterface{method: method, id: ""}, &shardRoot, shardID, blockNumber)
	return shardRoot
}

func (c *Chain) GetShardOwners(shardID int, blockNumber int) []string {
	const method = "chain_getShardOwners"
	var owners []string
	c.rpcClient.call(callInterface{method: method, id: ""}, &owners, shardID, blockNumber)
	return owners
}

func (c *Chain) GetShardUsers(shardID int, blockNumber int) []string {
	const method = "chain_getShardUsers"
	var users []string
	c.rpcClient.call(callInterface{method: method, id: ""}, &users, shardID, blockNumber)
	return users

}

func (c *Chain) GetMiningReward(blockNumber int) int {
	const method = "chain_getMiningReward"
	var reward int
	c.rpcClient.call(callInterface{method: method, id: ""}, &reward, blockNumber)
	return reward
}

func (c *Chain) GetMinTransactionFee(transactionType string, blockNumber int) int {
	const method = "chain_getMinTransactionFee"
	var minFee int
	c.rpcClient.call(callInterface{method: method, id: ""}, &minFee, transactionType, blockNumber)
	return minFee
}

func (c *Chain) GetCommonParams(blockNumber int) interface{} {
	const method = "chain_getCommonParamss"
	var commonParams interface{}
	c.rpcClient.call(callInterface{method: method, id: ""}, &commonParams, blockNumber)
	return commonParams
}

func (c *Chain) GetTermMetadata(blockNumber int) interface{} {
	const method = "chain_getTermMetadata"
	var metadata interface{}
	c.rpcClient.call(callInterface{method: method, id: ""}, &metadata, blockNumber)
	return metadata
}

func (c *Chain) ExecuteTransaction(transaction interface{}, sender string) string {
	const method = "chain_executeTransaction"
	var result string
	c.rpcClient.call(callInterface{method: method, id: ""}, &result, transaction, sender)
	return result
}

func (c *Chain) ExecuteVM(transaction interface{}, parameters [][][]int, indices int) []string {
	const method = "chain_executeVM"
	var result []string
	c.rpcClient.call(callInterface{method: method, id: ""}, &result, transaction, parameters, indices)
	return result
}

func (c *Chain) GetNetworkID() string {
	const method = "chain_getNetworkId"
	var networkID string
	c.rpcClient.call(callInterface{method: method, id: ""}, &networkID)
	return networkID
}
