package rpc

type Chain struct {
	rpcClient rpcClient
}

func (c *Chain) GetBestBlockNumber() (int, error) {
	const method = "chain_getBestBlockNumber"
	var blockNumber int
	err := c.rpcClient.call(callInterface{method: method, id: ""}, &blockNumber)
	return blockNumber, err
}

func (c *Chain) GetBestBlockID() (interface{}, error) {
	const method = "chain_getBestBlockId"
	var blockID interface{}
	err := c.rpcClient.call(callInterface{method: method, id: ""}, &blockID)
	return blockID, err
}

func (c *Chain) GetBlockHash(blockNumber int) (string, error) {
	const method = "chain_getBlockHash"
	var blockHash string
	err := c.rpcClient.call(callInterface{method: method, id: ""}, &blockHash, blockNumber)
	return blockHash, err
}

func (c *Chain) GetBlockByNumber(blockNumber int) (Block, error) {
	const method = "chain_getBlockByNumber"
	var block Block
	err := c.rpcClient.call(callInterface{method: method, id: ""}, &block, blockNumber)
	return block, err
}

func (c *Chain) GetBlockByHash(blockHash string) (interface{}, error) {
	const method = "chain_getBlockByHash"
	var block Block
	err := c.rpcClient.call(callInterface{method: method, id: ""}, &block, blockHash)
	return block, err
}

func (c *Chain) GetBlockTransactionCountByHash(blockHash string) (int, error) {
	const method = "chain_getBlockTransactionCountByHash"
	var count int
	err := c.rpcClient.call(callInterface{method: method, id: ""}, &count, blockHash)
	return count, err
}

func (c *Chain) GetTransaction(transactionHash string) (interface{}, error) {
	const method = "chain_getTransaction"
	var transaction interface{}
	err := c.rpcClient.call(callInterface{method: method, id: ""}, &transaction, transactionHash)
	return transaction, err
}

func (c *Chain) GetTransactionSigner(transactionHash string) (string, error) {
	const method = "chain_getTransactionSigner"
	var platformAddress string
	err := c.rpcClient.call(callInterface{method: method, id: ""}, &platformAddress, transactionHash)
	return platformAddress, err
}

func (c *Chain) ContainsTransaction(transactionHash string) (bool, error) {
	const method = "chain_containsTransaction"
	var contains bool
	err := c.rpcClient.call(callInterface{method: method, id: ""}, &contains, transactionHash)
	return contains, err
}

func (c *Chain) GetTransactionByTracker(tracker string) (interface{}, error) {
	const method = "chain_getTransactionByTracker"
	var transaction interface{}
	err := c.rpcClient.call(callInterface{method: method, id: ""}, &transaction, tracker)
	return transaction, err
}

func (c *Chain) GetAssetSchemeByTracker(tracker string) (interface{}, error) {
	const method = "chain_getAssetSchemeByTracker"
	var asset interface{}
	err := c.rpcClient.call(callInterface{method: method, id: ""}, &asset, tracker)
	return asset, err
}

func (c *Chain) GetAssetSchemeByType(assetType string) (interface{}, error) {
	const method = "chain_getAssetSchemeByType"
	var assetScheme interface{}
	err := c.rpcClient.call(callInterface{method: method, id: ""}, &assetScheme, assetType)
	return assetScheme, err
}

func (c *Chain) GetAsset(tracker string, transactionIndex int, shardID int, blockNumber int) (interface{}, error) {
	const method = "chain_getAsset"
	var asset interface{}
	err := c.rpcClient.call(callInterface{method: method, id: ""}, &asset, tracker, transactionIndex, shardID, blockNumber)
	return asset, err
}

func (c *Chain) GetText(transactionHash string, blockNumber int) (interface{}, error) {
	const method = "chain_getText"
	var text string
	err := c.rpcClient.call(callInterface{method: method, id: ""}, &text, transactionHash, blockNumber)
	return text, err
}

func (c *Chain) IsAssetSpent(tracker string, transactionIndex int, shardID int, blockNumber int) (interface{}, error) {
	const method = "chain_isAssetSpent"
	var spent bool
	err := c.rpcClient.call(callInterface{method: method, id: ""}, &spent, tracker, transactionIndex, shardID, blockNumber)
	return spent, err
}

func (c *Chain) GetSeq(address string, blockNumber int) (int, error) {
	const method = "chain_getSeq"
	var seq int
	err := c.rpcClient.call(callInterface{method: method, id: ""}, &seq, address, blockNumber)
	return seq, err
}

func (c *Chain) GetBalance(address string, blockNumber int) (int, error) {
	const method = "chain_getBalance"
	var balance int
	err := c.rpcClient.call(callInterface{method: method, id: ""}, &balance, address, blockNumber)
	return balance, err
}

func (c *Chain) GetRegularKey(address string, blockNumber int) (string, error) {
	const method = "chain_getRegularKey"
	var key string
	err := c.rpcClient.call(callInterface{method: method, id: ""}, &key, address, blockNumber)
	return key, err
}

func (c *Chain) GetRegularKeyOwner(publicKey string, blockNumber int) (string, error) {
	const method = "chain_getRegularKeyOwner"
	var platformAddress string
	err := c.rpcClient.call(callInterface{method: method, id: ""}, &platformAddress, publicKey, blockNumber)
	return platformAddress, err
}

func (c *Chain) GetGenesisAccounts() ([]interface{}, error) {
	const method = "chain_getGenesisAccounts"
	var accounts []interface{}
	err := c.rpcClient.call(callInterface{method: method, id: ""}, &accounts)
	return accounts, err
}

func (c *Chain) GetNumberOfShards(blockNumber int) (float64, error) {
	const method = "chain_getNumberOfShards"
	var num float64
	err := c.rpcClient.call(callInterface{method: method, id: ""}, &num, blockNumber)
	return num, err
}

func (c *Chain) GetShardIDByHash(transactionHash string, blockNumber int) (int, error) {
	const method = "chain_getShardIdByHash"
	var shardID int
	err := c.rpcClient.call(callInterface{method: method, id: ""}, &shardID, transactionHash, blockNumber)
	return shardID, err
}

func (c *Chain) GetShardRoot(shardID int, blockNumber int) (string, error) {
	const method = "chain_getShardRoot"
	var shardRoot string
	err := c.rpcClient.call(callInterface{method: method, id: ""}, &shardRoot, shardID, blockNumber)
	return shardRoot, err
}

func (c *Chain) GetShardOwners(shardID int, blockNumber int) ([]string, error) {
	const method = "chain_getShardOwners"
	var owners []string
	err := c.rpcClient.call(callInterface{method: method, id: ""}, &owners, shardID, blockNumber)
	return owners, err
}

func (c *Chain) GetShardUsers(shardID int, blockNumber int) ([]string, error) {
	const method = "chain_getShardUsers"
	var users []string
	err := c.rpcClient.call(callInterface{method: method, id: ""}, &users, shardID, blockNumber)
	return users, err

}

func (c *Chain) GetMiningReward(blockNumber int) (int, error) {
	const method = "chain_getMiningReward"
	var reward int
	err := c.rpcClient.call(callInterface{method: method, id: ""}, &reward, blockNumber)
	return reward, err
}

func (c *Chain) GetMinTransactionFee(transactionType string, blockNumber int) (int, error) {
	const method = "chain_getMinTransactionFee"
	var minFee int
	err := c.rpcClient.call(callInterface{method: method, id: ""}, &minFee, transactionType, blockNumber)
	return minFee, err
}

func (c *Chain) GetCommonParams(blockNumber int) (interface{}, error) {
	const method = "chain_getCommonParamss"
	var commonParams interface{}
	err := c.rpcClient.call(callInterface{method: method, id: ""}, &commonParams, blockNumber)
	return commonParams, err
}

func (c *Chain) GetTermMetadata(blockNumber int) (interface{}, error) {
	const method = "chain_getTermMetadata"
	var metadata interface{}
	err := c.rpcClient.call(callInterface{method: method, id: ""}, &metadata, blockNumber)
	return metadata, err
}

func (c *Chain) ExecuteTransaction(transaction interface{}, sender string) (string, error) {
	const method = "chain_executeTransaction"
	var result string
	err := c.rpcClient.call(callInterface{method: method, id: ""}, &result, transaction, sender)
	return result, err
}

func (c *Chain) ExecuteVM(transaction interface{}, parameters [][][]int, indices int) ([]string, error) {
	const method = "chain_executeVM"
	var result []string
	err := c.rpcClient.call(callInterface{method: method, id: ""}, &result, transaction, parameters, indices)
	return result, err
}

func (c *Chain) GetNetworkID() (string, error) {
	const method = "chain_getNetworkId"
	var networkID string
	err := c.rpcClient.call(callInterface{method: method, id: ""}, &networkID)
	return networkID, err
}
